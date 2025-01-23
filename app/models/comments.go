package models

type Comments struct {
	ID        int64  `json:"id"`
	Content   string `json:"content"`
	CreatedAT string `json:"created_at"`
	UpdatedAT string `json:"upated_at"`
	User_id   int64  `json:"user_id"`
	Book_id   int64  `jsom:"book_id"`
}
