package model

import "github.com/jinzhu/gorm"

// DBInfo is db connection info
type DBInfo struct {
	Host       string
	Username   string
	Password   string
	DBName     string
	Port       string
	Connection *gorm.DB
}
