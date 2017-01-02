package models

import (
	"github.com/jarlaak/mtg-archive/server"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var logger server.Logger

func Init(database *gorm.DB, log server.Logger) {
	db = database
	logger = log
	db.LogMode(true)
}
