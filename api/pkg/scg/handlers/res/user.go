package res

import (
	"github.com/soulski/test-scg/pkg/scg/model"
)

type User struct {
	Id string
}

func NewUser(userModel *model.User) *User {
	return &User{
		Id: userModel.GetId().String(),
	}
}
