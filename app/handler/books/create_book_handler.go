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

	"github.com/gin-gonic/gin"
)

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
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID do autor inválido"})
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
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "erro ao receber a imagem"})
			return
		}

		uploadPath := "./uploads/capas"
		err = os.MkdirAll(uploadPath, os.ModePerm)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar diretório para salvar imagem"})
			return
		}

		nomeImg := filepath.Base(file.Filename)
		caminhoDestino := filepath.Join(uploadPath, nomeImg)

		err = ctx.SaveUploadedFile(file, caminhoDestino)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Erro ao copiar imagem: %v", err)})
			return
		}

		respose, err := books.CreateBook(db, &book, categories, caminhoDestino)
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "erro interno do servidor"})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{"message": "Livro criado!", "livro": respose})

	}
}
