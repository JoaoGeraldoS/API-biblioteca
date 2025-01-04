package utils

import (
	"database/sql"
	"fmt"
)

func IdsCategories(tx *sql.Tx, names []string) ([]int64, error) {
	ids := []int64{}
	query := `SELECT id FROM categories WHERE name = ?`

	for _, name := range names {
		var id int64

		err := tx.QueryRow(query, name).Scan(&id)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, fmt.Errorf("Categoria '%s' não encontrada", name)
			}
			return nil, fmt.Errorf("Erro ao buscar ID da categoria '%s': %v", name, err)
		}
		ids = append(ids, id)
	}
	return ids, nil
}

func Relationship(tx *sql.Tx, bookID int64, categoriesIDs []int64) error {
	query := `INSERT INTO intermediaria(book_id, category_id) VALUES (?, ?)`

	for _, categoryID := range categoriesIDs {
		_, err := tx.Exec(query, bookID, categoryID)
		if err != nil {
			return fmt.Errorf("Erro ao inserir relacionamento %v", err)
		}
	}
	return nil
}
