package books

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/JoaoGeraldoS/API-biblioteca/app/models"
	"github.com/JoaoGeraldoS/API-biblioteca/app/utils"
	"github.com/JoaoGeraldoS/API-biblioteca/app/validacao"
)

func CreateBook(db *sql.DB, book *models.Books, categories []string, file string) (*models.Books, error) {
	tx, err := db.Begin()
	if err != nil {
		log.Printf("Erro ao iniciar transação: %v", err)
	}

	query := `INSERT INTO books (title, author, description, content, img) VALUES (?, ?, ?, ?, ?)`

	response, err := db.Exec(query, book.Title, book.Author, book.Description, book.Content, file)
	if err != nil {
		fmt.Println("Erro ao inserir os dados!", err)
	}

	id, err := response.LastInsertId()
	if err != nil {
		fmt.Println("Erro ao verificar id", err)
	}

	book.ID = id
	book.Img = file

	for _, category := range book.Categories {
		_, err := validacao.ValidaCategories(db, category.Name)
		if err != nil {
			log.Printf("Erro ao verificar ou criar categoria: %v", err)
		}
	}

	categoryIDs, err := utils.IdsCategories(tx, categories)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := utils.Relationship(tx, book.ID, categoryIDs); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		log.Printf("Erro ao confirmar transação")
	}

	return book, nil
}
