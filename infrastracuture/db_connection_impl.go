package infrastracuture

import (
	"github.com/jinzhu/configor"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/morihara-y/test-mvc/domain/model"
)

type dbConnectionImpl struct {
	dbconfigPath string
}

// NewDBConnectionImpl returns DBConnection
func NewDBConnectionImpl(dbconfigPath string) DBConnection {
	return &dbConnectionImpl{
		dbconfigPath: dbconfigPath,
	}
}

// Config meta
var Config = struct {
	DBInfo struct {
		Host     string
		Username string `default:"root"`
		Password string `required:"true" env:"root"`
		DBName   string
		Port     string `default:"3306"`
	}
}{}

// MakeDBConnection creates DB meta from config yml
func (d *dbConnectionImpl) MakeDBConnection() *model.DBInfo {
	configor.Load(&Config, d.dbconfigPath)
	return newDB(&model.DBInfo{
		Host:     Config.DBInfo.Host,
		Username: Config.DBInfo.Username,
		Password: Config.DBInfo.Password,
		DBName:   Config.DBInfo.DBName,
		Port:     Config.DBInfo.Port,
	})
}

func newDB(d *model.DBInfo) *model.DBInfo {
	db, err := gorm.Open("mysql", d.Username+":"+d.Password+"@tcp("+d.Host+":"+d.Port+")/"+d.DBName+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	d.Connection = db
	return d
}
