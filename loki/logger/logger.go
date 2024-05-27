package logger

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"path/filepath"
	"runtime"
	"time"
)

type TraceIDFn func(ctx context.Context) string

type Level slog.Level

func (l Level) String() string {
	return slog.Level(l).String()
}

type Attribute struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Record struct {
	Time       time.Time
	Message    string
	Level      Level
	Attributes []Attribute
}

const (
	LevelInfo  Level = Level(slog.LevelInfo)
	LevelWarn  Level = Level(slog.LevelWarn)
	LevelError Level = Level(slog.LevelError)
	LevelDebug Level = Level(slog.LevelDebug)
)

func toRecord(r slog.Record) Record {

	atts := make([]Attribute, r.NumAttrs())

	f := func(attr slog.Attr) bool {
		atts = append(atts, Attribute{
			Name:  attr.Key,
			Value: attr.Value.Resolve().String(),
		})
		return true
	}
	r.Attrs(f)

	return Record{
		Time:       r.Time,
		Message:    r.Message,
		Level:      Level(r.Level),
		Attributes: atts,
	}
}

type EventFn func(ctx context.Context, r Record)

type Events struct {
	Debug EventFn
	Info  EventFn
	Warn  EventFn
	Error EventFn
}

type Logger struct {
	handler   slog.Handler
	traceIDFn TraceIDFn
	events    Events
}

func (l *Logger) GetHandler() slog.Handler {
	return l.handler
}

func (h *Logger) handleLog(ctx context.Context, r slog.Record) error {

	switch r.Level {
	case slog.LevelDebug:
		if h.events.Debug != nil {
			go h.events.Debug(ctx, toRecord(r))
		}

	case slog.LevelError:
		if h.events.Error != nil {
			go h.events.Error(ctx, toRecord(r))
		}

	case slog.LevelWarn:
		if h.events.Warn != nil {
			go h.events.Warn(ctx, toRecord(r))
		}

	case slog.LevelInfo:
		if h.events.Info != nil {
			go h.events.Info(ctx, toRecord(r))
		}
	}

	return h.handler.Handle(ctx, r)
}

func New(serviceName string, w io.Writer, minLevel Level, traceIDFn TraceIDFn, events Events) *Logger {
	f := func(groups []string, a slog.Attr) slog.Attr {
		if a.Key == slog.SourceKey {
			if source, ok := a.Value.Any().(*slog.Source); ok {
				v := fmt.Sprintf("%s:%d", filepath.Base(source.File), source.Line)
				return slog.Attr{Key: "file", Value: slog.StringValue(v)}
			}
		}

		return a
	}
	handler := slog.Handler(slog.NewJSONHandler(w, &slog.HandlerOptions{
		AddSource:   true,
		Level:       slog.Level(minLevel),
		ReplaceAttr: f,
	}))
	attrs := []slog.Attr{
		slog.String("service", serviceName),
	}

	return &Logger{
		handler:   handler.WithAttrs(attrs),
		traceIDFn: traceIDFn,
		events:    events,
	}
}

func (log *Logger) write(ctx context.Context, level slog.Level, msg string, args ...interface{}) {
	slogLevel := slog.Level(level)
	var pcs [1]uintptr
	runtime.Callers(3, pcs[:])

	r := slog.Record{
		Level:   slogLevel,
		Message: msg,
		Time:    time.Now(),
		PC:      pcs[0],
	}
	if log.traceIDFn != nil {
		r.Add(slog.String("trace_id", log.traceIDFn(ctx)))
	}
	r.Add(args...)
	log.handleLog(ctx, r)
}

func (log *Logger) Debug(ctx context.Context, msg string, args ...interface{}) {
	log.write(ctx, slog.LevelDebug, msg, args...)
}

func (log *Logger) Info(ctx context.Context, msg string, args ...interface{}) {
	log.write(ctx, slog.LevelInfo, msg, args...)
}

func (log *Logger) Warn(ctx context.Context, msg string, args ...interface{}) {
	log.write(ctx, slog.LevelWarn, msg, args...)
}

func (log *Logger) Error(ctx context.Context, msg string, args ...interface{}) {
	log.write(ctx, slog.LevelError, msg, args...)
}
