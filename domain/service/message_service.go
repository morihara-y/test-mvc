package service

import (
	"github.com/morihara-y/test-mvc/domain/model"
)

// MessageService is interface to do for message
type MessageService interface {
	FetchAll() ([]model.Message, error)
	Insert(model.Message) (int, error)
	DeleteAll() (int, error)
}
