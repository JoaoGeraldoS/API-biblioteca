package main

import (
	"os"

	"github.com/JoaoGeraldoS/API-biblioteca/app/config"
	"github.com/JoaoGeraldoS/API-biblioteca/app/handler"

	_ "github.com/JoaoGeraldoS/API-biblioteca/app/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title API com JWT
// @version 1.0
// @description API com autenticação JWT e Swagger
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	db := config.Conection()

	os.MkdirAll("./uploads/capas", os.ModePerm)

	r := handler.Routes(db)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}
