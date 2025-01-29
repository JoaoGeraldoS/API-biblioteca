package users

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/JoaoGeraldoS/API-biblioteca/app/controller/users"
	"github.com/JoaoGeraldoS/API-biblioteca/app/models"
	"github.com/JoaoGeraldoS/API-biblioteca/app/validacao"
	"github.com/gin-gonic/gin"
)

// LoginHandler Faz o login do usuario
// @Summary Faz o login do usuario
// @Description Faz o login do usuario
// @Tags Usuarios
// @Accept json
// @Produce json
// @Param login body models.LoginRequest true "Dados do login do usuário"
// @Success 200 {object} validacao.Response "Usuario criado com sucesso"
// @Failure 400 {string} validacao.ErrorResponse "Erro nos dados fornecidos"
// @Failure 500 {string} validacao.ErrorResponse "Erro interno do servidor"
// @Router /login [post]
func LoginHandler(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var login models.LoginRequest
		if err := ctx.ShouldBindJSON(&login); err != nil {
			ctx.JSON(http.StatusBadRequest, validacao.ErrorResponse{Message: "Dados inválidos"})
			return
		}

		token, err := users.Login(db, login.Username, login.Password)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, validacao.ErrorResponse{Message: err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, validacao.Response{Message: fmt.Sprintf("Token: %s", token)})
	}
}
