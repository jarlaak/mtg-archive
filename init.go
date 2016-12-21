package archive

import (
        "os"
        "github.com/op/go-logging"
)

func InitializeLogger() *logging.Logger {
    log := logging.MustGetLogger("mtg")
    loggerBackend := logging.NewLogBackend(os.Stderr,"mtg-archive: ",0)
    leveledLogger := logging.AddModuleLevel(loggerBackend)
    leveledLogger.SetLevel(logging.DEBUG,"")
    log.SetBackend(leveledLogger)
    return log
}

