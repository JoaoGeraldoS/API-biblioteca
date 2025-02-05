package books

import (
	"database/sql"
	"fmt"

	"github.com/JoaoGeraldoS/API-biblioteca/app/models"
)

func Relation(db *sql.DB, relation models.Intermediary) (*models.Intermediary, error) {
	query := `INSERT INTO intermediary(book_id, category_id) VALUES (?, ?)`

	response, err := db.Exec(query, relation.BookID, relation.CategoryID)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar livros ou categorias! %v", err)
	}

	id, err := response.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("erro ao adicionar id")
	}

	relation.ID = id

	return &relation, nil
}
