module github.com/bruceesmith/logger

go 1.24

require (
	github.com/bruceesmith/set v1.0.0
	github.com/urfave/cli/v3 v3.0.0-beta1
)

require (
	github.com/BurntSushi/toml v1.4.1-0.20240526193622-a339e1f7089c // indirect
	github.com/gojp/goreportcard v0.0.0-20241203091639-8d0356773220 // indirect
	golang.org/x/exp/typeparams v0.0.0-20231108232855-2478ac86f678 // indirect
	golang.org/x/mod v0.23.0 // indirect
	golang.org/x/sync v0.11.0 // indirect
	golang.org/x/sys v0.30.0 // indirect
	golang.org/x/telemetry v0.0.0-20240522233618-39ace7a40ae7 // indirect
	golang.org/x/tools v0.30.0 // indirect
	golang.org/x/vuln v1.1.4 // indirect
	honnef.co/go/tools v0.6.0 // indirect
)

tool (
	github.com/gojp/goreportcard/cmd/goreportcard-cli
	golang.org/x/tools/cmd/stringer
	golang.org/x/vuln/cmd/govulncheck
	honnef.co/go/tools/cmd/staticcheck
)
