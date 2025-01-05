package handler

// Rotas do endpoints

import (
	"database/sql"

	"github.com/JoaoGeraldoS/API-biblioteca/app/handler/books"
	"github.com/JoaoGeraldoS/API-biblioteca/app/handler/categories"
	"github.com/JoaoGeraldoS/API-biblioteca/app/handler/users"
	"github.com/gin-gonic/gin"
)

func Routes(db *sql.DB) *gin.Engine {
	r := gin.Default()

	// Rota livros
	r.POST("/books", books.CreateBookHandler(db))
	r.GET("/books", books.ReadBookHandler(db))
	r.GET("/books/:id", books.ReadUnicBookHandler(db))
	r.PUT("/books/:id", books.UpdateBookHandler(db))
	r.DELETE("/books/:id", books.DeleteBookHandler(db))

	// Rota categorias
	r.POST("/categories", categories.CreateCategoryHandler(db))
	r.GET("/categories", categories.ReadCategoriesHandler(db))
	r.DELETE("/categories/:id", categories.DeleteCategoryHandler(db))

	// Rota Usuarios
	r.POST("/users", users.CerateUserHandler(db))

	return r
}
