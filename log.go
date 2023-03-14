package main

import (
	"log"
	"os"
)

// provide vendor agnostic logging to add additional features in the future,
// while providing a consistent interface

var (
	warningLogger *log.Logger
	infoLogger    *log.Logger
	errorLogger   *log.Logger
	fatalLogger   *log.Logger
	panicLogger   *log.Logger
)

func init() {
	// setup logger levels, currently outputting to stdio
	warningLogger = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime)
	infoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	errorLogger = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime)
	fatalLogger = log.New(os.Stdout, "FATAL: ", log.Ldate|log.Ltime)
	panicLogger = log.New(os.Stdout, "PANIC: ", log.Ldate|log.Ltime)
}

func LogWarning(vals ...interface{}) {
	warningLogger.Println(vals)
}

func LogInfo(vals ...interface{}) {
	infoLogger.Println(vals)
}

func LogError(vals ...interface{}) {
	errorLogger.Println(vals)
}

func LogFatal(vals ...interface{}) {
	fatalLogger.Println(vals)
	os.Exit(1)
}

func LogPanic(vals ...interface{}) {
	panicLogger.Println(vals)
	panic(vals)
}
