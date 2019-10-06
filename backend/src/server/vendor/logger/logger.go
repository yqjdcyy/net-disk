package logger

import (
	"fmt"
	"log"
	"os"
	"utils"

	"bitbucket.org/ansenwork/ilog"
)

// Logger 本地记录登记
type Logger struct {
	printer *log.Logger
	file    *os.File
}

var logger = new(Logger)
var (
	prefixDebug = " [debug] "
	prefixTrace = " [trace] "
	prefixInfo  = " [info] "
	prefixWarn  = " [warn] "
	prefixError = " [error] "
	prefixPanic = " [panic] "
)

// SetLogPath 设置日志目录
func SetLogPath(path string) {

	file, err := utils.OpenOrCreate(path)
	if nil != err {
		ilog.Panicf("fail to open or crate file[%v]: %v", path, err.Error())
	}
	logger.file = file
	logger.printer = log.New(file, "[Statistic]", log.LstdFlags)
	ilog.Infof("Statistic.Logger.init bind to file[%v]", path)
}

// Close 关闭文件资源
func Close() {

	if nil != logger.file {
		err := logger.file.Close()
		if nil != err {
			ilog.Panicf("fail to close log file: %v", err.Error())
		}
	}
}

// Debug write the debug msg
func Debug(format string) {
	logger.printer.Println(prefixDebug, format)
}

// Debugf write the debug msg
func Debugf(format string, args ...interface{}) {
	logger.printer.Printf("%s%s", prefixDebug, fmt.Sprintf(format, args...))
}

// Trace write the trace msg
func Trace(format string) {
	logger.printer.Println(prefixTrace, format)
}

// Tracef write the trace msg
func Tracef(format string, args ...interface{}) {
	logger.printer.Printf("%s%s", prefixTrace, fmt.Sprintf(format, args...))
}

// Info write the info msg
func Info(format string) {
	logger.printer.Println(prefixInfo, format)
}

// Infof write the info msg
func Infof(format string, args ...interface{}) {
	logger.printer.Printf("%s%s", prefixInfo, fmt.Sprintf(format, args...))
}

// Warn write the warn msg
func Warn(format string) {
	logger.printer.Println(prefixWarn, format)
}

// Warnf write the warn msg
func Warnf(format string, args ...interface{}) {
	logger.printer.Printf("%s%s", prefixWarn, fmt.Sprintf(format, args...))
}

// Error write the error msg
func Error(format string) {
	logger.printer.Println(prefixError, format)
}

// Errorf write the error msg
func Errorf(format string, args ...interface{}) {
	logger.printer.Printf("%s%s", prefixError, fmt.Sprintf(format, args...))
}

// Panic write the panic msg
func Panic(format string) {
	logger.printer.Fatal(prefixPanic, format)
}

// Panicf write the panic msg
func Panicf(format string, args ...interface{}) {
	logger.printer.Fatalf("%s%s", prefixPanic, fmt.Sprintf(format, args...))
}
