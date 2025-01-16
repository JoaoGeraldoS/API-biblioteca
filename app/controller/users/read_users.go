package users

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/JoaoGeraldoS/API-biblioteca/app/models"
)

func ReadUsers(id, name string, db *sql.DB) ([]models.Users, error) {
	query := `SELECT id, name, email, DATE_FORMAT(created_at, '%d/%m/%y %H:%i:%s') as created_at,
		DATE_FORMAT(update_at, '%d/%m/%y %H:%i:%s') as updated_at FROM users`

	conditions := []string{}
	params := []interface{}{}

	if id != "" {
		conditions = append(conditions, "id = ?")
		params = append(params, id)
	}

	if name != "" {
		conditions = append(conditions, "name like ?")
		params = append(params, fmt.Sprintf("%s%%", name))
	}

	if len(conditions) > 0 {
		query += " where " + strings.Join(conditions, " ")
	}

	row, err := db.Query(query, params...)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar usuarios! %v", err)
	}
	defer row.Close()

	var users []models.Users

	for row.Next() {
		var user models.Users

		err = row.Scan(&user.ID, &user.Name, &user.Email, &user.Created_at, &user.Updated_at)
		if err != nil {
			return nil, fmt.Errorf("erro ao escanear dados! %v", err)
		}

		users = append(users, user)
	}

	return users, nil

}
