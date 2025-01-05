package books

import (
	"database/sql"
	"log"

	"github.com/JoaoGeraldoS/API-biblioteca/app/models"
)

func ReadUnicBook(id string, db *sql.DB) (models.Books, error) {

	query := (`select b.id, b.title, b.author, b.description, b.content,
		DATE_FORMAT(b.created_at, '%d/%m/%y %H:%i:%s') AS created_at, DATE_FORMAT(b.updated_at, '%d/%m/%y %H:%i:%s') AS updated_at,
		c.id, c.name, DATE_FORMAT(c.created_at, '%d/%m/%y %H:%i:%s') AS created_at
		from intermediaria i
		join books b on i.book_id = b.id
		join categories c on i.category_id = c.id WHERE b.id = ?`)

	rows, err := db.Query(query, id)
	if err != nil {
		log.Printf("Erro ao buscar dados! %v", err)
	}
	defer rows.Close()

	var book models.Books
	book.Categories = []models.Categories{}

	for rows.Next() {
		var category models.Categories

		if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Description, &book.Content, &book.Created_at, &book.Updated_at, &category.ID, &category.Name, &category.Created_at_c); err != nil {
			log.Println("Erro ao buscar dados")
		}
		book.Categories = append(book.Categories, category)
	}

	if book.ID == 0 {
		log.Printf("Livro não encontrado %v", err)
	}

	return book, nil

}
