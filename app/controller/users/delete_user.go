package users

import (
	"database/sql"
	"fmt"
)

func DeleteUser(id string, db *sql.DB) error {
	query := `DELETE FROM users WHERE id = ?`

	row, err := db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("erro ao deletar usuario")
	}

	rowAffected, err := row.RowsAffected()
	if err != nil || rowAffected == 0 {
		return fmt.Errorf("nenhuma linha afetada!")
	}

	return nil
}
