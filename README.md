<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# logger

```go
import "github.com/bruceesmith/logger"
```

Package logger supports logging and tracing based on the standard library package [log/slog](<https://pkg.go.dev/log/slog/>).

[goreference\\\_badge](<https://pkg.go.dev/badge/github.com/bruceesmith/logger/v3.svg>)[goreference\\\_link](<https://pkg.go.dev/github.com/bruceesmith/logger>) [goreportcard\\\_badge](<https://goreportcard.com/badge/github.com/bruceesmith/logger>)[goreportcard\\\_link](<https://goreportcard.com/report/github.com/bruceesmith/logger>)

Debug, Error, Info and Warn operate like their package slog equivalents, with the level of logging modifiable using SetLevel.

A custom logging level \(LevelTrace\) can be supplied to SetLevel to enable tracing. Tracing can be unconditional when calling Trace, or only enabled for pre\-defined identifiers when calling TraceID. Identifiers for TraceID are registered by calling SetTraceIDs.

By default, all debug, error, info and warn messages go to Stdout, and traces go to Stderr; these destinations can be changed by calling RedirectNormal and RedirectTrace respectively.

When used in [cli applications](<https://github.com/urfave/cli>), a cli.Flag representing a LogLevel can be provided using the LogLevelFlag type.

## Index

- [func Debug\(msg string, args ...any\)](<#Debug>)
- [func Error\(msg string, args ...any\)](<#Error>)
- [func Info\(msg string, args ...any\)](<#Info>)
- [func Level\(\) string](<#Level>)
- [func RedirectStandard\(w io.Writer\)](<#RedirectStandard>)
- [func RedirectTrace\(w io.Writer\)](<#RedirectTrace>)
- [func SetFormat\(f Format\)](<#SetFormat>)
- [func SetLevel\(l slog.Level\)](<#SetLevel>)
- [func SetTraceIds\(ids ...string\)](<#SetTraceIds>)
- [func Trace\(msg string, args ...any\)](<#Trace>)
- [func TraceID\(id string, msg string, args ...any\)](<#TraceID>)
- [func Warn\(msg string, args ...any\)](<#Warn>)
- [type Format](<#Format>)
- [type LogLevel](<#LogLevel>)
  - [func \(ll \*LogLevel\) Set\(ls string\) \(err error\)](<#LogLevel.Set>)
  - [func \(ll \*LogLevel\) String\(\) \(s string\)](<#LogLevel.String>)
  - [func \(ll \*LogLevel\) Type\(\) string](<#LogLevel.Type>)
  - [func \(ll \*LogLevel\) UnmarshalJSON\(jason \[\]byte\) \(err error\)](<#LogLevel.UnmarshalJSON>)
- [type LogLevelFlag](<#LogLevelFlag>)
- [type Traces](<#Traces>)
  - [func \(t \*Traces\) Set\(ts string\) \(err error\)](<#Traces.Set>)
  - [func \(t \*Traces\) String\(\) \(s string\)](<#Traces.String>)
  - [func \(t \*Traces\) Type\(\) string](<#Traces.Type>)


<a name="Debug"></a>
## func Debug

```go
func Debug(msg string, args ...any)
```

Debug emits a debug log

<a name="Error"></a>
## func Error

```go
func Error(msg string, args ...any)
```

Error emits an error log

<a name="Info"></a>
## func Info

```go
func Info(msg string, args ...any)
```

Info emits an info log

<a name="Level"></a>
## func Level

```go
func Level() string
```



<a name="RedirectStandard"></a>
## func RedirectStandard

```go
func RedirectStandard(w io.Writer)
```

RedirectStandard changes the destination for normal \(non\-trace\) logs

<a name="RedirectTrace"></a>
## func RedirectTrace

```go
func RedirectTrace(w io.Writer)
```

RedirectTrace changes the destination for normal \(non\-trace\) logs

<a name="SetFormat"></a>
## func SetFormat

```go
func SetFormat(f Format)
```

SetFormat changes the format of log entries

<a name="SetLevel"></a>
## func SetLevel

```go
func SetLevel(l slog.Level)
```

SetLevel sets the default level of logging

<a name="SetTraceIds"></a>
## func SetTraceIds

```go
func SetTraceIds(ids ...string)
```

SetTraceIds registers identifiers for future tracing

<a name="Trace"></a>
## func Trace

```go
func Trace(msg string, args ...any)
```

Trace emits one JSON\-formatted log entry if trace level logging is enabled

<a name="TraceID"></a>
## func TraceID

```go
func TraceID(id string, msg string, args ...any)
```

TraceID emits one JSON\-formatted log entry if tracing is enabled for the requested ID

<a name="Warn"></a>
## func Warn

```go
func Warn(msg string, args ...any)
```

Warn emits a warning log

<a name="Format"></a>
## type Format

Format determines the format of each log entry

```go
type Format string
```

<a name="LevelTrace"></a>

```go
const (
    // LevelTrace can be set to enable tracing
    LevelTrace slog.Level = -10
    // Text format
    Text Format = "text"
    // JSON format
    JSON Format = "json"
)
```

<a name="LogLevel"></a>
## type LogLevel

LogLevel is the level of logging

```go
type LogLevel int
```

<a name="LogLevel.Set"></a>
### func \(\*LogLevel\) Set

```go
func (ll *LogLevel) Set(ls string) (err error)
```

Set is a convenience method for pflag.Value

<a name="LogLevel.String"></a>
### func \(\*LogLevel\) String

```go
func (ll *LogLevel) String() (s string)
```

String is a convenience method for pflag.Value

<a name="LogLevel.Type"></a>
### func \(\*LogLevel\) Type

```go
func (ll *LogLevel) Type() string
```

Type is a conveniene method for pflag.Value

<a name="LogLevel.UnmarshalJSON"></a>
### func \(\*LogLevel\) UnmarshalJSON

```go
func (ll *LogLevel) UnmarshalJSON(jason []byte) (err error)
```

UnmarshalJSON is a convenience method for Kong

<a name="LogLevelFlag"></a>
## type LogLevelFlag

LogLevelFlag is useful for using a LogLevel as a command\-line flag in CLI applications

```go
type LogLevelFlag = cli.FlagBase[LogLevel, cli.NoConfig, logLevelValue]
```

<a name="Traces"></a>
## type Traces

Traces is the list of trace IDs enabled

```go
type Traces []string
```

<a name="Traces.Set"></a>
### func \(\*Traces\) Set

```go
func (t *Traces) Set(ts string) (err error)
```

Set is a convenience method for pflag.Value

<a name="Traces.String"></a>
### func \(\*Traces\) String

```go
func (t *Traces) String() (s string)
```

String is a convenience method for pflag.Value

<a name="Traces.Type"></a>
### func \(\*Traces\) Type

```go
func (t *Traces) Type() string
```

Type is a conveniene method for pflag.Value

Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)
