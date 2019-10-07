package model

type Id interface {
	GetId() interface{}
	Equal(id Id) bool
	String() string
}
