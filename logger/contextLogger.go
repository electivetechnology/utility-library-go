package logger

import (
	"context"
	"fmt"
	"github.com/electivetechnology/utility-library-go/hash"
	"github.com/gin-gonic/gin"
	"time"
)

const RequestIdKey = "requestId"

var startTime time.Time

type ContextLogging interface {
	AdvancedLogging
	PrintFContext(ctx context.Context, format string, v ...interface{})
	CriticalFContext(ctx context.Context, format string, v ...interface{})
	WarningFContext(ctx context.Context, format string, v ...interface{})
	ErrorFContext(ctx context.Context, format string, v ...interface{})
	NoticeFContext(ctx context.Context, format string, v ...interface{})
	InfoFContext(ctx context.Context, format string, v ...interface{})
	DebugFContext(ctx context.Context, format string, v ...interface{})
	PrintFRequest(requestId string, format string, v ...interface{})
	CriticalFRequest(requestId string, format string, v ...interface{})
	WarningFRequest(requestId string, format string, v ...interface{})
	ErrorFRequest(requestId string, format string, v ...interface{})
	NoticeFRequest(requestId string, format string, v ...interface{})
	InfoFRequest(requestId string, format string, v ...interface{})
	DebugFRequest(requestId string, format string, v ...interface{})
}

func (l *Logger) LoggerRequestHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestId := hash.GenerateHash(12)
		l.StartContext(requestId)

		c.Set(RequestIdKey, requestId)

		c.Writer.Header().Set("X-Log-Id", requestId)
		c.Next()
		l.EndContext(requestId)
	}
}

func contextFormat(ctx context.Context, format string) string {
	return fmt.Sprintf("[%v] %v ", ctx.Value(RequestIdKey), format)
}

func requestFormat(requestId string, format string) string {
	return fmt.Sprintf("[%v] %v ", requestId, format)
}

func (l *Logger) PrintFContext(ctx context.Context, format string, err ...interface{}) {
	l.Printf(contextFormat(ctx, format), err)
}

func (l *Logger) CriticalFContext(ctx context.Context, format string, err ...interface{}) {
	l.CriticalF(contextFormat(ctx, format), err)
}

func (l *Logger) WarningFContext(ctx context.Context, format string, err ...interface{}) {
	l.WarningF(contextFormat(ctx, format), err)
}

func (l *Logger) ErrorFContext(ctx context.Context, format string, err ...interface{}) {
	l.ErrorF(contextFormat(ctx, format), err)
}

func (l *Logger) NoticeFContext(ctx context.Context, format string, err ...interface{}) {
	l.NoticeF(contextFormat(ctx, format), err)
}

func (l *Logger) InfoFContext(ctx context.Context, format string, err ...interface{}) {
	l.InfoF(contextFormat(ctx, format), err)
}

func (l *Logger) DebugFContext(ctx context.Context, format string, err ...interface{}) {
	l.DebugF(contextFormat(ctx, format), err)
}

func (l *Logger) PrintFRequest(requestId string, format string, err ...interface{}) {
	l.Printf(requestFormat(requestId, format), err)
}

func (l *Logger) CriticalFRequest(requestId string, format string, err ...interface{}) {
	l.CriticalF(requestFormat(requestId, format), err)
}

func (l *Logger) WarningFRequest(requestId string, format string, err ...interface{}) {
	l.WarningF(requestFormat(requestId, format), err)
}

func (l *Logger) ErrorFRequest(requestId string, format string, err ...interface{}) {
	l.ErrorF(requestFormat(requestId, format), err)
}

func (l *Logger) NoticeFRequest(requestId string, format string, err ...interface{}) {
	l.NoticeF(requestFormat(requestId, format), err)
}

func (l *Logger) InfoFRequest(requestId string, format string, err ...interface{}) {
	l.InfoF(requestFormat(requestId, format), err)
}

func (l *Logger) DebugFRequest(requestId string, format string, err ...interface{}) {
	l.DebugF(requestFormat(requestId, format), err)
}

func (l *Logger) StartContext(requestId string) {
	startTime = time.Now()
	withContext := fmt.Sprintf("[%v] %v ", requestId, "Request Started")
	l.Printf(withContext)
}

func (l *Logger) EndContext(requestId string) {
	currentTime := time.Now()
	diff := currentTime.Sub(startTime).Milliseconds()
	withContext := fmt.Sprintf("[%v] Request Completed after %v ms", requestId, diff)
	l.Printf(withContext)
}
