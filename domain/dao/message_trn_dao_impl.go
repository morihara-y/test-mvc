package dao

import (
	"github.com/morihara-y/test-mvc/domain/model"
	"github.com/morihara-y/test-mvc/infrastracuture"
)

type messageTrnDaoImpl struct {
	db infrastracuture.DBConnection
}

// NewMessageTrnDaoImpl returns MessageTrnDao
func NewMessageTrnDaoImpl(db infrastracuture.DBConnection) MessageTrnDao {
	return &messageTrnDaoImpl{
		db: db,
	}
}

func (m *messageTrnDaoImpl) FindAll() []model.Message {
	dbcon := m.db.MakeDBConnection()
	db := dbcon.Connection
	defer db.Close()
	messages := []model.Message{}
	db.Find(&messages)
	return messages
}

func (m *messageTrnDaoImpl) Create(message *model.Message) int {
	dbcon := m.db.MakeDBConnection()
	db := dbcon.Connection
	defer db.Close()
	if !db.NewRecord(&message) {
		panic("could not create new record")
	}
	return 1
}

func (m *messageTrnDaoImpl) DeleteAll() int {
	dbcon := m.db.MakeDBConnection()
	db := dbcon.Connection
	defer db.Close()
	db.Exec("TRUNCATE TABLE message_trn")
	return 1
}
