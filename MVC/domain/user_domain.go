package domain

type User struct {
 Id int `json:"user_id"`
 Firstname string
Lastname string
Email string
}

type Query struct {
	Id   int    `form:"id"`
	Name string `form:"name"`
}