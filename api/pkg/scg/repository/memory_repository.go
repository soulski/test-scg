package repository

import (
	"errors"
	"fmt"

	model "github.com/soulski/test-scg/pkg/scg/model"
)

type MemoryRepository struct {
	storage map[string]model.Model
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		storage: make(map[string]model.Model),
	}
}

func (mr *MemoryRepository) Create(model model.Model) error {
	keyStr := model.GetId().String()

	_, ok := mr.storage[keyStr]
	if ok {
		return fmt.Errorf("Id %s already exists", model.GetId().String())
	}

	mr.storage[keyStr] = model

	fmt.Printf("Create %v+\n", mr.storage)

	return nil
}

func (mr *MemoryRepository) Update(model model.Model) error {
	keyStr := model.GetId().String()

	_, ok := mr.storage[keyStr]
	if !ok {
		return fmt.Errorf("Id %s is not exists", model.GetId().String())
	}

	mr.storage[keyStr] = model

	return nil
}

func (mr *MemoryRepository) Delete(id model.Id) error {
	keyStr := id.String()

	_, ok := mr.storage[keyStr]
	if !ok {
		return fmt.Errorf("Id %s is not exists", id.String())
	}

	delete(mr.storage, keyStr)

	return nil
}

func (mr *MemoryRepository) FindById(id model.Id) (model.Model, error) {
	keyStr := id.String()

	model, ok := mr.storage[keyStr]

	if ok {
		return model, nil
	} else {
		return nil, errors.New("Not Found")
	}
}

func (mr *MemoryRepository) FindAll() []model.Model {
	models := make([]model.Model, 0, len(mr.storage))

	for _, model := range mr.storage {
		models = append(models, model)
	}

	return models
}
