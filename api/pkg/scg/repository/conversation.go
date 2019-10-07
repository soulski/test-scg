package repository

type ConversationRepository interface {
	Repository
}

type memoryConversationRepository struct {
	*MemoryRepository
}

func NewConversationRepository() *memoryConversationRepository {
	return &memoryConversationRepository{NewMemoryRepository()}
}
