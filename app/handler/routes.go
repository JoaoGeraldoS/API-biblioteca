package handler

// Rotas do endpoints

import (
	"database/sql"

	"github.com/JoaoGeraldoS/API-biblioteca/app/handler/books"
	"github.com/gin-gonic/gin"
)

func Routes(db *sql.DB) *gin.Engine {
	r := gin.Default()

	r.POST("/books", books.CreateBookHandler(db))
	r.GET("/books", books.ReadBookHandler(db))
	r.PUT("/books/:id", books.UpdateBookHandler(db))
	r.DELETE("/books/:id", books.DeleteBookHandler(db))

	return r
}
