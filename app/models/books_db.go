package models

// Reponsavel pelo instancia das tabelas no banco de dados

type Books struct {
	ID          int64        `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Content     string       `json:"content"`
	CreatedAt   string       `json:"created_at"`
	UpdatedAt   string       `json:"updated_at"`
	Img         string       `json:"img"`
	Categories  []Categories `form:"categories" json:"categories"`
	AuthorId    int64        `json:"author_id"`
	Authors     []Authors    `json:"author"`
}

type Intermediary struct {
	ID         int64 `json:"id,omitempty"`
	BookID     int64 `json:"book_id"`
	CategoryID int64 `json:"category_id"`
}

type Categories struct {
	ID         int64  `json:"id"`
	Name       string `json:"name" binding:"required,min=3"`
	CreatedAtC string `json:"created_at"`
}

type Authors struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Photo       string `json:"photo"`
}
