package server

import ()

func GetLogger() Logger {
	return NewIOLogger()
}
