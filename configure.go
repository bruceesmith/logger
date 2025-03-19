package logger

import (
	"fmt"
	"io"
	"log/slog"
	"os"

	set "github.com/deckarep/golang-set/v2"
)

//go:generate go tool stringer -type LogID
//go:generate go tool stringer -type SettingKey

// loggerConfig is the modifiable settings of a logger
type loggerConfig struct {
	Destination io.Writer
	Format      Format
	OmitTime    bool
}

// configuration of this package
type configuration struct {
	Normal      loggerConfig
	Trace       loggerConfig
	traceIds    set.Set[string]
	traceLogger *slog.Logger
}

var (
	config                   configuration
	defaultNormalDestination = os.Stdout
	defaultTraceDestination  = os.Stderr
	level                    slog.LevelVar
)

// Format determines the format of each log entry
type Format string

const (
	Text Format = "text" // Text format logs
	JSON Format = "json" // JSON format logs
)

// LogID defines the identifier of a logger
type LogID int

const (
	Norm  LogID = iota // The normal logger
	Tracy              // The trace logger
)

// SettingKey defines a logger setting that can be set or
// changed via the Configure function
type SettingKey int

const (
	DestinationSetting SettingKey = iota // Output writer / destination for a logger
	FormatSetting                        // Format of log entries
	OmitTimeSetting                      // Whether a timestamp is included in log entries
)

// ConfigSetting is an argument to Configure()
type ConfigSetting struct {
	AppliesTo LogID      // Logger whose setting is set/changed
	Key       SettingKey // Attribute of the logger that is srt/changed
	Value     any        // New value for this attribute
}

func init() {
	config = configuration{
		Normal: loggerConfig{
			Destination: os.Stdout,
			Format:      Text,
			OmitTime:    false,
		},
		Trace: loggerConfig{
			Destination: os.Stderr,
			Format:      Text,
			OmitTime:    false,
		},
		traceIds: set.NewSet[string](),
		traceLogger: slog.New(
			textHandler(defaultTraceDestination, true),
		),
	}
	level.Set(slog.LevelInfo)
	slog.SetDefault(
		slog.New(
			textHandler(defaultNormalDestination, false),
		),
	)
}

// Configure sets or changes attributes of either the normal
// or trace loggers
func Configure(setting ...ConfigSetting) error {
	for _, s := range setting {
		switch s.AppliesTo {
		case Norm, Tracy:
		default:
			return fmt.Errorf("there is no logger identied as %s", s.AppliesTo.String())
		}
		switch s.Key {
		case DestinationSetting:
			w, ok := s.Value.(io.Writer)
			if !ok {
				return fmt.Errorf("unknown destination %v", s.Value)
			}
			destination(s.AppliesTo, w)
		case FormatSetting:
			f, ok := s.Value.(Format)
			if !ok {
				return fmt.Errorf("unknown logger Format value %v", s.Value)
			}
			formats(s.AppliesTo, f)
		case OmitTimeSetting:
			b, ok := s.Value.(bool)
			if !ok {
				return fmt.Errorf("unknown imit time value %v", s.Value)
			}
			omitTime(s.AppliesTo, b)
		default:
			return fmt.Errorf("there is no configuration setting called %s", s.Key.String())
		}
	}
	return nil
}

// formats adjusts the format (JSON or text) of loggers
func formats(log LogID, f Format) {
	switch log {
	case Norm:
		if f != config.Normal.Format {
			normalFormat(f)
		}
	case Tracy:
		if f != config.Trace.Format {
			traceFormat(f)
		}
	}
}

// normalFormat adjusts the format (JSON or text) of the normal logger
func normalFormat(f Format) {
	config.Normal.Format = f
	switch f {
	case JSON:
		slog.SetDefault(slog.New(jsonHandler(config.Normal.Destination, false)))
	case Text:
		slog.SetDefault(slog.New(textHandler(config.Normal.Destination, false)))
	}
}

// traceFormat adjusts the format (JSON or text) of the trace logger
func traceFormat(f Format) {
	config.Trace.Format = f
	switch f {
	case JSON:
		config.traceLogger = slog.New(jsonHandler(config.Trace.Destination, true))
	case Text:
		config.traceLogger = slog.New(textHandler(config.Trace.Destination, true))
	}
}

// destination adjusts the output writer of loggers
func destination(log LogID, w io.Writer) {
	switch log {
	case Norm:
		if w != config.Normal.Destination {
			normalDestination(w)
		}
	case Tracy:
		if w != config.Trace.Destination {
			traceDestination(w)
		}
	}
}

// normalDestination adjusts the normal logger's writer
func normalDestination(w io.Writer) {
	config.Normal.Destination = w
	switch config.Normal.Format {
	case JSON:
		slog.SetDefault(slog.New(jsonHandler(w, false)))
	case Text:
		slog.SetDefault(slog.New(textHandler(w, false)))
	}
}

// traceDestination adjusts the trace logger's writer
func traceDestination(w io.Writer) {
	config.Trace.Destination = w
	switch config.Trace.Format {
	case JSON:
		config.traceLogger = slog.New(jsonHandler(w, true))
	case Text:
		config.traceLogger = slog.New(textHandler(w, true))
	}
}

// omitTime determines if either logger will include timestamps
func omitTime(log LogID, omit bool) {
	switch log {
	case Norm:
		config.Normal.OmitTime = omit
	case Tracy:
		config.Trace.OmitTime = omit
	}
}
