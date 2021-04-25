package logger

import (
	"os"

	log "github.com/apsdehal/go-logger"
)

const (
	PROD = "prod"
	DEV  = "dev"
)

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
