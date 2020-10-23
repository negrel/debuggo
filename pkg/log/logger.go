package log

import (
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
