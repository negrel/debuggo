package log

import (
	"fmt"
	"log"
	"os"
)

var std = New(os.Stderr, "", log.LstdFlags)

// Panic level

func Panic(args ...interface{}) {
	std.Panic(args...)
}

func Panicf(format string, args ...interface{}) {
	std.Panicf(format, args...)
}

func Panicln(args ...interface{}) {
	std.Panicln(args...)
}

func Panicfn(fn func() []interface{}) {
	std.Panic(fn()...)
}

// Fatal level

func Fatal(args ...interface{}) {
	std.Fatal(args...)
}

func Fatalf(format string, args ...interface{}) {
	Fatal(fmt.Sprintf(format, args...))
}

func Fatalln(args ...interface{}) {
	Fatal(fmt.Sprintln(args...))
}

func Fatalfn(fn func() []interface{}) {
	Fatal(fn()...)
}

// Error level

func Error(args ...interface{}) {
	std.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	Error(fmt.Sprintf(format, args...))
}

func Errorln(args ...interface{}) {
	Error(fmt.Sprintln(args...))
}

func Errorfn(fn func() []interface{}) {
	Error(fn()...)
}

// Warn level

func Warn(args ...interface{}) {
	std.Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	Warn(fmt.Sprintf(format, args...))
}

func Warnln(args ...interface{}) {
	Warn(fmt.Sprintln(args...))
}

func Warnfn(fn func() []interface{}) {
	Warn(fn()...)
}

// Info level

func Info(args ...interface{}) {
	std.Info(args...)
}

func Infof(format string, args ...interface{}) {
	Info(fmt.Sprintf(format, args...))
}

func Infoln(args ...interface{}) {
	Info(fmt.Sprintln(args...))
}

func Infofn(fn func() []interface{}) {
	Info(fn()...)
}

// Debug level

func Debug(args ...interface{}) {
	std.Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	Debug(fmt.Sprintf(format, args...))
}

func Debugln(args ...interface{}) {
	Debug(fmt.Sprintln(args...))
}

func Debugfn(fn func() []interface{}) {
	Debug(fn()...)
}

// Trace level

func Trace(args ...interface{}) {
	std.Trace(args...)
}

func Tracef(format string, args ...interface{}) {
	Trace(fmt.Sprintf(format, args...))
}

func Traceln(args ...interface{}) {
	Trace(fmt.Sprintln(args...))
}

func Tracefn(fn func() []interface{}) {
	Trace(fn()...)
}
