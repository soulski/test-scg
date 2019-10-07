package model

import (
	"github.com/google/uuid"
)

type RandomId struct {
	id uuid.UUID
}

func NewRandomId() *RandomId {
	return &RandomId{
		id: uuid.New(),
	}
}

func (d *RandomId) GetId() interface{} {
	return d.id
}

func (d *RandomId) Equal(id Id) bool {
	return d.GetId() == id.GetId()
}

func (d *RandomId) String() string {
	return d.id.String()
}
