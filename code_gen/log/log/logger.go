package log

import (
	"fmt"
	"io"
	"log"
)

type Logger struct {
	internal *log.Logger
}

// New creates a new Logger. The out variable sets the destination to which log data will be written.
// The prefix appears at the beginning of each generated log line, or after the log header if the Lmsgprefix flag is
// provided. The flag argument defines the logging properties.
func New(out io.Writer, prefix string, flag int) *Logger {
	return &Logger{
		log.New(out, prefix, flag),
	}
}

// SetOutput sets the output destination for the logger.
func (l *Logger) SetOutput(w io.Writer) {
	l.internal.SetOutput(w)
}

// SetFlags sets the output flags for the logger.
// The flag bits are Ldate, Ltime, and so on.
func (l *Logger) SetFlags(flag int) {
	l.internal.SetFlags(flag)
}

// SetPrefix sets the output prefix for the logger.
func (l *Logger) SetPrefix(prefix string) {
	l.internal.SetPrefix(prefix)
}

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
	l.internal.Print(args...)
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
	l.internal.Print(args...)
}

func (l *Logger) Tracef(format string, args ...interface{}) {
	l.Trace(fmt.Sprintf(format, args...))
}

func (l *Logger) Traceln(args ...interface{}) {
	l.Trace(fmt.Sprintln(args...))
}

func (l *Logger) Tracefn(fn func() []interface{}) {
	l.Trace(fn()...)
}
