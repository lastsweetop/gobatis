package model

type User struct {
	Id      int    `json:"id"`
	Account string `json:"account"`
}

type UserParam struct {
	Id int `json:"id"`
}
