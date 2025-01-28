package categories

import (
	"database/sql"
	"net/http"

	"github.com/JoaoGeraldoS/API-biblioteca/app/controller/categories"
	"github.com/JoaoGeraldoS/API-biblioteca/app/validacao"
	"github.com/gin-gonic/gin"
)

// Delete uma categoria do banco
// @Summary Apaga uma categoria livro
// @Description Apaga uma categoria fornecido pelo id
// @Tags Categorias
// @Accept json
// @Produce json
// @Param id path int true "Recebe o id do livro"
// @Success 200 {string} validacao.Response "Execultada com sucesso"
// @Failure 404 {string} validacao.ErrorResponse "Dados não existentes"
// @Router /admin/categories/{id} [delete]
func DeleteCategoryHandler(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		if err := categories.DeleteCategory(id, db); err != nil {
			ctx.JSON(http.StatusNotFound, validacao.ErrorResponse{Message: "id não encontrado!"})
		}
		ctx.JSON(http.StatusOK, validacao.Response{Message: "categoria apagada!"})

	}
}
