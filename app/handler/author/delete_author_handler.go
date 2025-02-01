package author

import (
	"database/sql"
	"net/http"

	"github.com/JoaoGeraldoS/API-biblioteca/app/controller/author"
	"github.com/JoaoGeraldoS/API-biblioteca/app/validacao"
	"github.com/gin-gonic/gin"
)

// Delete um autor do banco
// @Summary Apaga um autor
// @Description Apaga um autor fornecido pelo id
// @Tags Authors
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path int true "Recebe o id do autor"
// @Success 200 {string} validacao.Response "Execultada com sucesso"
// @Failure 404 {string} validacao.ErrorResponse "Dados não existentes"
// @Router /admin/authors/{id} [delete]
func DeleteAuthorsHandler(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		if err := author.DeleteAuthors(id, db); err != nil {
			ctx.JSON(http.StatusNotFound, validacao.ErrorResponse{Message: "dados não existente"})
			return
		}

		ctx.JSON(http.StatusOK, validacao.Response{Message: "autor apagado com sucesso!"})
	}
}
