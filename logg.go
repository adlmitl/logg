package logg

import (
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
)

// List variables for recording in log.
var (
	Error   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
)

// InitLogg - Initialization logg.
// pathToFileLogg - specify the path where the log file will be stored.
func InitLogg(pathToFileLogg string) {
	file, err := os.OpenFile(pathToFileLogg, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("\033[1;31m[E] Error while opening logger file: %s\033[0m", err.Error())
	}
	Info = log.New(file, " INFO ---", log.Lmsgprefix|log.LstdFlags)
	Error = log.New(file, "ERROR --- ", log.Lmsgprefix|log.LstdFlags)
	Warning = log.New(file, " WARN --- ", log.Lmsgprefix|log.LstdFlags)
}

// LogI - Print of info msg in console and log file.
func LogI(msg string) {
	log.Printf("\033[1;34m[I] %s\033[0m", msg)
	Info.Printf(" %s - %s", fileWithFuncAndLineNum(), msg)
}

// LogE - Print of errors msg in console and log file.
func LogE(msg error, err string) {
	log.Printf("%s \033[1;33m[E] %s: %s\033[0m", fileWithLineNum(), msg, err)
	Error.Printf(" %s - %s", fileWithFuncAndLineNum(), msg, err)
}

// LogW - Print of warning msg in console and log file.
func LogW(msg string, err string) {
	log.Printf("%s \033[1;33m[W] %s: %s\033[0m", fileWithLineNum(), msg, err)
	Warning.Printf("%s - %s: %s", fileWithFuncAndLineNum(), msg, err)
}

/* The number of stack frames to skip before recording to the PC, where 0 identifies the frame to the callers
themselves and 1 identifies the caller. Returns the number of records written to the computer.*/
const skipNumOfStackFrame = 3

// fileWithLineNum - Return name file and number string current file.
func fileWithLineNum() string {
	pc := make([]uintptr, 15)
	n := runtime.Callers(skipNumOfStackFrame, pc)
	frame, _ := runtime.CallersFrames(pc[:n]).Next()
	idxFile := strings.LastIndexByte(frame.File, '/')

	return frame.File[idxFile+1:] + ":" + strconv.FormatInt(int64(frame.Line), 10)
}

// fileWithFuncAndLineNum - Return name file with function and number string current file.
func fileWithFuncAndLineNum() string {
	pc := make([]uintptr, 15)
	n := runtime.Callers(skipNumOfStackFrame, pc)
	frame, _ := runtime.CallersFrames(pc[:n]).Next()
	idxFile := strings.LastIndexByte(strconv.Itoa(frame.Line), '/')

	return "[" + frame.Function + "] - " + frame.File[idxFile+1:] + ":" + strconv.FormatInt(int64(frame.Line), 10)
}
