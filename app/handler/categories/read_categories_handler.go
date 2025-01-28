package categories

import (
	"database/sql"
	"net/http"

	"github.com/JoaoGeraldoS/API-biblioteca/app/controller/categories"
	"github.com/JoaoGeraldoS/API-biblioteca/app/models"
	"github.com/JoaoGeraldoS/API-biblioteca/app/validacao"
	"github.com/gin-gonic/gin"
)

// Ler as e retorna os dados do banco
// @Summary Ler as categorias
// @Description Ler e restorna os dados das categorias
// @Tags Categorias
// @Accept json
// @Produce json
// @Success 200 {object} validacao.GenericResponse[models.Categories] "Execultada com sucesso"
// @Failure 404 {string} validacao.ErrorResponse "Dados não existentes"
// @Router /admin/categories [get]
func ReadCategoriesHandler(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		response, err := categories.ReadCategories(db)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"erro": "dados nao encontrados!"})
		}

		ctx.JSON(http.StatusOK, validacao.GenericResponse[models.Categories]{Items: response})

	}
}
