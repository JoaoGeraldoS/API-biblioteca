package books

import (
	"database/sql"
	"net/http"

	"github.com/JoaoGeraldoS/API-biblioteca/app/controller/books"
	"github.com/JoaoGeraldoS/API-biblioteca/app/models"

	"github.com/gin-gonic/gin"
)

func CreateBookHandler(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var book models.Books

		if err := ctx.ShouldBindJSON(&book); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "erro ao adicionar dados!"})
			return
		}

		var categories []string
		for _, category := range book.Categories {
			categories = append(categories, category.Name)
		}

		respose, err := books.CreateBook(db, &book, categories)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "erro interno do servidor"})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{"message": "Livro criado!", "livro": respose})

	}
}
