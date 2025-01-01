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

	entrada := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		os.Getenv("USER"), os.Getenv("PASSWORD"), os.Getenv("HOST"), os.Getenv("PORT"), os.Getenv("DATABASE"),
	)

	conexao, err := sql.Open("mysql", entrada)
	if err != nil {
		log.Fatalln("Erro ao caonectar com o banco de dados!", err)
	}

	if err := conexao.Ping(); err != nil {
		log.Fatalln("Erro ao verficar conexão")
	}
	fmt.Println("Conexão estabelecida!")

	return conexao
}
