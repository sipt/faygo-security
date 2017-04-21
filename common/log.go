package common

import (
	"fmt"
	"time"
)

func init() {
	logger = new(Console)
}

//Console console log
type Console struct{}

//Info log info
func (Console) Info(args ...interface{}) {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), " [INFO] ", args)
}

//Warn log info
func (Console) Warn(args ...interface{}) {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), " [WARN] ", args)
}

//Error log info
func (Console) Error(args ...interface{}) {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), " [ERROR] ", args)
}

//ILogger logger interface
type ILogger interface {
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
}

//logger sigle logger entry
var logger ILogger

//Info log info
func Info(args ...interface{}) {
	logger.Info(args)
}

//Warn log info
func Warn(args ...interface{}) {
	logger.Info(args)
}

//Error log info
func Error(args ...interface{}) {
	logger.Error(args)
}

//SetLogger custom logger
func SetLogger(l ILogger) {
	logger = l
}
