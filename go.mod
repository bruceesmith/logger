module github.com/bruceesmith/logger

go 1.24.4

require (
	github.com/deckarep/golang-set/v2 v2.8.0
	github.com/urfave/cli/v3 v3.3.8
)

require (
	github.com/BurntSushi/toml v1.5.0 // indirect
	github.com/gojp/goreportcard v0.0.0-20250418060254-1060522058eb // indirect
	golang.org/x/exp/typeparams v0.0.0-20250620022241-b7579e27df2b // indirect
	golang.org/x/mod v0.25.0 // indirect
	golang.org/x/sync v0.15.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/telemetry v0.0.0-20250701142123-2f1f7721456e // indirect
	golang.org/x/tools v0.34.0 // indirect
	golang.org/x/vuln v1.1.4 // indirect
	honnef.co/go/tools v0.6.1 // indirect
)

tool (
	github.com/gojp/goreportcard/cmd/goreportcard-cli
	golang.org/x/tools/cmd/stringer
	golang.org/x/vuln/cmd/govulncheck
	honnef.co/go/tools/cmd/staticcheck
)
