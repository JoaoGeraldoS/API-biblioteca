package books

import (
	"database/sql"
	"fmt"

	"github.com/JoaoGeraldoS/API-biblioteca/app/models"
)

func CreateBook(db *sql.DB, book models.Books) (models.Books, error) {
	query := `INSERT INTO books (title, author, description, content) VALUES (?, ?, ?, ?)`

	response, err := db.Exec(query, &book.Title, &book.Author, &book.Descrition, &book.Content)
	if err != nil {
		fmt.Println("Erro ao inserir os dados!", err)
	}

	id, err := response.LastInsertId()
	if err != nil {
		fmt.Println("Erro ao verificar id", err)
	}

	book.ID = id

	return book, nil
}
