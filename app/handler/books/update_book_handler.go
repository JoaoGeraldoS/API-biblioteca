package books

import (
	"database/sql"
	"net/http"

	"github.com/JoaoGeraldoS/API-biblioteca/app/controller/books"
	"github.com/JoaoGeraldoS/API-biblioteca/app/models"
	"github.com/JoaoGeraldoS/API-biblioteca/app/validacao"
	"github.com/gin-gonic/gin"
)

// Atualiza um livro do banco
// @Summary Atualiza um novo livro
// @Description Atualiza um livro fornecido pelo id
// @Tags Livros
// @Accept json
// @Produce json
// @Param id path int true "Recebe o id do livro"
// @Success 200 {object} validacao.Response "Execultada com sucesso"
// @Failure 400 {string} validacao.ErrorResponse "Requisição nao execultada"
// @Failure 500 {string} validacao.ErrorResponse "Erro no servidor"
// @Router /admin/books/{id} [put]
func UpdateBookHandler(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		var book models.Books

		if err := ctx.ShouldBindJSON(&book); err != nil {
			ctx.JSON(http.StatusBadRequest, validacao.ErrorResponse{Message: "erro ao atualizar dados"})
			return
		}

		response, err := books.UpdateBook(id, db, &book)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, validacao.ErrorResponse{Message: "erro interno do servidor!"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "Livro atualizado", "livro": response})
	}
}
