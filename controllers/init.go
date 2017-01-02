package controllers

import (
	"github.com/jarlaak/mtg-archive/server"
)

var logger server.Logger

func Init(log server.Logger) {
	logger = log
}
