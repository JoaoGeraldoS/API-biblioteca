package books

import (
	"database/sql"
	"net/http"

	"github.com/JoaoGeraldoS/API-biblioteca/app/controller/books"
	"github.com/JoaoGeraldoS/API-biblioteca/app/validacao"
	"github.com/gin-gonic/gin"
)

// Delete um livro do banco
// @Summary Apaga um novo livro
// @Description Apaga um livro fornecido pelo id
// @Tags Livros
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path int true "Recebe o id do livro"
// @Success 200 {string} validacao.Response "Execultada com sucesso"
// @Failure 404 {string} validacao.ErrorResponse "Dados não existentes"
// @Router /admin/books/{id} [delete]
func DeleteBookHandler(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		if err := books.DeleteBook(id, db); err != nil {
			ctx.JSON(http.StatusNotFound, validacao.ErrorResponse{Message: "dados não existente"})
		}

		ctx.JSON(http.StatusOK, validacao.Response{Message: "livro apagado com sucesso!"})
	}
}
