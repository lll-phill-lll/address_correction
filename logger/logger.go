package logger

import (
	"io"
	"log"
	"os"
)

// https://www.ardanlabs.com/blog/2013/11/using-log-package-in-go.html?showComment=1396035887595

var (
	// Trace level of logging
	Debug *log.Logger
	// Info level of logging
	Info *log.Logger
	// Warning level of logging
	Warning *log.Logger
	// Error level of logging
	Error *log.Logger
)

// SetLogger setups logger instance for whole project
func SetLogger(traceHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) {

	file, err := os.OpenFile("Logs.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666) // file for logs
	if err != nil {
		log.Println("Failed to open log file:", err)
	}

	outDebug := io.MultiWriter(file, traceHandle) // copy logs stream to file
	outInfo := io.MultiWriter(file, infoHandle)
	outWarning := io.MultiWriter(file, warningHandle)
	outError := io.MultiWriter(file, errorHandle)

	// set function to write log, example: Info.Println("someText")
	Debug = log.New(outDebug,
		"TRACE: ",
		log.Ldate|log.Lmicroseconds|log.Lshortfile)

	Info = log.New(outInfo,
		"INFO: ",
		log.Ldate|log.Lmicroseconds|log.Lshortfile)

	Warning = log.New(outWarning,
		"WARNING: ",
		log.Ldate|log.Lmicroseconds|log.Lshortfile)

	Error = log.New(outError,
		"ERROR: ",
		log.Ldate|log.Lmicroseconds|log.Lshortfile)
}
