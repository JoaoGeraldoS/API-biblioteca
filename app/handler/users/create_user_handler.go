package users

import (
	"database/sql"
	"net/http"

	"github.com/JoaoGeraldoS/API-biblioteca/app/controller/users"
	"github.com/JoaoGeraldoS/API-biblioteca/app/models"
	"github.com/gin-gonic/gin"
)

func CerateUserHandler(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user models.Users

		if err := ctx.ShouldBindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao inserir dados"})
			return
		}

		response, err := users.CreateUsers(db, &user)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"erro": "erro interno!"})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{"user": response})
	}
}
