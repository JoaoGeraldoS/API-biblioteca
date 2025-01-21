package handler

// Rotas do endpoints

import (
	"database/sql"
	"time"

	"github.com/JoaoGeraldoS/API-biblioteca/app/handler/books"
	"github.com/JoaoGeraldoS/API-biblioteca/app/handler/categories"
	"github.com/JoaoGeraldoS/API-biblioteca/app/handler/users"
	"github.com/JoaoGeraldoS/API-biblioteca/app/middleware"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func Routes(db *sql.DB) *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://127.0.0.1:5500", "http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.Static("/uploads", "./uploads")

	public := r.Group("/public")
	{
		public.GET("/books", books.ReadBookHandler(db))
	}

	userComun := r.Group("/users")
	userComun.Use(middleware.AuthMiddleware())
	{
		userComun.GET("/users", users.ReadUsersHandler(db))
		userComun.PUT("/users/:id", users.UpdateUserHandler(db))
	}

	admin := r.Group("/admin")
	admin.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
	{
		// Rotas livros
		admin.POST("/books", books.CreateBookHandler(db))
		admin.GET("/books", books.ReadBookHandler(db))
		admin.PUT("/books/:id", books.UpdateBookHandler(db))
		admin.DELETE("/books/:id", books.DeleteBookHandler(db))

		// Rota categorias
		admin.POST("/categories", categories.CreateCategoryHandler(db))
		admin.GET("/categories", categories.ReadCategoriesHandler(db))
		admin.DELETE("/categories/:id", categories.DeleteCategoryHandler(db))

		// Rota Usuarios
		r.GET("/users", users.ReadUsersHandler(db))
	}

	r.POST("/users", users.CerateUserHandler(db))
	r.POST("/login", users.LoginHandler(db))

	return r
}
