module github.com/bruceesmith/logger

go 1.26

require (
	github.com/deckarep/golang-set/v2 v2.9.0
	github.com/urfave/cli/v3 v3.10.1
)

require (
	github.com/BurntSushi/toml v1.6.0 // indirect
	github.com/gojp/goreportcard v0.0.0-20260605163032-af15decf135b // indirect
	go.mongodb.org/mongo-driver v1.17.9 // indirect
	golang.org/x/exp/typeparams v0.0.0-20260727155853-b88d891fe743 // indirect
	golang.org/x/mod v0.38.0 // indirect
	golang.org/x/sync v0.22.0 // indirect
	golang.org/x/sys v0.47.0 // indirect
	golang.org/x/telemetry v0.0.0-20260804195142-bdd03c3c8848 // indirect
	golang.org/x/tools v0.48.0 // indirect
	golang.org/x/vuln v1.6.0 // indirect
	honnef.co/go/tools v0.7.0 // indirect
)

tool (
	github.com/gojp/goreportcard/cmd/goreportcard-cli
	golang.org/x/tools/cmd/stringer
	golang.org/x/vuln/cmd/govulncheck
	honnef.co/go/tools/cmd/staticcheck
)
