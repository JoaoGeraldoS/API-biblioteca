package users

import (
	"database/sql"
	"net/http"

	"github.com/JoaoGeraldoS/API-biblioteca/app/controller/users"
	"github.com/JoaoGeraldoS/API-biblioteca/app/models"
	"github.com/JoaoGeraldoS/API-biblioteca/app/validacao"
	"github.com/gin-gonic/gin"
)

// Ler as e retorna os dados do banco
// @Summary Ler os Usuarios
// @Description Ler e restorna os dados dos usuarios com  filtros
// @Tags Usuarios
// @Accept json
// @Produce json
// @Query id "Retorno de um usuario"
// @Query name "Busca pelo nome"
// @Success 200 {object} validacao.GenericResponse[models.Users] "Execultada com sucesso"
// @Failure 404 {string} validacao.ErrorResponse "Dados não existentes"
// @Router /public/books [get]
func ReadUsersHandler(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id := ctx.Query("id")
		name := ctx.Query("name")

		response, err := users.ReadUsers(id, name, db)
		if err != nil {
			ctx.JSON(http.StatusNotFound, validacao.ErrorResponse{Message: "usuario não encontrdo"})
			return
		}

		ctx.JSON(http.StatusOK, validacao.GenericResponse[models.Users]{Items: response})
	}
}
