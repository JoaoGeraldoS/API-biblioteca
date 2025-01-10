package models

// Reponsavel pelo instancia das tabelas no banco de dados

type Books struct {
	ID          int64        `json:"id"`
	Title       string       `json:"title" binding:"required,min=3"`
	Author      string       `json:"author" binding:"required,min=2"`
	Description string       `json:"description"`
	Content     string       `json:"content"`
	Created_at  string       `json:"created_at"`
	Updated_at  string       `json:"updated_at"`
	Img         string       `json:"img"`
	Categories  []Categories `json:"categories"`
}

type Intermediaria struct {
	ID         int64 `json:"id,omitempty"`
	BookID     int64 `json:"book_id"`
	CategoryID int64 `json:"category_id"`
}

type Categories struct {
	ID           int64  `json:"id"`
	Name         string `json:"name" binding:"required,min=3"`
	Created_at_c string `json:"created_at"`
}
