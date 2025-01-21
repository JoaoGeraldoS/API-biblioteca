package categories

import (
	"database/sql"
	"log"

	"github.com/JoaoGeraldoS/API-biblioteca/app/models"
)

func CreateCategory(db *sql.DB, cat models.Categories) (models.Categories, error) {
	query := `INSERT INTO categories (name) VALUE (?)`

	response, err := db.Exec(query, &cat.Name)
	if err != nil {
		log.Println("Erro ao salvar categoria!")
	}

	id, err := response.LastInsertId()
	if err != nil {
		log.Println("Erro ao verificar id")
	}

	cat.ID = id
	return cat, nil
}
