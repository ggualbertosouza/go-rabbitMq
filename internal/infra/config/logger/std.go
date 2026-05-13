package logger

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type StdLogger struct {
	info  *log.Logger
	warn  *log.Logger
	error *log.Logger
}

func NewLogger(application string) *StdLogger {
	return &StdLogger{
		info:  log.New(os.Stdout, formatPrefix(application, "INFO"), log.Ldate|log.Ltime|log.Lshortfile),
		warn:  log.New(os.Stdout, formatPrefix(application, "WARN"), log.Ldate|log.Ltime|log.Lshortfile),
		error: log.New(os.Stdout, formatPrefix(application, "ERROR"), log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (l *StdLogger) Info(msg string, fields ...Field) {
	l.info.Println(format(msg, fields...))
}

func (l *StdLogger) Warn(msg string, fields ...Field) {
	l.warn.Println(format(msg, fields...))
}

func (l *StdLogger) Error(msg string, fields ...Field) {
	l.error.Println(format(msg, fields...))
}

func format(msg string, fields ...Field) string {
	if len(fields) == 0 {
		return msg
	}

	out := msg
	for _, f := range fields {
		out += fmt.Sprintf(" %s=%v", f.Key, f.Value)
	}
	return out
}

func formatPrefix(application, level string) string {
	return strings.ToUpper(application) + " " + level + ": "
}
