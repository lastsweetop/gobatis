# gobatis

Golang version of mybatis.

Generate dao code from tag of structure.

# install

`go install github.com/lastsweetop/gobatis`

# generate

`//go:generate gobatis -type=User`

# tag

```golang
type User struct {
	GetAll  func() []model.User                `batis:"select id,account from user"`
	GetUser func(*model.UserParam) *model.User `batis:"select id,account from user where id=?"`
}
```