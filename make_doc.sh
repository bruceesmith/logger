#!/bin/bash
echo '[![Go Reference][goreference_badge]][goreference_link]' > temp
echo '[![Go Report Card][goreportcard_badge]][goreportcard_link]' >> temp
echo '[goreference_badge]: https://pkg.go.dev/badge/github.com/bruceesmith/logger/v3.svg' >> temp
echo '[goreference_link]: https://pkg.go.dev/github.com/bruceesmith/logger' >> temp
echo '[goreportcard_badge]: https://goreportcard.com/badge/github.com/bruceesmith/logger' >> temp
echo '[goreportcard_link]: https://goreportcard.com/report/github.com/bruceesmith/logger' >> temp
go run github.com/princjef/gomarkdoc/cmd/gomarkdoc@latest ./... --output read
cat temp read > README.md
rm temp read
