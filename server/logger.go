package server

import (
	"fmt"
	"github.com/op/go-logging"
	"os"
)

// TODO: add logging level
type Logger interface {
	Debug(...interface{})
	Info(...interface{})
	Warn(...interface{})
	Error(...interface{})
	Fatal(...interface{})
	/* NO NOT USE: this is redirected to  ERROR to comply
	   mux/handler RecoverHandler Error mechanism */
	Println(...interface{})
	Print(...interface{})
}

var logger Logger

func UseLogger(log Logger) {
	logger = log
}

// STDIO Logger (TODO: move to its own file when more than one logger is implemented)
type IOLogger struct {
	*logging.Logger
}

func NewIOLogger() IOLogger {
	log := logging.MustGetLogger("mtg")
	loggerBackend := logging.NewLogBackend(os.Stderr, "mtg-archive: ", 0)
	leveledLogger := logging.AddModuleLevel(loggerBackend)
	leveledLogger.SetLevel(logging.DEBUG, "")
	log.SetBackend(leveledLogger)
	return IOLogger{log}
}

func (logger IOLogger) Debug(v ...interface{}) {
	logger.Logger.Debug(fmt.Sprint(v...))
}

func (logger IOLogger) Info(v ...interface{}) {
	logger.Logger.Info(fmt.Sprint(v...))
}

func (logger IOLogger) Warn(v ...interface{}) {
	logger.Logger.Warning(fmt.Sprint(v...))
}

func (logger IOLogger) Error(v ...interface{}) {
	logger.Logger.Error(fmt.Sprint(v...))
}

func (logger IOLogger) Fatal(v ...interface{}) {
	logger.Logger.Fatal(fmt.Sprint(v...))
}

func (logger IOLogger) Println(v ...interface{}) {
	logger.Logger.Error(v)
}

func (logger IOLogger) Print(v ...interface{}) {
	logger.Logger.Error(v)
}
