package books

import (
	"database/sql"
	"net/http"

	"github.com/JoaoGeraldoS/API-biblioteca/app/controller/books"
	"github.com/gin-gonic/gin"
)

func DeleteBookHandler(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		if err := books.DeleteBook(id, db); err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"Erro": "dados não existente"})
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "livro apagado com sucesso!"})
	}
}
