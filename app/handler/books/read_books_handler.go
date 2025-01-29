package books

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/JoaoGeraldoS/API-biblioteca/app/controller/books"
	"github.com/JoaoGeraldoS/API-biblioteca/app/models"
	"github.com/JoaoGeraldoS/API-biblioteca/app/validacao"
	"github.com/gin-gonic/gin"
)

// Ler as e retorna os dados do banco
// @Summary Ler os livros
// @Description Ler e restorna os dados dos livros com categorias e autores, e filtros
// @Tags Livros
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Query page "Paginação"
// @Query id "Retorno de um livro"
// @Query title "Busca pelo titulo"
// @Query author "Busca pelo autor"
// @Query category "Busca por categoria"
// @Success 200 {object} validacao.GenericResponse[models.Books] "Execultada com sucesso"
// @Failure 404 {string} validacao.ErrorResponse "Dados não existentes"
// @Router /public/books [get]
// @Router /admin/books [get]
func ReadBookHandler(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		page := ctx.Query("page")
		getid := ctx.Query("id")
		title := ctx.Query("title")
		author := ctx.Query("author")
		category := ctx.Query("category")

		response, err := books.ReadBook(getid, title, author, category, page, db)

		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusNotFound, validacao.ErrorResponse{Message: "livros não encontrados!"})
			return
		}

		ctx.JSON(http.StatusOK, validacao.GenericResponse[models.Books]{Items: response})
	}
}
