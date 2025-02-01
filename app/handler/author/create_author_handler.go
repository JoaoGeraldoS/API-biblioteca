package author

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/JoaoGeraldoS/API-biblioteca/app/controller/author"
	"github.com/JoaoGeraldoS/API-biblioteca/app/models"
	"github.com/JoaoGeraldoS/API-biblioteca/app/validacao"
	"github.com/gin-gonic/gin"
)

// CreateAuthorHandler cria um novo livro no banco de dados
// @Summary Cria um Autor
// @Description Cria um novo autor com os dados fornecidos, incluindo name, descrição, foto
// @Tags Authors
// @Accept multipart/form-data
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param name formData string true "Nome do autor"
// @Param description formData string true "Descrição do autor"
// @Param photo formData file true "Foto do autor"
// @Success 201 {object} validacao.ResponseGeneric[models.Authors] "Author criado com sucesso"
// @Failure 400 {string} validacao.ErrorResponse "Erro nos dados fornecidos"
// @Failure 500 {string} validacao.ErrorResponse "Erro interno do servidor"
// @Router /admin/author [post]
func CreateAuthorHandler(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		name := ctx.PostForm("name")
		decription := ctx.PostForm("description")

		getAuthor := models.Authors{
			Name:        name,
			Description: decription,
		}

		file, err := ctx.FormFile("photo")
		if err != nil {
			ctx.JSON(http.StatusBadRequest, validacao.ErrorResponse{Message: "erro ao receber a imagem"})
			return
		}

		uploadPath := "./uploads/authors"
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

		response, err := author.CreateAuthor(db, &getAuthor, caminhoDestino)
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, validacao.ErrorResponse{Message: "erro interno do servidor"})
			return
		}

		ctx.JSON(http.StatusCreated, validacao.ResponseGeneric[models.Authors]{Items: *response})

	}
}
