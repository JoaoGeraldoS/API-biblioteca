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

		respose, err := books.CreateBook(db, book)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "erro interno do servidor"})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{"message": "Livro criado!", "livro": respose})

	}
}
