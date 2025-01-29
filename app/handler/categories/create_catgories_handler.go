package categories

import (
	"database/sql"
	"net/http"

	"github.com/JoaoGeraldoS/API-biblioteca/app/controller/categories"
	"github.com/JoaoGeraldoS/API-biblioteca/app/models"
	"github.com/JoaoGeraldoS/API-biblioteca/app/validacao"
	"github.com/gin-gonic/gin"
)

// CreateCategoryHandler cria uma nova categoria no banco de dados
// @Summary Cria uma categoria
// @Description Cria uma categoria com os dados fornecidos, nome
// @Tags Categorias
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param Categories body models.Categories true "Cria categoria, parametros: Name"
// @Success 201 {object} validacao.ResponseGeneric[models.Categories] "Categoria criada com sucesso"
// @Failure 400 {string} validacao.ErrorResponse "Requisição nao execultada"
// @Failure 500 {string} validacao.ErrorResponse "Erro interno do servidor"
// @Router /admin/categories [post]
func CreateCategoryHandler(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var category models.Categories

		if err := ctx.ShouldBindJSON(&category); err != nil {
			ctx.JSON(http.StatusBadRequest, validacao.ErrorResponse{Message: "erro ao criar dados"})
			return
		}

		response, err := categories.CreateCategory(db, category)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, validacao.ErrorResponse{Message: "erro interno do servidor"})
			return
		}

		ctx.JSON(http.StatusCreated, validacao.ResponseGeneric[models.Categories]{Items: response})
	}
}
