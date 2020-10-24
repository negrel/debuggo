// +build panic

package log

import (
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

}

func Fatalf(format string, args ...interface{}) {

}

func Fatalln(args ...interface{}) {

}

func Fatalfn(fn func() []interface{}) {

}

// Error level

func Error(args ...interface{}) {

}

func Errorf(format string, args ...interface{}) {

}

func Errorln(args ...interface{}) {

}

func Errorfn(fn func() []interface{}) {

}

// Warn level

func Warn(args ...interface{}) {

}

func Warnf(format string, args ...interface{}) {

}

func Warnln(args ...interface{}) {

}

func Warnfn(fn func() []interface{}) {

}

// Info level

func Info(args ...interface{}) {

}

func Infof(format string, args ...interface{}) {

}

func Infoln(args ...interface{}) {

}

func Infofn(fn func() []interface{}) {

}

// Debug level

func Debug(args ...interface{}) {

}

func Debugf(format string, args ...interface{}) {

}

func Debugln(args ...interface{}) {

}

func Debugfn(fn func() []interface{}) {

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

}

func (l *Logger) Fatalf(format string, args ...interface{}) {

}

func (l *Logger) Fatalln(args ...interface{}) {

}

func (l *Logger) Fatalfn(fn func() []interface{}) {

}

// Error level

func (l *Logger) Error(args ...interface{}) {

}

func (l *Logger) Errorf(format string, args ...interface{}) {

}

func (l *Logger) Errorln(args ...interface{}) {

}

func (l *Logger) Errorfn(fn func() []interface{}) {

}

// Warn level

func (l *Logger) Warn(args ...interface{}) {

}

func (l *Logger) Warnf(format string, args ...interface{}) {

}

func (l *Logger) Warnln(args ...interface{}) {

}

func (l *Logger) Warnfn(fn func() []interface{}) {

}

// Info level

func (l *Logger) Info(args ...interface{}) {

}

func (l *Logger) Infof(format string, args ...interface{}) {

}

func (l *Logger) Infoln(args ...interface{}) {

}

func (l *Logger) Infofn(fn func() []interface{}) {

}

// Debug level

func (l *Logger) Debug(args ...interface{}) {

}

func (l *Logger) Debugf(format string, args ...interface{}) {

}

func (l *Logger) Debugln(args ...interface{}) {

}

func (l *Logger) Debugfn(fn func() []interface{}) {

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
