package common

import (
	"context"
	"net/http"
	"time"
)

type Logger interface {
	Debug(ctx context.Context, msg string, args ...interface{})
	Info(ctx context.Context, msg string, args ...interface{})
	Warn(ctx context.Context, msg string, args ...interface{})
	Error(ctx context.Context, msg string, args ...interface{})
}

type HttpContext interface {
	GetContext() context.Context
	SetContext(ctx context.Context)
	GetTraceID() string
	GetUserID() string
	SetTraceID(traceID string)
	SetUserID(userID string)
	GetTime() time.Time
	GetStatusCode() int
	SetStatusCode(statusCode int)
	Text(status int, body string) error
	SetHeader(key, value string)
	JSON(status int, body interface{}) error
	Decode(val any) error
	GetRemoteAddr() string
	GetRequest() *http.Request
	GetWriter() http.ResponseWriter
	GetLogger() Logger
	GetMethod() string
	GetPath() string
	GetParam(key string) string
}
