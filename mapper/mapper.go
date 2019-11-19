//go:generate gobatis -type=User
package mapper

import "gobatis/model"

type User struct {
	GetAll  func() []model.User                `batis:"select id,account from user"`
	GetUser func(*model.UserParam) *model.User `batis:"select id,account from user where id=?"`
}
