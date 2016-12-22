package archive

import (
	"github.com/op/go-logging"
	"os"
)

var logger *logging.Logger

func InitializeLogger() {
	log := logging.MustGetLogger("mtg")
	loggerBackend := logging.NewLogBackend(os.Stderr, "mtg-archive: ", 0)
	leveledLogger := logging.AddModuleLevel(loggerBackend)
	leveledLogger.SetLevel(logging.DEBUG, "")
	log.SetBackend(leveledLogger)
	logger = log
}

func GetLogger() *logging.Logger {
	if logger == nil {
		InitializeLogger()
	}
	return logger
}
