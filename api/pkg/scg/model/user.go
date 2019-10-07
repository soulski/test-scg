package model

type User struct {
	Id Id
}

func NewUser(id string) *User {
	return &User{Id: NewDefinedId(id)}
}

func (u *User) GetId() Id {
	return u.Id
}
