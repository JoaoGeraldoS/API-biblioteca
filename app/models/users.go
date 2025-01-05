package models

type Users struct {
	ID         int64  `json:"id"`
	Name       string `json:"name" binding:"required,min=3"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required,min=6,max=12"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}
