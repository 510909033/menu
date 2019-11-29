package applog

import (
	"io"
	"log"
	"os"
)

var (
	LogInfo  *log.Logger
	LogError *log.Logger
)

func init() {

	errFile, err := os.OpenFile("errors.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("打开日志文件失败：", err)
	}

	LogInfo = log.New(io.MultiWriter(os.Stdout, errFile), "Info:", log.Ldate|log.Ltime|log.Lshortfile)
	LogError = log.New(io.MultiWriter(os.Stderr, errFile), "Error:", log.Ldate|log.Ltime|log.Lshortfile)

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
