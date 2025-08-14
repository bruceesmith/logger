module github.com/bruceesmith/logger

go 1.25

require (
	github.com/deckarep/golang-set/v2 v2.8.0
	github.com/urfave/cli/v3 v3.4.1
)

require (
	github.com/BurntSushi/toml v1.5.0 // indirect
	github.com/gojp/goreportcard v0.0.0-20250418060254-1060522058eb // indirect
	golang.org/x/exp/typeparams v0.0.0-20250813145105-42675adae3e6 // indirect
	golang.org/x/mod v0.27.0 // indirect
	golang.org/x/sync v0.16.0 // indirect
	golang.org/x/sys v0.35.0 // indirect
	golang.org/x/telemetry v0.0.0-20250813145757-41cd51e6ab6a // indirect
	golang.org/x/tools v0.36.0 // indirect
	golang.org/x/tools/go/expect v0.1.1-deprecated // indirect
	golang.org/x/tools/go/packages/packagestest v0.1.1-deprecated // indirect
	golang.org/x/vuln v1.1.4 // indirect
	honnef.co/go/tools v0.6.1 // indirect
)

tool (
	github.com/gojp/goreportcard/cmd/goreportcard-cli
	golang.org/x/tools/cmd/stringer
	golang.org/x/vuln/cmd/govulncheck
	honnef.co/go/tools/cmd/staticcheck
)
