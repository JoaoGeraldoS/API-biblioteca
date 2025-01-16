package books

// Leitura dos livros, retorna uma lista de todos os livors do sistema

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/JoaoGeraldoS/API-biblioteca/app/models"
)

func ReadBook(id string, page int, db *sql.DB) ([]models.Books, error) {
	booksMap := make(map[int64]*models.Books)

	size := 10 // Tamanho padrão de dados por paginas

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * size // Calculo das paginas

	query := `select b.id, b.title, b.author, b.description, b.content,
		DATE_FORMAT(b.created_at, '%d/%m/%y %H:%i:%s') AS created_at, DATE_FORMAT(b.update_at, '%d/%m/%y %H:%i:%s') AS updated_at,
		c.id, c.name, DATE_FORMAT(c.created_at, '%d/%m/%y %H:%i:%s') AS created_at, ifnull(b.img,"Imagem não informada!") as imagem
		from intermediaria i
		join books b on i.book_id = b.id
		join categories c on i.category_id = c.id`

	condicions := []string{}  // Array que salva os dados do friltro
	params := []interface{}{} // Array que salva as querys dos filtros

	if id != "" { // Verfica se foi passado um id para buscas com filtros
		condicions = append(condicions, "b.id = ?")
		params = append(params, id)
	}

	if len(condicions) > 0 { // Verfica se exite algum dados no arry, se existir ele faz a junção na query
		query += " where " + strings.Join(condicions, " ")
	}

	query += " limit ? offset ?"
	params = append(params, size, offset)

	rows, err := db.Query(query, params...)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar dados! %v", err)
	}
	defer rows.Close()

	for rows.Next() { // Faz a iteração dos dados recebidos pela query
		var ( // Variaveis usadas para cada item da struct
			bookID, categoryID                                               int64
			title, author, description, content, created_at, updated_at, img string
			categoryName, created_at_c                                       string
		)

		err := rows.Scan(&bookID, &title, &author, &description, &content, &created_at, &updated_at, &categoryID, &categoryName, &created_at_c, &img)
		if err != nil {
			return nil, fmt.Errorf("erro ao verificar dados! %v", err)
		}

		if _, exists := booksMap[bookID]; !exists { // Verfica se existe o livro no map, essa verificação é feita pelo id
			booksMap[bookID] = &models.Books{ // Se não existir e adicionado no arry do map a struct com os valores mapeados
				ID:          bookID,
				Title:       title,
				Author:      author,
				Description: description,
				Content:     content,
				Created_at:  created_at,
				Updated_at:  updated_at,
				Img:         img,
				Categories:  []models.Categories{},
			}
		}

		booksMap[bookID].Categories = append(booksMap[bookID].Categories, models.Categories{ // Adiciona a categoria no livro dentro do map
			ID:           categoryID,
			Name:         categoryName,
			Created_at_c: created_at_c,
		})
	}

	var books []models.Books
	for _, book := range booksMap { // Faz a iteração do map e adiciona em books, um map para struct
		books = append(books, *book)
	}

	return books, nil
}
