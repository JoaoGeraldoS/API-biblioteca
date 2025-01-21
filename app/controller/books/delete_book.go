package books

import (
	"database/sql"
	"log"
)

func DeleteBook(id string, db *sql.DB) error {
	query := `delete from books where id = ?`

	row, err := db.Exec(query, id)
	if err != nil {
		log.Println("Erro ao deletear dados")
	}

	rowAffected, err := row.RowsAffected()
	if err != nil || rowAffected == 0 {
		log.Println("Error dados nao encontrados")
	}

	return nil
}
