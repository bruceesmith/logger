module github.com/bruceesmith/logger

go 1.26

require (
	github.com/deckarep/golang-set/v2 v2.8.0
	github.com/urfave/cli/v3 v3.7.0
)

require (
	github.com/BurntSushi/toml v1.6.0 // indirect
	github.com/gojp/goreportcard v0.0.0-20260302084821-2ae51f4baa2e // indirect
	golang.org/x/exp/typeparams v0.0.0-20260312153236-7ab1446f8b90 // indirect
	golang.org/x/mod v0.34.0 // indirect
	golang.org/x/sync v0.20.0 // indirect
	golang.org/x/sys v0.42.0 // indirect
	golang.org/x/telemetry v0.0.0-20260316223853-b6b0c46d1ccd // indirect
	golang.org/x/tools v0.43.0 // indirect
	golang.org/x/tools/go/packages/packagestest v0.1.1-deprecated // indirect
	golang.org/x/vuln v1.1.4 // indirect
	honnef.co/go/tools v0.7.0 // indirect
)

tool (
	github.com/gojp/goreportcard/cmd/goreportcard-cli
	golang.org/x/tools/cmd/stringer
	golang.org/x/vuln/cmd/govulncheck
	honnef.co/go/tools/cmd/staticcheck
)
