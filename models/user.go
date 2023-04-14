package model

type User struct{
	ID string `json:"_id"`
	Email string `json:"email"`
	Password string `json:"password"`
	Posts []string `json:"posts"`
	Comments []string `json:"comments"`
}