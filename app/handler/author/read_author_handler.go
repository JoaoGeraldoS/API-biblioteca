package author

import (
	"database/sql"
	"net/http"

	"github.com/JoaoGeraldoS/API-biblioteca/app/controller/author"
	"github.com/JoaoGeraldoS/API-biblioteca/app/models"
	"github.com/JoaoGeraldoS/API-biblioteca/app/validacao"
	"github.com/gin-gonic/gin"
)

// Ler as e retorna os dados do banco
// @Summary Ler os autores
// @Description Ler e restorna os dados dos autores
// @Tags Authors
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} validacao.GenericResponse[models.Authors] "Execultada com sucesso"
// @Failure 404 {string} validacao.ErrorResponse "Dados não existentes"
// @Router /admin/authors [get]
func ReadAuthorsHandler(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		response, err := author.ReadAuthors(db)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"erro": "dados nao encontrados!"})
			return
		}

		ctx.JSON(http.StatusOK, validacao.GenericResponse[models.Authors]{Items: response})
	}
}
