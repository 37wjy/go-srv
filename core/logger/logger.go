package logger

import (
	"log"
	"os"
)

var (
	info  = newlog("INFO")
	debug = newlog("DEBUG")
)

func newlog(prefex string) *log.Logger { //TODO
	// logFile, _ := os.OpenFile("./xx.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	// mw := io.MultiWriter(os.Stdout, logFile)
	return log.New(os.Stdout, prefex+": ", log.Ldate|log.Ltime|log.Lmsgprefix)
}

func INFO(msg interface{}) {
	info.Println(msg)
}

func DEBUG(msg interface{}) {
	debug.Println(msg)
}
