package logger

import (
	"context"
	"fmt"
	log "github.com/apsdehal/go-logger"
	"os"
)

const (
	PROD = "prod"
	DEV  = "dev"
)
const RequestIdKey = "requestId"

type Mode string

const (
	PRINT  Mode = "printF"
	NOTICE Mode = "noticeF"
)

func (l *Logger) PrintContext(ctx context.Context, mode Mode, format string, err ...interface{}) {
	formatWithContext := fmt.Sprintf("[%v] %v ", ctx.Value(RequestIdKey), format)
	l.PrintOnMode(mode, formatWithContext, err...)
}

func (l *Logger) PrintRequestId(requestId string, mode Mode, format string, err ...interface{}) {
	formatWithContext := fmt.Sprintf("[%v] %v ", requestId, format)
	l.PrintOnMode(mode, formatWithContext, err...)
}

func (l *Logger) PrintOnMode(mode Mode, format string, err ...interface{}) {
	switch mode {
	case PRINT:
		l.Panicf(format, err...)
	case NOTICE:
		l.NoticeF(format, err...)
	default:
		l.Printf(format, err...)
	}
}

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

type Logger struct {
	Mode   string
	Logger *log.Logger
}

type ContextLogging interface {
	AdvancedLogging
	PrintContext(ctx context.Context, mode Mode, format string, v ...interface{})
	PrintRequestId(requestId string, mode Mode, format string, err ...interface{})
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
