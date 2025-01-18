package utils

import (
	"database/sql"
	"fmt"
)

func IdsCategories(tx *sql.Tx, names []string) ([]int64, error) {
	ids := []int64{}
	query := `SELECT id FROM categories WHERE name = ?`
	// Busca no banco se existe id relacionado ao nome da categoria

	for _, name := range names { // Itera sobre o lista de nomes de categorias recebidas
		var id int64

		err := tx.QueryRow(query, name).Scan(&id) // Realiza a query, e escanea os ids
		if err != nil {
			if err == sql.ErrNoRows { // Caso der não exista a categoria
				return nil, fmt.Errorf("categoria '%s' não encontrada", name)
			}
			return nil, fmt.Errorf("erro ao buscar ID da categoria '%s': %v", name, err)
		}
		ids = append(ids, id) // Adiciona o id ao array de a ids
	}
	return ids, nil // Retorna o arry e um erro
}

func Relationship(tx *sql.Tx, bookID int64, categoriesIDs []int64) error {
	query := `INSERT INTO intermediary(book_id, category_id) VALUES (?, ?)`
	// Insere a id da categoria e o id do livro relacionado na tabela intermediaria
	for _, categoryID := range categoriesIDs { // Faz a iteração e adiciona no banco
		_, err := tx.Exec(query, bookID, categoryID)
		if err != nil {
			return fmt.Errorf("Erro ao inserir relacionamento %v", err)
		}
	}
	return nil
}
