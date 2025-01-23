package comments

import (
	"database/sql"
	"fmt"

	"github.com/JoaoGeraldoS/API-biblioteca/app/models"
)

func CreateComments(db *sql.DB, comment *models.Comments) (*models.Comments, error) {
	query := `INSERT INTO comments (content, user_id, book_id) VALUES (?, ?, ?)`

	response, err := db.Exec(query, comment.Content, comment.User_id, comment.Book_id)
	if err != nil {
		return nil, fmt.Errorf("erro ao inserir comentarios")
	}

	id, err := response.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("erro ao verificar o id")
	}

	comment.ID = id

	return comment, nil
}
