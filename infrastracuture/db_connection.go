package infrastracuture

import (
	"github.com/morihara-y/test-mvc/domain/model"
)

// DBConnection is interface to make connection
type DBConnection interface {
	MakeDBConnection() *model.DBInfo
}
