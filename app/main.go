package main

import (
	"os"

	"github.com/JoaoGeraldoS/API-biblioteca/app/config"
	"github.com/JoaoGeraldoS/API-biblioteca/app/handler"

	_ "github.com/JoaoGeraldoS/API-biblioteca/app/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	db := config.Conection()

	os.MkdirAll("./uploads/capas", os.ModePerm)

	r := handler.Routes(db)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}
