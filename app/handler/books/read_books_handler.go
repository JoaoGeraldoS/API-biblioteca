package books

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/JoaoGeraldoS/API-biblioteca/app/controller/books"
	"github.com/gin-gonic/gin"
)

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
			ctx.JSON(http.StatusNotFound, gin.H{"error": "livros não encontrados!"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"livros": response})
	}
}
