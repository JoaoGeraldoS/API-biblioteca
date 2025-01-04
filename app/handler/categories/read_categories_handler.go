package categories

import (
	"database/sql"
	"net/http"

	"github.com/JoaoGeraldoS/API-biblioteca/app/controller/categories"
	"github.com/gin-gonic/gin"
)

func ReadCategoriesHandler(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		response, err := categories.ReadCategories(db)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"erro": "dados nao encontrados!"})
		}

		ctx.JSON(http.StatusOK, gin.H{"categories": response})

	}
}
