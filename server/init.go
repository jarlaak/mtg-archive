package server

import (
	"github.com/op/go-logging"
)

var logger *logging.Logger

func UseLogger(log *logging.Logger) {
	logger = log
}
