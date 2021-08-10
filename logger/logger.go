package logger

import (
	"context"
	"fmt"
	"github.com/electivetechnology/utility-library-go/hash"
	"github.com/gin-gonic/gin"
	"os"
	"time"

	log "github.com/apsdehal/go-logger"
)

const (
	PROD = "prod"
	DEV  = "dev"
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

type Logger struct {
	Mode   string
	Logger *log.Logger
}

type ContextLogging interface {
	AdvancedLogging
	WithRequestContext(ctx context.Context, format string, v ...interface{})
	WithRequestId(string, format string, v ...interface{})
	StartRequestContext(requestId string)
	EndRequestContext(requestId string)
}

const RequestIdKey = "requestId"

var startTime time.Time

func (l *Logger) LoggerRequestHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime = time.Now()
		requestId := hash.GenerateHash(12)
		l.StartRequestContext(requestId)

		c.Set(RequestIdKey, requestId)

		c.Writer.Header().Set("X-Log-Id", requestId)
		c.Next()
		l.EndRequestContext(requestId)
	}
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

func (l *Logger) WithRequestContext(ctx context.Context, format string, err ...interface{}) {
	withContext := fmt.Sprintf("[%v] %v ", ctx.Value(RequestIdKey), format)
	l.Printf(withContext, err...)
}

func (l *Logger) WithRequestId(requestId string, format string, err ...interface{}) {
	withId := fmt.Sprintf("[%v] %v ", requestId, format)
	l.Printf(withId, err...)
}

func (l *Logger) WithWorkerContext(ctx context.Context, format string, err ...interface{}) {
	withContext := fmt.Sprintf("[%v] %v ", ctx.Value(RequestIdKey), format)
	l.Printf(withContext, err...)
}

func (l *Logger) StartRequestContext(requestId string) {
	withContext := fmt.Sprintf("[%v] %v ", requestId, "Request Started")
	l.Printf(withContext)
}

func (l *Logger) EndRequestContext(requestId string) {
	currentTime := time.Now()
	diff := currentTime.Sub(startTime).Milliseconds()
	withContext := fmt.Sprintf("[%v] Request Completed after %v ms", requestId, diff)
	l.Printf(withContext)
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
