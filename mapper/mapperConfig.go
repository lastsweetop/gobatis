//go:generate gobatis
package mapper

import "gobatis/model"
import _ "gobatis/db"
import _ "log"

type User struct {
	GetAll     func() []model.User                `gobatis:"select id,account,nickName from user"`
	GetUser    func(*model.UserParam) *model.User `gobatis:"select id,account,nickName from user where id=?"`
	DeleteUser func(*model.UserParam) error       `gobatis:"delete from user where id=?"`
	AddUser    func(param *model.UserParam) error `gobatis:"insert user (account,nickName) values (?,?)"`
	UpdateUser func(param *model.UserParam) error `gobatis:"update user set account=?,nickName=? where id=?"`
}

type Device struct {
}
