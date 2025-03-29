// Copyright © 2024 Bruce Smith <bruceesmith@gmail.com>
// Use of this source code is governed by the MIT
// License that can be found in the LICENSE file.

/*
Package logger supports logging and tracing based on the standard library package [log/slog].

Debug, Error, Info and Warn operate like their package slog equivalents, with the level of logging modifiable
using SetLevel.

A custom logging level (LevelTrace) can be supplied to SetLevel to enable tracing. Tracing can
be unconditional when calling Trace, or only enabled for pre-defined identifiers when calling TraceID. Identifiers
for TraceID are registered by calling SetTraceIDs.

By default, all debug, error, info and warn messages go to Stdout, and traces go to Stderr; these destinations
can be changed by calling RedirectNormal and RedirectTrace respectively.

A number of settings can be changed for one or both of the normal (non-trace) and trace loggers by calling
[Configure] - the format of log records, their destination, and whether each record contains a timestamp.

When used in [cli applications], a cli.Flag representing a LogLevel can be provided using the LogLevelFlag type.

[cli applications]: https://github.com/urfave/cli
*/
package logger

//go:generate ./make_doc.sh

import (
	"context"
	"io"
	"log/slog"
	"runtime"
	"strings"
	"time"
)

// jsonHandler returns a JSONHandler configured per the config settings
func jsonHandler(w io.Writer, trace bool) slog.Handler {
	return slog.NewJSONHandler(
		w,
		&slog.HandlerOptions{
			AddSource:   trace,
			Level:       &level,
			ReplaceAttr: replacer(trace),
		},
	)
}

// levelAttr replaces custom log levels with their String name in log records
func levelAttr(a slog.Attr) slog.Attr {
	if a.Key == slog.LevelKey {
		level, ok := a.Value.Any().(slog.Level)
		if ok {
			loglevel := LogLevel(level)
			a.Value = slog.StringValue((&loglevel).String())
		}
	}
	return a
}

// replcer returns a function used as ReplaceAttr in loggers
func replacer(trace bool) func(_ []string, a slog.Attr) slog.Attr {
	return func(_ []string, a slog.Attr) slog.Attr {
		a = levelAttr(a)
		a = timeAttr(a, trace)
		return a
	}
}

// textHandler returns a TextNHandler configured per the config settings
func textHandler(w io.Writer, trace bool) slog.Handler {
	return slog.NewTextHandler(
		w,
		&slog.HandlerOptions{
			Level:       &level,
			ReplaceAttr: replacer(trace),
		},
	)
}

// timeAttr removes the "Time" fragment from a log record if so configured
func timeAttr(a slog.Attr, trace bool) slog.Attr {
	if a.Key == slog.TimeKey &&
		((trace && config.Trace.OmitTime) ||
			(!trace && config.Normal.OmitTime)) {
		return slog.Attr{}
	}
	return a
}

// Debug emits a debug log
func Debug(msg string, args ...any) {
	slog.Debug(msg, args...)
}

// Error emits an error log
func Error(msg string, args ...any) {
	slog.Error(msg, args...)
}

// Info emits an info log
func Info(msg string, args ...any) {
	slog.Info(msg, args...)
}

// Level returns the current logging level as a string
func Level() string {
	ll := LogLevel(level.Level())
	return (&ll).String()
}

// RedirectStandard changes the destination for normal (non-trace) logsDestinationSetting argument
//
// Deprecated: RedirectStandard() should be replaced by a call to Configure()
// with a DestinationSetting argument
func RedirectStandard(w io.Writer) {
	normalDestination(w)
}

// RedirectTrace changes the destination for normal (non-trace) logs
//
// Deprecated: RedirectTrace() should be replaced by a call to Configure()
// with a DestinationSetting argument
func RedirectTrace(w io.Writer) {
	traceDestination(w)
}

// SetFormat changes the format of log entries
//
// Deprecated: SetFormat() should be replaced by calls to Configure() with a
// FormatSetting argument. An advantage of Configure() is that the format of
// the standard logger can be configured differently to that of the Trace logger
func SetFormat(f Format) {
	normalFormat(f)
	traceFormat(f)
}

const (
	// LevelTrace can be set to enable tracing
	LevelTrace slog.Level = -10
)

// SetLevel sets the default level of logging
func SetLevel(l slog.Level) {
	level.Set(l)
}

// SetTraceIds registers identifiers for future tracing
func SetTraceIds(ids ...string) {
	for _, id := range ids {
		_ = config.traceIds.Add(strings.ToLower(id))
	}
}

// Trace emits one JSON-formatted log entry if trace level logging is enabled
func Trace(msg string, args ...any) {
	if level.Level() == LevelTrace {
		var pcs [1]uintptr
		runtime.Callers(2, pcs[:]) // skip [Callers, Infof]
		r := slog.NewRecord(time.Now(), LevelTrace, msg, pcs[0])
		r.Add(args...)
		_ = config.traceLogger.Handler().Handle(context.Background(), r)
	}
}

// TraceID emits one JSON-formatted log entry if tracing is enabled for the requested ID
func TraceID(id string, msg string, args ...any) {
	if level.Level() == LevelTrace && (config.traceIds.Contains(strings.ToLower(id)) || config.traceIds.Contains("all")) {
		var pcs [1]uintptr
		runtime.Callers(2, pcs[:]) // skip [Callers, Infof]
		r := slog.NewRecord(time.Now(), LevelTrace, msg, pcs[0])
		r.Add(args...)
		_ = config.traceLogger.Handler().Handle(context.Background(), r)
	}
}

// TraceIDs returns the list of enabled trace IDs
func TraceIDs() []string {
	return config.traceIds.ToSlice()
}

// Warn emits a warning log
func Warn(msg string, args ...any) {
	slog.Warn(msg, args...)
}
