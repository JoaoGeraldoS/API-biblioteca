package handler

import (
	"database/sql"

	"github.com/JoaoGeraldoS/API-biblioteca/app/handler/books"
	"github.com/gin-gonic/gin"
)

func Routes(db *sql.DB) *gin.Engine {
	r := gin.Default()

	r.POST("/books", books.CreateBookHandler(db))
	r.GET("/books", books.ReadBookHandler(db))

	return r
}
