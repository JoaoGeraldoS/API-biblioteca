package books

import (
	"database/sql"
	"fmt"
)

func DeleteBook(id string, db *sql.DB) error {
	query := `delete from books where id = ?`

	row, err := db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("erro ao deletear dados! %v", err)
	}

	rowAffected, err := row.RowsAffected()
	if err != nil || rowAffected == 0 {
		return fmt.Errorf("error dados nao encontrados! %v", err)
	}

	return nil
}
