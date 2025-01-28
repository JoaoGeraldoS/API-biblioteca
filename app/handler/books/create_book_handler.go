package books

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/JoaoGeraldoS/API-biblioteca/app/controller/books"
	"github.com/JoaoGeraldoS/API-biblioteca/app/models"
	"github.com/JoaoGeraldoS/API-biblioteca/app/validacao"

	"github.com/gin-gonic/gin"
)

// CreateBookHandler cria um novo livro no banco de dados
// @Summary Cria um novo livro
// @Description Cria um novo livro com os dados fornecidos, incluindo título, descrição, conteúdo, autor, categorias e imagem.
// @Tags Livros
// @Accept multipart/form-data
// @Produce json
// @Param title formData string true "Título do Livro"
// @Param description formData string true "Descrição do Livro"
// @Param content formData string true "Conteúdo do Livro"
// @Param author_id formData int true "ID do Autor"
// @Param categories formData []string false "Categorias do Livro"
// @Param image formData file true "Imagem do Livro"
// @Success 201 {object} validacao.ResponseGeneric[models.Books] "Livro criado com sucesso"
// @Failure 400 {string} validacao.ErrorResponse "Erro nos dados fornecidos"
// @Failure 500 {string} validacao.ErrorResponse "Erro interno do servidor"
// @Router /admin/books [post]
func CreateBookHandler(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		title := ctx.PostForm("title")
		description := ctx.PostForm("description")
		content := ctx.PostForm("content")
		authorIDStr := ctx.PostForm("author_id")
		categories := ctx.PostFormArray("categories[]")

		// Converta o authorId para int64
		authorID, err := strconv.ParseInt(authorIDStr, 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, validacao.ErrorResponse{Message: "ID do autor inválido"})
			return
		}

		book := models.Books{
			Title:       title,
			Description: description,
			Content:     content,
			AuthorId:    authorID,
		}

		file, err := ctx.FormFile("image")
		if err != nil {
			ctx.JSON(http.StatusBadRequest, validacao.ErrorResponse{Message: "erro ao receber a imagem"})
			return
		}

		uploadPath := "./uploads/capas"
		err = os.MkdirAll(uploadPath, os.ModePerm)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, validacao.ErrorResponse{Message: "Erro ao criar diretório para salvar imagem"})
			return
		}

		nomeImg := filepath.Base(file.Filename)
		caminhoDestino := filepath.Join(uploadPath, nomeImg)

		err = ctx.SaveUploadedFile(file, caminhoDestino)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, validacao.ErrorResponse{Message: fmt.Sprintf("Erro ao copiar imagem: %v", err)})
			return
		}

		respose, err := books.CreateBook(db, &book, categories, caminhoDestino)
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, validacao.ErrorResponse{Message: "erro interno do servidor"})
			return
		}

		ctx.JSON(http.StatusCreated, validacao.ResponseGeneric[models.Books]{Items: *respose})
	}
}
