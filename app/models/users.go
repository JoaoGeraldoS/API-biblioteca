package models

type Users struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Bio        string `json:"bio"`
	Photo      string `json:"photo"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}
