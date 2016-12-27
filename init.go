package archive

import (
	"github.com/jarlaak/mtg-archive/server"
        "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var logger server.Logger

var db *gorm.DB

func InitializeLogger() {
	logger = server.NewIOLogger()
}

func GetLogger() server.Logger {
	if logger == nil {
		InitializeLogger()
	}
	return logger
}


func InitializeDatabase() {
	database, err := gorm.Open("postgres","dbname=mtg-local")
	if err != nil {
		panic(err)
	}
	db = database
}
