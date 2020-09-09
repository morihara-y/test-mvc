package service

import (
	"github.com/morihara-y/test-mvc/domain/model"
)

// MessageService is interface to do for message
type MessageService interface {
	FetchAll() []model.Message
	Insert(*model.Message) int
	DeleteAll() int
}
