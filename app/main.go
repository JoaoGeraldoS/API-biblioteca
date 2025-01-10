package main

import (
	"os"

	"github.com/JoaoGeraldoS/API-biblioteca/app/config"
	"github.com/JoaoGeraldoS/API-biblioteca/app/handler"
)

func main() {
	db := config.Conection()

	os.MkdirAll("./uploads/capas", os.ModePerm)

	r := handler.Routes(db)
	r.Run(":8080")
}
