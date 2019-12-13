package applog

import (
	"io"
	"log"
	"os"
)

type LEVEL int

const (
	DEBUG LEVEL = iota
	INFO

	ERROR
)

var (
	LogDebug *log.Logger
	LogInfo  *log.Logger
	LogError *log.Logger
)

func init() {

	errError, err := os.OpenFile("log_errors.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	errInfo, err := os.OpenFile("log_info.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	errDebug, err := os.OpenFile("log_debug.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("打开日志文件失败：", err)
	}

	LogDebug = log.New(io.MultiWriter(os.Stdout, errDebug), "Debug:", log.Ldate|log.Ltime|log.Lshortfile)
	LogInfo = log.New(io.MultiWriter(os.Stdout, errInfo), "Info:", log.Ldate|log.Ltime|log.Lshortfile)
	LogError = log.New(io.MultiWriter(os.Stderr, errError, errDebug, errInfo), "Error:", log.Ldate|log.Ltime|log.Lshortfile)

}

func Log(msg interface{}, level LEVEL) {
	switch level {
	case DEBUG:
		LogDebug.Print(msg)
	case INFO:
		LogInfo.Print(msg)
	case ERROR:
		LogError.Print(msg)
		break
	}
}

/*
		errFile, err := os.OpenFile("errors.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("打开日志文件失败：", err)
	}

	//	Info = log.New(io.MultiWriter(errFile), "Info:", log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(io.MultiWriter(os.Stdout, errFile), "Info:", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(os.Stdout, "Warning:", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(io.MultiWriter(os.Stderr, errFile), "Error:", log.Ldate|log.Ltime|log.Lshortfile)
*/
