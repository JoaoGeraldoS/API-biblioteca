package books

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/JoaoGeraldoS/API-biblioteca/app/controller/books"
	"github.com/JoaoGeraldoS/API-biblioteca/app/models"
	"github.com/JoaoGeraldoS/API-biblioteca/app/utils"

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

		uploadPath := "./uploads/capas"
		err := os.MkdirAll(uploadPath, os.ModePerm)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar diretório para salvar imagem"})
			return
		}

		nomeImg := filepath.Base(book.Img)
		caminhoDestino := filepath.Join(uploadPath, nomeImg)

		err = utils.CopyImg(book.Img, caminhoDestino)
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
