package users

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/JoaoGeraldoS/API-biblioteca/app/controller/users"
	"github.com/JoaoGeraldoS/API-biblioteca/app/models"
	"github.com/JoaoGeraldoS/API-biblioteca/app/validacao"
	"github.com/gin-gonic/gin"
)

// UpdateUserHandler Atualiza usuarios
// @Summary Atualiza usuarios
// @Description Atualiza usuarios com os dados fornecidos, name, bio, foto
// @Tags Usuarios
// @Accept multipart/form-data
// @Produce json
// @Param name formData string true "Nome do usuario"
// @Param bio formData string true "Descrição do usuario"
// @Param image formData file true "Foto do usuario"
// @Success 200 {object} validacao.ResponseGeneric[models.Users] "Livro criado com sucesso"
// @Failure 400 {string} validacao.ErrorResponse "Erro nos dados fornecidos"
// @Failure 500 {string} validacao.ErrorResponse "Erro interno do servidor"
// @Router /users/users/{id} [put]
func UpdateUserHandler(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		name := ctx.PostForm("name")
		bio := ctx.PostForm("bio")

		user := models.Users{
			Name: name,
			Bio:  bio,
		}

		file, err := ctx.FormFile("image")
		if err != nil {
			ctx.JSON(http.StatusBadRequest, validacao.ErrorResponse{Message: "erro ao receber a imagem"})
			return
		}

		uploadPath := "./uploads/perfil"
		err = os.MkdirAll(uploadPath, os.ModePerm)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, validacao.ErrorResponse{Message: "Erro ao criar diretório para salvar imagem"})
			return
		}

		namePhoto := filepath.Base(file.Filename)
		caminhoDestino := filepath.Join(uploadPath, namePhoto)

		err = ctx.SaveUploadedFile(file, caminhoDestino)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, validacao.ErrorResponse{Message: fmt.Sprintf("Erro ao copiar imagem: %v", err)})
			return
		}

		response, err := users.UpdateUser(id, db, &user, caminhoDestino)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, validacao.ErrorResponse{Message: "erro interno do servidor"})
			return
		}

		ctx.JSON(http.StatusOK, validacao.ResponseGeneric[models.Users]{Items: *response})
	}
}
