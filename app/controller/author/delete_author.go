package author

import (
	"database/sql"
	"fmt"
)

func DeleteAuthors(id string, db *sql.DB) error {
	query := `DELETE FROM authors WHERE id = ?`

	response, err := db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("erro ao deletear dados! %v", err)
	}

	rowsAffected, err := response.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return fmt.Errorf("error dados nao encontrados! %v", err)
	}

	return err
}
