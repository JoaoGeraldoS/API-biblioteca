package validacao

import (
	"database/sql"
	"fmt"
)

func ValidaCategories(db *sql.DB, categoryName string) (int64, error) {
	var categoryID int64
	query := `SELECT id FROM categories WHERE name = ?`

	err := db.QueryRow(query, categoryName).Scan(&categoryID)
	if err != nil {
		if err == sql.ErrNoRows {
			insert := `INSERT INTO categories (name) VALUES (?)`
			response, err := db.Exec(insert, categoryName)
			if err != nil {
				return 0, fmt.Errorf("Erro ao inserir categoria: %v", err)
			}
			categoryID, err = response.LastInsertId()
			if err != nil {
				return 0, fmt.Errorf("Erro ao obter o ID da categoria inserida: %v", err)
			}
			return categoryID, nil
		}
		return 0, fmt.Errorf("Erro ao verificar a categoria: %v", err)
	}
	return categoryID, nil
}
