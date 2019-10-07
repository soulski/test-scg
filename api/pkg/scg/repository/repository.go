package repository

import model "github.com/soulski/test-scg/pkg/scg/model"

type Repository interface {
	Create(model model.Model) error
	Update(model model.Model) error
	Delete(id model.Id) error
	FindById(id model.Id) (model.Model, error)
	FindAll() (models []model.Model)
}
