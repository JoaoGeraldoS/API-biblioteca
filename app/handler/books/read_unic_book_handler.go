package books

import (
	"database/sql"
	"net/http"

	"github.com/JoaoGeraldoS/API-biblioteca/app/controller/books"
	"github.com/gin-gonic/gin"
)

func ReadUnicBookHandler(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		response, err := books.ReadUnicBook(id, db)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"Erro": "Livro não encontrado"})
		}
		ctx.JSON(http.StatusOK, gin.H{"livro": response})
	}
}
