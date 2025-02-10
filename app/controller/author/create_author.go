package author

import (
	"database/sql"
	"fmt"

	"github.com/JoaoGeraldoS/API-biblioteca/app/models"
)

func CreateAuthor(db *sql.DB, author *models.Authors, file string) (*models.Authors, error) {
	query := `INSERT INTO authors (name, description, photo) VALUES (?, ?, ?)`

	response, err := db.Exec(query, author.Name, author.Description, file)
	if err != nil {
		return nil, fmt.Errorf("erro ao inserir os dados! %v", err)
	}

	id, err := response.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("erro ao verificar id! %v", err)
	}

	author.ID = id
	author.Photo = file

	return author, nil
}
