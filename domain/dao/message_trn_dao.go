package dao

import (
	"github.com/morihara-y/test-mvc/domain/model"
)

// MessageTrnDao is interface of dao to do message_trn
type MessageTrnDao interface {
	FindAll() []model.Message
	Create(model.Message) int
	DeleteAll() int
}
