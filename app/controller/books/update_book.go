package books

import (
	"database/sql"
	"log"

	"github.com/JoaoGeraldoS/API-biblioteca/app/models"
)

func UpdateBook(id string, db *sql.DB, book *models.Books) (*models.Books, error) {
	query := `UPDATE books SET title = ? , author = ? , description = ? , content = ? WHERE id = ?`

	response, err := db.Exec(query, book.Title, book.Author, book.Description, book.Content, id)
	if err != nil {
		log.Println("Erro ao atualizar dados!")
	}

	rowsAffected, err := response.RowsAffected()
	if err != nil || rowsAffected == 0 {
		log.Println("Erro ao validar dados")
	}

	id_book, err := response.LastInsertId()
	if err != nil {
		log.Println("Erro ao verificar id")
	}

	book.ID = id_book
	return book, nil

}
