package books

import (
	"database/sql"
	"log"
	"net/http"
	"path/filepath"

	"github.com/JoaoGeraldoS/API-biblioteca/app/controller/books"
	"github.com/JoaoGeraldoS/API-biblioteca/app/models"

	"github.com/gin-gonic/gin"
)

func CreateBookHandler(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var book models.Books

		if err := ctx.ShouldBindJSON(&book); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "dados nao encontrados"})
		}

		// Itera sobre as categorias vindo do json
		var categories []string
		for _, category := range book.Categories {
			categories = append(categories, category.Name)
		}

		file, err := ctx.FormFile("img")
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Capa é obrigatória"})
			return
		}

		uploadPath := "./uploads/capas"
		filePath := filepath.Join(uploadPath, file.Filename)

		if err := ctx.SaveUploadedFile(file, filePath); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar a capa"})
			return
		}

		respose, err := books.CreateBook(db, &book, categories, filePath)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "erro interno do servidor"})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{"message": "Livro criado!", "livro": respose})

	}
}
