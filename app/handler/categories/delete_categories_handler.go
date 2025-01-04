package categories

import (
	"database/sql"
	"net/http"

	"github.com/JoaoGeraldoS/API-biblioteca/app/controller/categories"
	"github.com/gin-gonic/gin"
)

func DeleteCategoryHandler(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		if err := categories.DeleteCategory(id, db); err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"erro": "id não encontrado!"})
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "categoria apagada!"})

	}
}
