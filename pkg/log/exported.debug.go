// +build debug

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

}

func Tracef(format string, args ...interface{}) {

}

func Traceln(args ...interface{}) {

}

func Tracefn(fn func() []interface{}) {

}

// Logger

// Panic level

func (l *Logger) Panic(args ...interface{}) {
	l.internal.Panic(args...)
}

func (l *Logger) Panicf(format string, args ...interface{}) {
	l.internal.Panicf(format, args...)
}

func (l *Logger) Panicln(args ...interface{}) {
	l.internal.Panicln(args...)
}

func (l *Logger) Panicfn(fn func() []interface{}) {
	l.internal.Panic(fn()...)
}

// Fatal level

func (l *Logger) Fatal(args ...interface{}) {
	l.internal.Fatal(args...)
}

func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.Fatal(fmt.Sprintf(format, args...))
}

func (l *Logger) Fatalln(args ...interface{}) {
	l.Fatal(fmt.Sprintln(args...))
}

func (l *Logger) Fatalfn(fn func() []interface{}) {
	l.Fatal(fn()...)
}

// Error level

func (l *Logger) Error(args ...interface{}) {
	l.internal.Print(args...)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.Error(fmt.Sprintf(format, args...))
}

func (l *Logger) Errorln(args ...interface{}) {
	l.Error(fmt.Sprintln(args...))
}

func (l *Logger) Errorfn(fn func() []interface{}) {
	l.Error(fn()...)
}

// Warn level

func (l *Logger) Warn(args ...interface{}) {
	l.internal.Print(args...)
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	l.Warn(fmt.Sprintf(format, args...))
}

func (l *Logger) Warnln(args ...interface{}) {
	l.Warn(fmt.Sprintln(args...))
}

func (l *Logger) Warnfn(fn func() []interface{}) {
	l.Warn(fn()...)
}

// Info level

func (l *Logger) Info(args ...interface{}) {
	l.internal.Print(args...)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.Info(fmt.Sprintf(format, args...))
}

func (l *Logger) Infoln(args ...interface{}) {
	l.Info(fmt.Sprintln(args...))
}

func (l *Logger) Infofn(fn func() []interface{}) {
	l.Info(fn()...)
}

// Debug level

func (l *Logger) Debug(args ...interface{}) {
	l.internal.Print(args...)
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	l.Debug(fmt.Sprintf(format, args...))
}

func (l *Logger) Debugln(args ...interface{}) {
	l.Debug(fmt.Sprintln(args...))
}

func (l *Logger) Debugfn(fn func() []interface{}) {
	l.Debug(fn()...)
}

// Trace level

func (l *Logger) Trace(args ...interface{}) {

}

func (l *Logger) Tracef(format string, args ...interface{}) {

}

func (l *Logger) Traceln(args ...interface{}) {

}

func (l *Logger) Tracefn(fn func() []interface{}) {

}
