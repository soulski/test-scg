package model

type DefinedId struct {
	id string
}

func NewDefinedId(id string) *DefinedId {
	return &DefinedId{id}
}

func (d *DefinedId) GetId() interface{} {
	return d.id
}

func (d *DefinedId) Equal(id Id) bool {
	return d.id == id.String()
}

func (d *DefinedId) String() string {
	return d.id
}
