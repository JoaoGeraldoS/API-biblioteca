package users

import (
	"database/sql"
	"net/http"

	"github.com/JoaoGeraldoS/API-biblioteca/app/controller/users"
	"github.com/gin-gonic/gin"
)

func ReadUsersHandler(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id := ctx.Query("id")
		name := ctx.Query("name")

		response, err := users.ReadUsers(id, name, db)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "usuario não encontrdo"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"usuarios": response})
	}
}
