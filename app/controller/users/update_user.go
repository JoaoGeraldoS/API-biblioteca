package users

import (
	"database/sql"
	"fmt"

	"github.com/JoaoGeraldoS/API-biblioteca/app/models"
)

func UpdateUser(id string, db *sql.DB, user *models.Users, photo string) (*models.Users, error) {
	query := `UPDATE users SET name = ?, bio = ?, photo = ? WHERE id = ?`

	response, err := db.Exec(query, user.Name, user.Bio, photo, id)
	if err != nil {
		return nil, fmt.Errorf("erro ao atualizar perfil! %v", err)
	}

	rowsAffected, err := response.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return nil, fmt.Errorf("erro, nenhuma linha afetada! %v", err)
	}

	id_user, err := response.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("erro ao verificar id! %v", err)
	}

	user.ID = id_user
	user.Photo = photo
	return user, nil
}
