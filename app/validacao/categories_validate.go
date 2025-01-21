package validacao

import (
	"database/sql"
	"fmt"
)

func ValidaCategories(db *sql.DB, categoryName string) (int64, error) {
	var categoryID int64
	query := `SELECT id FROM categories WHERE name = ?`
	// Verifica se existe a categoria recebida pelo nome no banco de dados
	err := db.QueryRow(query, categoryName).Scan(&categoryID)
	if err != nil {
		if err == sql.ErrNoRows { // Se a categoria não existir, a categoria é criada no banco
			insert := `INSERT INTO categories (name) VALUES (?)`
			response, err := db.Exec(insert, categoryName)
			if err != nil {
				return 0, fmt.Errorf("erro ao inserir categoria: %v", err)
			}

			categoryID, err = response.LastInsertId()
			if err != nil {
				return 0, fmt.Errorf("erro ao obter o ID da categoria inserida: %v", err)
			}
			return categoryID, nil // Depois de criada, retorna a id da categoria e o erro nulo
		}
		return 0, fmt.Errorf("erro ao verificar a categoria: %v", err)
	}
	return categoryID, nil // Se a categoria existir, retorna a id da categoria e o erro nulo
}
