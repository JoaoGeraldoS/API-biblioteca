package users

import (
	"database/sql"
	"net/http"

	"github.com/JoaoGeraldoS/API-biblioteca/app/controller/users"
	"github.com/JoaoGeraldoS/API-biblioteca/app/models"
	"github.com/JoaoGeraldoS/API-biblioteca/app/validacao"
	"github.com/gin-gonic/gin"
)

// CreateUserHandler cria um novo usuario no banco de dados
// @Summary Cria um novo Usuario
// @Description Cria um novo usuario com os dados fornecidos, name, email, senha, username
// @Tags Usuarios
// @Accept json
// @Produce json
// @Success 201 {object} validacao.ResponseGeneric[models.Users] "Usuario criado com sucesso"
// @Failure 400 {string} validacao.ErrorResponse "Erro nos dados fornecidos"
// @Failure 500 {string} validacao.ErrorResponse "Erro interno do servidor"
// @Router /users [post]
func CerateUserHandler(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user models.Users

		if err := ctx.ShouldBindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, validacao.ErrorResponse{Message: "Erro ao inserir dados"})
			return
		}

		response, err := users.CreateUsers(db, &user)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, validacao.ErrorResponse{Message: "erro interno!"})
			return
		}

		ctx.JSON(http.StatusCreated, validacao.ResponseGeneric[models.Users]{Items: *response})
	}
}
