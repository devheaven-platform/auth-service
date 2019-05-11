package utils

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/sirupsen/logrus"
)

// StructuredLogger represents an instance of an logrus
// logger.
type StructuredLogger struct {
	Logger *logrus.Logger
}

// StructuredLoggerEntry represents an instance of an
// logrus log entry.
type StructuredLoggerEntry struct {
	Logger logrus.FieldLogger
}

// NewStructuredLogger is used to create a new request
// logger. It takes an http.Handler as parameter and
// returns an instance of an chi request logger.
func NewStructuredLogger(logger *logrus.Logger) func(next http.Handler) http.Handler {
	return middleware.RequestLogger(&StructuredLogger{logger})
}

// NewLogEntry is invoked every time a request is made.
// It takes an instance of http.Request as parameter and
// returns an chi log entry.
func (l *StructuredLogger) NewLogEntry(r *http.Request) middleware.LogEntry {
	entry := &StructuredLoggerEntry{Logger: logrus.NewEntry(l.Logger)}
	logFields := logrus.Fields{}

	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	logFields["http_scheme"] = scheme
	logFields["http_proto"] = r.Proto
	logFields["http_method"] = r.Method
	logFields["remote_addr"] = r.RemoteAddr

	logFields["uri"] = fmt.Sprintf("%s://%s%s", scheme, r.Host, r.RequestURI)

	entry.Logger = entry.Logger.WithFields(logFields)

	entry.Logger.Infoln("Request received")

	return entry
}

// Write is invoked be logrus to write the log message
// to std out. It takes response status, reponse bytes
// length and an response time in ms as parameters.
func (l *StructuredLoggerEntry) Write(status, bytes int, elapsed time.Duration) {
	l.Logger = l.Logger.WithFields(logrus.Fields{
		"response_status":       status,
		"response_bytes_length": bytes,
		"response_time_ms":      float64(elapsed.Nanoseconds()) / 1000000.0,
	})

	l.Logger.Infoln("Request processed")
}

// Panic is invoked by logrus when a panic event occurred.
// It takes an interface and stack trace a parameters.
func (l *StructuredLoggerEntry) Panic(v interface{}, stack []byte) {
	l.Logger = l.Logger.WithFields(logrus.Fields{
		"stack": string(stack),
		"panic": fmt.Sprintf("%+v", v),
	})
}
