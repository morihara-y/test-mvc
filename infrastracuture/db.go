package infrastracuture

import (
	"github.com/jinzhu/configor"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB connection info
type DB struct {
	Host       string
	Username   string
	Password   string
	DBName     string
	Port       string
	Connection *gorm.DB
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

// NewDB creates DB meta from config yml
func NewDB() *DB {
	configor.Load(&Config, "dbconfig.yml")
	return newDB(&DB{
		Host:     Config.DBInfo.Host,
		Username: Config.DBInfo.Username,
		Password: Config.DBInfo.Password,
		DBName:   Config.DBInfo.DBName,
		Port:     Config.DBInfo.Port,
	})
}

func newDB(d *DB) *DB {
	db, err := gorm.Open("mysql", d.Username+":"+d.Password+"@tcp("+d.Host+":"+d.Port+")/"+d.DBName+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	d.Connection = db
	return d
}

// Begin begins a transaction
func (db *DB) Begin() *gorm.DB {
	return db.Connection.Begin()
}

// Connect returns db connection
func (db *DB) Connect() *gorm.DB {
	return db.Connection
}
