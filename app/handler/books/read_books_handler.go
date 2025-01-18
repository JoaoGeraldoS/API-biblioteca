package books

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/JoaoGeraldoS/API-biblioteca/app/controller/books"
	"github.com/gin-gonic/gin"
)

func ReadBookHandler(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		page := ctx.Param("page")

		var covPage string
		if len(page) > 1 {
			covPage = page[5:]
		} else {
			covPage = page
		}

		intPage, err := strconv.Atoi(covPage)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro na gerção de pages"})
			return
		}

		getid := ctx.Query("id")
		title := ctx.Query("title")
		author := ctx.Query("author")

		response, err := books.ReadBook(getid, title, author, intPage, db)
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusNotFound, gin.H{"error": "livros não encontrados!"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"livros": response})
	}
}
