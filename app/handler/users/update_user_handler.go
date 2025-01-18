package users

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/JoaoGeraldoS/API-biblioteca/app/controller/users"
	"github.com/JoaoGeraldoS/API-biblioteca/app/models"
	"github.com/JoaoGeraldoS/API-biblioteca/app/utils"
	"github.com/gin-gonic/gin"
)

func UpdateUserHandler(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		var user models.Users

		if err := ctx.ShouldBindJSON(&user); err != nil {
			log.Print(err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "erro ao atualizar dados"})
			return
		}

		uploadPath := "./uploads/perfil"
		err := os.MkdirAll(uploadPath, os.ModePerm)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar diretório para salvar imagem"})
			return
		}

		namePhoto := filepath.Base(user.Photo)
		caminhoDestino := filepath.Join(uploadPath, namePhoto)

		err = utils.CopyImg(user.Photo, caminhoDestino)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Erro ao copiar imagem: %v", err)})
			return
		}

		response, err := users.UpdateUser(id, db, &user, caminhoDestino)
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "erro interno do servidor!"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"user": response})
	}
}
