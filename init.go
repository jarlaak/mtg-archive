package archive

import (
	"github.com/jarlaak/mtg-archive/server"
)

var logger server.Logger

func InitializeLogger() {
	logger = server.NewIOLogger()
}

func GetLogger() server.Logger {
	if logger == nil {
		InitializeLogger()
	}
	return logger
}
