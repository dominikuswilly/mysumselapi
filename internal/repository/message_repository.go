package repository

import (
	"mysumselapi/internal/model"
)

type MessageRepository interface {
	GetWelcomeMessage() (*model.Message, error)
}

type messageRepository struct{}

func NewMessageRepository() MessageRepository {
	return &messageRepository{}
}

func (r *messageRepository) GetWelcomeMessage() (*model.Message, error) {
	// Mocking a database or external call
	return &model.Message{
		ID:      "1",
		Content: "Welcome to the Clean Architecture API!",
	}, nil
}
