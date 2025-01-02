package models

// Reponsavel pelo instancia das tabelas no banco de dados

type Books struct {
	ID         int64  `json:"id"`
	Title      string `json:"title" binding:"required,min=3"`
	Author     string `json:"author" binding:"required,min=2"`
	Descrition string `json:"description"`
	Content    string `json:"content"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}
