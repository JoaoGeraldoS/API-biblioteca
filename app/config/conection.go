package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func Conection() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}

	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		os.Getenv("USER"), os.Getenv("PASSWORD"), os.Getenv("HOST"), os.Getenv("PORT"), os.Getenv("DATABASE"),
	)

	session, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatalln("Erro ao caonectar com o banco de dados!", err)
	}

	if err := session.Ping(); err != nil {
		log.Fatalln("Erro ao verficar conexão")
	}

	return session
}
