package main

import (
	"github.com/JoaoGeraldoS/API-biblioteca/app/config"
	"github.com/JoaoGeraldoS/API-biblioteca/app/handler"
)

func main() {
	db := config.Conection()

	r := handler.Routes(db)
	r.Run(":8080")
}
