# gobatis

Golang version of mybatis.

Generate dao code from tag of structure.

So you don't have to worry about the performance issues caused by orm reflection, but you can still get rid of the boring dao layer code

# install

`go install github.com/lastsweetop/gobatis`

# generate

`//go:generate gobatis

# tag

```golang
type User struct {
	GetAll     func() []model.User                `gobatis:"select id,account,nickName from user"`
	GetUser    func(*model.UserParam) *model.User `gobatis:"select id,account,nickName from user where id=?"`
	DeleteUser func(*model.UserParam) error       `gobatis:"delete from user where id=?"`
	AddUser    func(param *model.UserParam) error `gobatis:"insert user (account,nickName) values (?,?)"`
}

type Device struct {
}
```