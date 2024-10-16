package model

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
	Password string `json:"password"`
	Created  string `json:"created"`
}
