package repository

type UserRepository interface {
	Repository
}

type memoryUserRepository struct {
	*MemoryRepository
}

func NewUserRepository() *memoryUserRepository {
	return &memoryUserRepository{NewMemoryRepository()}
}
