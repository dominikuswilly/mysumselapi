package service

import (
	"mysumselapi/internal/model"
	"mysumselapi/internal/repository"
)

type MessageService interface {
	GetWelcome() (*model.Message, error)
}

type messageService struct {
	repo repository.MessageRepository
}

func NewMessageService(repo repository.MessageRepository) MessageService {
	return &messageService{repo: repo}
}

func (s *messageService) GetWelcome() (*model.Message, error) {
	return s.repo.GetWelcomeMessage()
}
