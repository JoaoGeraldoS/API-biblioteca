package books

// Leitura dos livros, retorna uma lista de todos os livors do sistema

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/JoaoGeraldoS/API-biblioteca/app/models"
)

func ReadBook(id, title, author, category string, page string, db *sql.DB) ([]models.Books, error) {
	booksMap := make(map[int64]*models.Books)

	size := 10 // Tamanho padrão de dados por paginas

	query := `select b.id, b.title, b.author_id, b.description, b.content,
		DATE_FORMAT(b.created_at, '%d/%m/%y %H:%i:%s') AS created_at, DATE_FORMAT(b.updated_at, '%d/%m/%y %H:%i:%s') AS updated_at,
		c.id, c.name, DATE_FORMAT(c.created_at, '%d/%m/%y %H:%i:%s') AS created_at, ifnull(b.img,"Imagem não informada!") as imagem,
		a.id, a.name, a.description
		from intermediary i
		join books b on i.book_id = b.id
		join categories c on i.category_id = c.id
		join authors a on b.author_id = a.id`

	condicions := []string{}  // Array que salva os dados do friltro
	params := []interface{}{} // Array que salva as querys dos filtros

	if id != "" { // Verfica se foi passado um id para buscas com filtros
		condicions = append(condicions, "b.id = ?")
		params = append(params, id)
	} // Adiciona o id no array de filtros

	if title != "" {
		condicions = append(condicions, "b.title like ?")
		params = append(params, fmt.Sprintf("%s%%", title))
	}

	if author != "" {
		condicions = append(condicions, "a.name like ?")
		params = append(params, fmt.Sprintf("%s%%", author))
	}

	if category != "" {
		condicions = append(condicions, "c.name like ?")
		params = append(params, fmt.Sprintf("%s%%", category))
	}

	if page != "" {
		intPage, err := strconv.Atoi(page)
		if err != nil {
			return nil, fmt.Errorf("erro na gerção de pages")
		}

		if intPage < 1 {
			intPage = 1
		}
		offset := (intPage - 1) * size // Calculo das paginas

		query += " limit ? offset ?"
		params = append(params, size, offset)
	}

	if len(condicions) > 0 { // Verfica se exite algum dados no arry, se existir ele faz a junção na query
		query += " where " + strings.Join(condicions, " ")
	}

	rows, err := db.Query(query, params...)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar dados! %v", err)
	}
	defer rows.Close()

	for rows.Next() { // Faz a iteração dos dados recebidos pela query
		var ( // Variaveis usadas para cada item da struct
			bookID, categoryID, authorID, IdAuthor                   int64
			title, description, content, created_at, updated_at, img string
			categoryName, created_at_c                               string
			authorName, authorDescription                            string
		)

		err := rows.Scan(&bookID, &title, &authorID, &description, &content, &created_at, &updated_at, &categoryID, &categoryName, &created_at_c, &img,
			&IdAuthor, &authorName, &authorDescription)
		if err != nil {
			return nil, fmt.Errorf("erro ao verificar dados! %v", err)
		}

		baseURL := "http://localhost:8080/"
		fullImageURL := baseURL + img

		if _, exists := booksMap[bookID]; !exists { // Verfica se existe o livro no map, essa verificação é feita pelo id
			booksMap[bookID] = &models.Books{ // Se não existir e adicionado no arry do map a struct com os valores mapeados
				ID:          bookID,
				Title:       title,
				AuthorId:    authorID,
				Description: description,
				Content:     content,
				CreatedAt:   created_at,
				UpdatedAt:   updated_at,
				Img:         fullImageURL,
				Categories:  []models.Categories{},
				Authors:     []models.Authors{},
			}
		}

		booksMap[bookID].Categories = append(booksMap[bookID].Categories, models.Categories{ // Adiciona a categoria no livro dentro do map
			ID:         categoryID,
			Name:       categoryName,
			CreatedAtC: created_at_c,
		})

		booksMap[bookID].Authors = append(booksMap[bookID].Authors, models.Authors{
			ID:          IdAuthor,
			Name:        authorName,
			Description: authorDescription,
		})
	}

	var books []models.Books
	for _, book := range booksMap { // Faz a iteração do map e adiciona em books, um map para struct
		books = append(books, *book)
	}

	return books, nil
}
