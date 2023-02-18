package service

import (
	"store/database"
	"store/entity"
	"time"

	log "github.com/sirupsen/logrus"
)

type MessageService struct {
	repo database.MessageRepository
}

func (s *MessageService) Init() {
	s.repo.Init()
}

func (s *MessageService) SaveMessage(content string) (*entity.Message, error) {
	message := entity.Message{
		CreateTime: time.Now(),
		Content:    content,
	}
	err := s.repo.Insert(&message)
	if err != nil {
		log.Errorf(err.Error())
		return nil, err
	}
	return &message, nil
}
