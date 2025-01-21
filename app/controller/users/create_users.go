package users

import (
	"database/sql"
	"log"

	"github.com/JoaoGeraldoS/API-biblioteca/app/models"
)

func CreateUsers(db *sql.DB, user *models.Users) (*models.Users, error) {
	query := `INSERT INTO users(name, email, password, username, role) VALUES (?, ?, ?, ?, ?)`

	response, err := db.Exec(query, user.Name, user.Email, user.Password, user.Username, user.Role)

	if err != nil {
		log.Printf("Erro ao inserir dados! %v", err)
	}

	id, err := response.LastInsertId()
	if err != nil {
		log.Printf("Erro ao verificar id! %v", err)
	}

	user.ID = id

	return user, nil
}
