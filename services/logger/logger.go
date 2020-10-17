
package logger

import (
	"fmt"
	"log"
	"os"
)

// Levels of Log
const (
	TRACE   = "TRACE"
	DEBUG   = "DEBUG"
	INFO    = "INFO"
	WARNING = "WARNING"
	ERROR   = "ERROR"
)

//Mapping the  log levels to no
var LogLevelsMap = map[string]int{
	"TRACE":   0,
	"DEBUG":   1,
	"INFO":    2,
	"WARNING": 3,
	"ERROR":   4,
}

var logLevel = 2

var (
	traceLogger = log.New(os.Stdout, "[TRACE]: ", log.Ldate|log.Ltime|log.Llongfile)
	debugLogger = log.New(os.Stdout, "[DEBUG]: ", log.Ldate|log.Ltime|log.Llongfile)
	infoLogger  = log.New(os.Stdout, "[INFO]: ", log.Ldate|log.Ltime|log.Lshortfile)
	warnLogger  = log.New(os.Stdout, "[WARNING]: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(os.Stderr, "[ERROR]: ", log.Ldate|log.Ltime|log.Llongfile)
)

//Trace traces
func Trace(v ...interface{}) {
	if logLevel == LogLevelsMap[TRACE] {
		traceLogger.Output(2, fmt.Sprintln(v...))
	}
}

//Tracef trace with formatting
func Tracef(format string, v ...interface{}) {
	if logLevel == LogLevelsMap[TRACE] {
		traceLogger.Output(2, fmt.Sprintf(format, v...))
	}
}

//Debug debug 
func Debug(v ...interface{}) {
	if logLevel >= LogLevelsMap[TRACE] && logLevel <= LogLevelsMap[DEBUG] {
		debugLogger.Output(2, fmt.Sprintln(v...))
	}
}

//Debugf debug with format
func Debugf(format string, v ...interface{}) {
	if logLevel >= LogLevelsMap[TRACE] && logLevel <= LogLevelsMap[DEBUG] {
		debugLogger.Output(2, fmt.Sprintf(format, v...))
	}
}

//Info info 
func Info(v ...interface{}) {
	if logLevel >= LogLevelsMap[TRACE] && logLevel <= LogLevelsMap[INFO] {
		infoLogger.Output(2, fmt.Sprintln(v...))
	}
}

//Infof info and format
func Infof(format string, v ...interface{}) {
	if logLevel >= LogLevelsMap[TRACE] && logLevel <= LogLevelsMap[INFO] {
		infoLogger.Output(2, fmt.Sprintf(format, v...))
	}
}

//Warn warning
func Warn(v ...interface{}) {
	if logLevel >= LogLevelsMap[TRACE] && logLevel <= LogLevelsMap[WARNING] {
		warnLogger.Output(2, fmt.Sprintln(v...))
	}
}

//Warnf warning and format
func Warnf(format string, v ...interface{}) {
	if logLevel >= LogLevelsMap[TRACE] && logLevel <= LogLevelsMap[WARNING] {
		warnLogger.Output(2, fmt.Sprintf(format, v...))
	}
}

//Error error
func Error(v ...interface{}) {
	if logLevel >= LogLevelsMap[TRACE] && logLevel <= LogLevelsMap[ERROR] {
		errorLogger.Output(2, fmt.Sprintln(v...))
	}
}

//Errorf error and format
func Errorf(format string, v ...interface{}) {
	if logLevel >= LogLevelsMap[TRACE] && logLevel <= LogLevelsMap[ERROR] {
		errorLogger.Output(2, fmt.Sprintf(format, v...))
	}
}

//SetLogLevel define the level
func SetLogLevel(loggingLevel string) {
	logLevel = LogLevelsMap[loggingLevel]
}
