package books

// Leitura dos livros, retorna uma lista de todos os livors do sistema

import (
	"database/sql"
	"log"

	"github.com/JoaoGeraldoS/API-biblioteca/app/models"
)

func ReadBook(db *sql.DB) ([]models.Books, error) {
	rows, err := db.Query(`select id, title, author, description, content,
	 DATE_FORMAT(created_at, '%d/%m/%y %H:%i:%s') AS created_at, DATE_FORMAT(updated_at, '%d/%m/%y %H:%i:%s') AS updated_at from books`)
	if err != nil {

		log.Println("Erro ao buscar dados")
	}
	defer rows.Close()

	var books []models.Books

	for rows.Next() {
		var book models.Books

		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Descrition, &book.Content, &book.Created_at, &book.Updated_at)
		if err != nil {
			log.Println("Erro ao verificar dados")
		}

		books = append(books, book)
	}

	return books, nil
}
