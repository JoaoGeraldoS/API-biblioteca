package categories

import (
	"database/sql"
	"log"
)

func DeleteCategory(id string, db *sql.DB) error {
	query := `DELETE FROM categories WHERE id = ?`

	response, err := db.Exec(query, id)
	if err != nil {
		log.Printf("Erro ao apagar dado! %v", err)
	}

	rowsAffected, err := response.RowsAffected()
	if err != nil || rowsAffected == 0 {
		log.Printf("Nenhum dado apagado! %v", err)
	}

	return nil
}
