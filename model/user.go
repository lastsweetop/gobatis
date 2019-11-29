package model

type User struct {
	Id       int    `json:"id"`
	Account  string `json:"account"`
	NickName string `json:"nickName"`
}

type UserParam struct {
	Id       int    `json:"id"`
	Account  string `json:"account"`
	NickName string `json:"nickName"`
}
