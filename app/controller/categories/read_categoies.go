package categories

import (
	"database/sql"
	"fmt"

	"github.com/JoaoGeraldoS/API-biblioteca/app/models"
)

func ReadCategories(db *sql.DB) ([]models.Categories, error) {
	query := `SELECT id, name, DATE_FORMAT(created_at, '%d/%m/%y %H:%i:%s') AS created_at FROM categories`

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar dados %v", err)
	}
	defer rows.Close()

	var categories []models.Categories

	for rows.Next() {
		var category models.Categories

		err = rows.Scan(&category.ID, &category.Name, &category.CreatedAtC)
		if err != nil {
			return nil, fmt.Errorf("erro ao scanear dados %v", err)
		}
		categories = append(categories, category)
	}

	return categories, nil
}
