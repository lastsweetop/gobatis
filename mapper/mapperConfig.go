//go:generate gobatis
package mapper

import "gobatis/model"

type User struct {
	GetAll     func() []model.User                `batis:"select id,account,nickName from user"`
	GetUser    func(*model.UserParam) *model.User `batis:"select id,account,nickName from user where id=?"`
	DeleteUser func(*model.UserParam) error       `batis:"delete from user where id=?"`
	AddUser    func(param *model.UserParam) error `batis:"insert user (account,nickName) values (?,?)"`
}

type Device struct {
}
