package author

import (
	"database/sql"
	"fmt"

	"github.com/JoaoGeraldoS/API-biblioteca/app/models"
)

func ReadAuthors(db *sql.DB) ([]models.Authors, error) {
	rows, err := db.Query(`SELECT * FROM authors`)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar dados! %v", err)
	}
	defer rows.Close()

	var authors []models.Authors

	for rows.Next() {
		var author models.Authors

		err = rows.Scan(&author.ID, &author.Name, &author.Description, &author.Photo)
		if err != nil {
			return nil, fmt.Errorf("erro ao scanear dados %v", err)
		}
		authors = append(authors, author)
	}

	return authors, nil
}
