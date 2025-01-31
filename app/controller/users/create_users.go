package users

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/JoaoGeraldoS/API-biblioteca/app/models"
	"golang.org/x/crypto/bcrypt"
)

func CreateUsers(db *sql.DB, user *models.Users) (*models.Users, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Erro ao gerar hash da senha! %v", err)
		return nil, fmt.Errorf("erro ao gerar hash da senha")
	}

	user.Password = string(hashedPassword)

	if user.Role == "" {
		user.Role = "User"
	}

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
