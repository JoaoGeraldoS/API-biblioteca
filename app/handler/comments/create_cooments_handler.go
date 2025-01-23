package comments

import (
	"database/sql"
	"net/http"

	"github.com/JoaoGeraldoS/API-biblioteca/app/controller/comments"
	"github.com/JoaoGeraldoS/API-biblioteca/app/models"
	"github.com/gin-gonic/gin"
)

func CreateCommentsHandler(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var comment models.Comments

		if err := ctx.ShouldBindJSON(&comment); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "erro ao verificar os dados"})
			return
		}

		response, err := comments.CreateComments(db, &comment)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "erro interno do servidor"})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{"comment": response})

	}
}
