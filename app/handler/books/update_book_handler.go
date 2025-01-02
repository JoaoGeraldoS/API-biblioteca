package books

import (
	"database/sql"
	"net/http"

	"github.com/JoaoGeraldoS/API-biblioteca/app/controller/books"
	"github.com/JoaoGeraldoS/API-biblioteca/app/models"
	"github.com/gin-gonic/gin"
)

func UpdateBookHandler(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		var book models.Books

		if err := ctx.ShouldBindJSON(&book); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "erro ao atualizar dados"})
			return
		}

		response, err := books.UpdateBook(id, db, &book)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "erro interno do servidor!"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "Livro atualizado", "livro": response})
	}
}
