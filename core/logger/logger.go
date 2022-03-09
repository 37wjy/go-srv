package logger

import (
	"log"
	"os"
)

var (
	Info  = newlog("INFO")
	Debug = newlog("DEBUG")
)

func newlog(prefex string) *log.Logger { //TODO
	logFile, _ := os.OpenFile("./xx.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	return log.New(logFile, prefex+": ", log.Ldate|log.Ltime|log.Lmsgprefix)
}
