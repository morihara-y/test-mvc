package service

import (
	"github.com/morihara-y/test-mvc/domain/dao"
	"github.com/morihara-y/test-mvc/domain/model"
)

type messageServiceImpl struct {
	dao dao.MessageTrnDao
}

// NewMessageServiceImpl returns MessageTrnDao
func NewMessageServiceImpl(dao dao.MessageTrnDao) MessageService {
	return &messageServiceImpl{
		dao: dao,
	}
}

func (m *messageServiceImpl) FetchAll() []model.MessageTrn {
	return m.dao.FindAll()
}

func (m *messageServiceImpl) Insert(message *model.MessageTrn) int {
	return m.dao.Create(message)
}

func (m *messageServiceImpl) DeleteAll() int {
	return m.dao.DeleteAll()
}
