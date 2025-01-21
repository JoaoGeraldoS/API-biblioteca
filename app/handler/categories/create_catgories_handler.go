package categories

import (
	"database/sql"
	"net/http"

	"github.com/JoaoGeraldoS/API-biblioteca/app/controller/categories"
	"github.com/JoaoGeraldoS/API-biblioteca/app/models"
	"github.com/gin-gonic/gin"
)

func CreateCategoryHandler(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var category models.Categories

		if err := ctx.ShouldBindJSON(&category); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"erro": "erro ao criar dados"})
			return
		}

		response, err := categories.CreateCategory(db, category)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "erro interno do servidor"})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{"message": "Categoria criada com sucesso!", "categoria": response})
	}
}
