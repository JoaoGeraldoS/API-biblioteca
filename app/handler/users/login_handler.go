package users

import (
	"database/sql"
	"net/http"

	"github.com/JoaoGeraldoS/API-biblioteca/app/controller/users"
	"github.com/JoaoGeraldoS/API-biblioteca/app/models"
	"github.com/gin-gonic/gin"
)

func LoginHandler(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var login models.Users
		if err := ctx.ShouldBindJSON(&login); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
			return
		}

		token, err := users.Login(db, login.Username, login.Password)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"token": token})
	}
}
