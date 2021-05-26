package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	DB *gorm.DB
)

func InitSqlite() (err error) {
	DB, err := gorm.Open("sqlite3", "./db.sqlite3")
	if err != nil {
		return
	}
	return DB.DB().Ping()
}
