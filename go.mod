module github.com/bruceesmith/logger

go 1.24

require (
	github.com/deckarep/golang-set/v2 v2.8.0
	github.com/urfave/cli/v3 v3.0.0-beta1
)

require (
	github.com/BurntSushi/toml v1.5.0 // indirect
	github.com/gojp/goreportcard v0.0.0-20250313055930-6c0fa925565c // indirect
	golang.org/x/exp/typeparams v0.0.0-20250305212735-054e65f0b394 // indirect
	golang.org/x/mod v0.24.0 // indirect
	golang.org/x/sync v0.12.0 // indirect
	golang.org/x/sys v0.31.0 // indirect
	golang.org/x/telemetry v0.0.0-20250310203348-fdfaad844314 // indirect
	golang.org/x/tools v0.31.0 // indirect
	golang.org/x/vuln v1.1.4 // indirect
	honnef.co/go/tools v0.6.1 // indirect
)

tool (
	github.com/gojp/goreportcard/cmd/goreportcard-cli
	golang.org/x/tools/cmd/stringer
	golang.org/x/vuln/cmd/govulncheck
	honnef.co/go/tools/cmd/staticcheck
)
