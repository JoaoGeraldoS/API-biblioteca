package books

// Leitura dos livros, retorna uma lista de todos os livors do sistema

import (
	"database/sql"
	"log"

	"github.com/JoaoGeraldoS/API-biblioteca/app/models"
)

func ReadBook(db *sql.DB) ([]models.Books, error) {
	booksMap := make(map[int64]*models.Books)
	rows, err := db.Query(`select b.id, b.title, b.author, b.description, b.content,
		DATE_FORMAT(b.created_at, '%d/%m/%y %H:%i:%s') AS created_at, DATE_FORMAT(b.updated_at, '%d/%m/%y %H:%i:%s') AS updated_at,
		c.id, c.name, DATE_FORMAT(c.created_at, '%d/%m/%y %H:%i:%s') AS created_at
		from intermediaria i
		join books b on i.book_id = b.id
		join categories c on i.category_id = c.id`,
	)
	if err != nil {
		log.Println("Erro ao buscar dados")
	}
	defer rows.Close()

	for rows.Next() {
		var (
			bookID, categoryID                                          int64
			title, author, description, content, created_at, updated_at string
			categoryName, created_at_c                                  string
		)

		err := rows.Scan(&bookID, &title, &author, &description, &content, &created_at, &updated_at, &categoryID, &categoryName, &created_at_c)
		if err != nil {
			log.Println("Erro ao verificar dados")
		}

		if _, exists := booksMap[bookID]; !exists {
			booksMap[bookID] = &models.Books{
				ID:          bookID,
				Title:       title,
				Author:      author,
				Description: description,
				Content:     content,
				Created_at:  created_at,
				Updated_at:  updated_at,
				Categories:  []models.Categories{},
			}
		}

		booksMap[bookID].Categories = append(booksMap[bookID].Categories, models.Categories{
			ID:           categoryID,
			Name:         categoryName,
			Created_at_c: created_at_c,
		})
	}

	var books []models.Books
	for _, book := range booksMap {
		books = append(books, *book)
	}

	return books, nil
}
