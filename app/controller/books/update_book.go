package books

import (
	"database/sql"
	"fmt"

	"github.com/JoaoGeraldoS/API-biblioteca/app/models"
)

func UpdateBook(id string, db *sql.DB, book *models.Books) (*models.Books, error) {
	query := `UPDATE books SET title = ? , author = ? , description = ? , content = ? WHERE id = ?`

	response, err := db.Exec(query, book.Title, book.Author, book.Description, book.Content, id)
	if err != nil {
		return nil, fmt.Errorf("erro ao atualizar dados! %v", err)
	}

	rowsAffected, err := response.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return nil, fmt.Errorf("erro ao validar dados! %v", err)
	}

	id_book, err := response.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("erro ao verificar id! %v", err)
	}

	book.ID = id_book
	return book, nil

}
