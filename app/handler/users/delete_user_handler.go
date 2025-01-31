package users

import (
	"database/sql"
	"net/http"

	"github.com/JoaoGeraldoS/API-biblioteca/app/controller/users"
	"github.com/JoaoGeraldoS/API-biblioteca/app/validacao"
	"github.com/gin-gonic/gin"
)

// Delete um usuario do banco
// @Summary Apaga um usuario
// @Description Apaga um usuario fornecido pelo id
// @Tags Usuarios
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path int true "Recebe o id do usuario"
// @Success 200 {string} validacao.Response "Execultada com sucesso"
// @Failure 404 {string} validacao.ErrorResponse "Dados não existentes"
// @Router /admin/users/{id} [delete]
func DeleteUserHandler(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		if err := users.DeleteUser(id, db); err != nil {
			ctx.JSON(http.StatusNotFound, validacao.ErrorResponse{Message: "erro usuario não encontrado!"})
		}

		ctx.JSON(http.StatusOK, validacao.Response{Message: "usuario apagado com sucesso!"})
	}
}
