package users

import (
	"database/sql"
	"fmt"

	"github.com/JoaoGeraldoS/API-biblioteca/app/middleware"
	"golang.org/x/crypto/bcrypt"
)

func Login(db *sql.DB, username, password string) (string, error) {
	var senhaHash, role string

	err := db.QueryRow(`SELECT password, role FROM users WHERE username = ?`, username).Scan(&senhaHash, &role)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("usuario ou senha invalidos")
		}
		return "", fmt.Errorf("erro ao buscar usuario")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(senhaHash), []byte(password)); err != nil {
		return "", fmt.Errorf("usuario ou senha invalidos")
	}

	token, err := middleware.GerarToken(username, role)
	if err != nil {
		return "", fmt.Errorf("erro ao gerar token")
	}

	return token, nil
}
