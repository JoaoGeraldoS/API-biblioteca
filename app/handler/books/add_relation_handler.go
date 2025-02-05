package books

import (
	"database/sql"
	"net/http"

	"github.com/JoaoGeraldoS/API-biblioteca/app/controller/books"
	"github.com/JoaoGeraldoS/API-biblioteca/app/models"
	"github.com/JoaoGeraldoS/API-biblioteca/app/validacao"
	"github.com/gin-gonic/gin"
)

// RelationHandler Relaciona livro a categoria
// @Summary Relaciona livro a categoria
// @Description Relaciona livro a categoria com os dados fornecidos, incluindo id do livro, id da categoria
// @Tags Livros
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param ids body models.Intermediary true "Recebe o id do livro e o id da categoria"
// @Success 201 {object} validacao.ResponseGeneric[models.Intermediary] "Relaciona livro a categoria"
// @Failure 400 {string} validacao.ErrorResponse "Erro nos dados fornecidos"
// @Failure 500 {string} validacao.ErrorResponse "Erro interno do servidor"
// @Router /admin/relation [post]
func RelationHandler(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var relation models.Intermediary

		if err := ctx.ShouldBindJSON(&relation); err != nil {
			ctx.JSON(http.StatusBadRequest, validacao.ErrorResponse{Message: "erro! dados invalidos!"})
			return
		}

		response, err := books.Relation(db, relation)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, validacao.ErrorResponse{Message: "erro interno do servidor!"})
			return
		}

		ctx.JSON(http.StatusOK, validacao.ResponseGeneric[models.Intermediary]{Items: *response})
	}
}
