package books

import (
	"database/sql"
	"net/http"

	"github.com/JoaoGeraldoS/API-biblioteca/app/controller/books"
	"github.com/gin-gonic/gin"
)

func ReadBookHandler(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		response, err := books.ReadBook(db)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "livros não encontrados!"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"livros": response})
	}
}
