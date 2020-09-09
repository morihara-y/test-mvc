package dao

import (
	"log"

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

func (m *messageTrnDaoImpl) FindAll() []model.MessageTrn {
	dbcon := m.db.MakeDBConnection()
	db := dbcon.Connection
	defer db.Close()
	messages := []model.MessageTrn{}
	db.Table("message_trn").Find(&messages)
	return messages
}

func (m *messageTrnDaoImpl) Create(message *model.MessageTrn) int {
	dbcon := m.db.MakeDBConnection()
	db := dbcon.Connection
	defer db.Close()
	if !db.Table("message_trn").NewRecord(&message) {
		log.Fatal("it is not new recoad")
		return 0
	}
	result := db.Table("message_trn").Create(&message)
	if result.Error != nil {
		log.Fatal(result.Error)
		log.Fatal("could not create new record")
		return 0
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
