#!/bin/bash
echo '[![Go Reference][goreference_badge]][goreference_link]' >temp1
echo " " >>temp1
echo " " >temp2
echo '[goreference_badge]: https://pkg.go.dev/badge/github.com/bruceesmith/logger/v3.svg' >>temp2
echo '[goreference_link]: https://pkg.go.dev/github.com/bruceesmith/logger' >>temp2
go tool -modfile=tools/go.mod github.com/princjef/gomarkdoc/cmd/gomarkdoc ./... --output read
cat temp1 read temp2 >README.md
rm temp1 temp2 read
