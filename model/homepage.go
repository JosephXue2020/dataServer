package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Notice struct {
	Text string
	Url  string
	gorm.Model
}
