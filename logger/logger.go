package logger

import (
	"context"
	"os"

	log "github.com/apsdehal/go-logger"
)

const (
	PROD = "prod"
	DEV  = "dev"
)

type correlationIdType int

const (
	requestIdKey correlationIdType = iota
	sessionIdKey
)

type Logging interface {
	Fatalf(format string, v ...interface{})
	Panicf(format string, v ...interface{})
	Printf(format string, v ...interface{})
}

type AdvancedLogging interface {
	Logging
	InfoF(format string, v ...interface{})
	NoticeF(format string, v ...interface{})
	WarningF(format string, v ...interface{})
	ErrorF(format string, v ...interface{})
}

// WithRqId returns a context which knows its request ID
func WithRqId(ctx context.Context, rqId string) context.Context {
	return context.WithValue(ctx, requestIdKey, requestId)
}

// WithSessionId returns a context which knows its session ID
func WithSessionId(ctx context.Context, sessionId string) context.Context {
	return context.WithValue(ctx, sessionIdKey, sessionId)
}

type Logger struct {
	Mode   string
	Logger *log.Logger
}

func NewLogger(module string) *Logger {
	log, _ := log.New(module, 1)
	log.SetFormat("#%{id} %{time} ▶ [%{module}][%{level}]: %{message}")

	mode := os.Getenv("LOGGER_MODE")
	if mode != PROD {
		mode = DEV
	}

	return &Logger{
		Mode:   mode,
		Logger: log,
	}
}

func ContextLogger(module string, ctx context.Context) *Logger {
	logger := NewLogger(module)

	return logger
}

func (l *Logger) Fatalf(format string, err ...interface{}) {
	l.CriticalF(format, err...)
}

func (l *Logger) Panicf(format string, err ...interface{}) {
	l.CriticalF(format, err...)
}

func (l *Logger) Printf(format string, err ...interface{}) {
	l.DebugF(format, err...)
}

func (l *Logger) CriticalF(format string, err ...interface{}) {
	l.Logger.CriticalF(format, err...)
}

func (l *Logger) WarningF(format string, err ...interface{}) {
	if l.Mode == DEV {
		l.Logger.SetLogLevel(log.DebugLevel)
		l.Logger.WarningF(format, err...)
	}
}

func (l *Logger) ErrorF(format string, err ...interface{}) {
	if l.Mode == DEV {
		l.Logger.SetLogLevel(log.DebugLevel)
		l.Logger.ErrorF(format, err...)
	}
}

func (l *Logger) NoticeF(format string, err ...interface{}) {
	if l.Mode == DEV {
		l.Logger.SetLogLevel(log.DebugLevel)
		l.Logger.NoticeF(format, err...)
	}
}

func (l *Logger) InfoF(format string, err ...interface{}) {
	if l.Mode == DEV {
		l.Logger.SetLogLevel(log.DebugLevel)
		l.Logger.InfoF(format, err...)
	}
}

func (l *Logger) DebugF(format string, err ...interface{}) {
	if l.Mode == DEV {
		l.Logger.SetLogLevel(log.DebugLevel)
		l.Logger.DebugF(format, err...)
	}
}
