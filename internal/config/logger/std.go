package logger

import (
	"fmt"
	"log"
	"os"
)

type StdLogger struct {
	info  *log.Logger
	warn  *log.Logger
	error *log.Logger
}

func NewLogger() *StdLogger {
	return &StdLogger{
		info:  log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		warn:  log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile),
		error: log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
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
