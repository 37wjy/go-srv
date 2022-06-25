package logger

import (
	"fmt"
	"io"
	"log"
	"os"
)

// private
var (
	info  = newlog("INFO")
	debug = newlog("DEBUG")
	fatal = newlog("FATAL")
)

func newlog(prefex string) *log.Logger { //TODO 实现本地或网络log
	s := fmt.Sprintf("/data/logs/s9998/%s.log", prefex)
	logFile, _ := os.OpenFile(s, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	mw := io.MultiWriter(os.Stdout, logFile)
	return log.New(mw, prefex+": ", log.Ldate|log.Ltime|log.Lmsgprefix)
}

func Info(msg ...interface{}) {
	info.Println(msg...)
}

func Infof(msg string, v ...interface{}) { //有想法改个名字
	info.Printf(msg, v...)
}

func Debug(msg ...interface{}) {
	debug.Println(msg...)
}

func Debugf(msg string, v ...interface{}) {
	debug.Printf(msg, v...)
}

func Fatal(msg ...interface{}) {
	fatal.Println(msg...)
}

func Fatalf(msg string, v ...interface{}) {
	fatal.Printf(msg, v...)
}
