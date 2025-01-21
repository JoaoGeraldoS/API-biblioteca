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
	tx, err := db.Begin() // Inicia uma transação
	if err != nil {
		return nil, fmt.Errorf("erro ao iniciar transação: %v", err)
	}

	query := `INSERT INTO books (title, author_id, description, content, img) VALUES (?, ?, ?, ?, ?)`
	// Insere os dados do livvro no banco de dados
	response, err := db.Exec(query, book.Title, book.AuthorId, book.Description, book.Content, file)
	if err != nil {
		return nil, fmt.Errorf("erro ao inserir os dados! %v", err)
	}

	id, err := response.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("erro ao verificar id! %v", err)
	}

	book.ID = id
	book.Img = file

	// Faz a iteração no array de categorias recebidas pelo json
	for _, category := range categories {
		_, err := validacao.ValidaCategories(db, category) // Passa o parametro para a função de validação de categorias

		if err != nil {
			return nil, fmt.Errorf("erro ao verificar ou criar categoria: %v", err) // Caso a verficação da erro, esse erro é rotornado
		}
	}

	categoryIDs, err := utils.IdsCategories(tx, categories) // Busca as categorias existentes pelo id
	if err != nil {
		tx.Rollback() // Se um não exitir, para a trasação
		return nil, err
	}

	if err := utils.Relationship(tx, book.ID, categoryIDs); err != nil { /* Passa os parametros para a inserção dos ids categoria e
		livros para tabela intermediaria*/
		tx.Rollback()
		return nil, err
	}

	if err = tx.Commit(); err != nil { // Se tudo ocorrer bem, confirma a trasação e salva no banco
		log.Fatalln("erro ao confirmar transação")
	}

	return book, nil
}
