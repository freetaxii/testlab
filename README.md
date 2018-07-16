# FreeTAXII/testlab #

[![Go Report Card](https://goreportcard.com/badge/github.com/freetaxii/testlab)](https://goreportcard.com/report/github.com/freetaxii/testlab) [![GoDoc](https://godoc.org/github.com/freetaxii/testlab?status.png)](https://godoc.org/github.com/freetaxii/testlab)

testlab contains a series of testing tools to automatically test a TAXII 2.x 
implementation. It was written in the Go (Golang) programming language.

## Installation ##

This package can be installed with the go get command:

```
go get github.com/freetaxii/testlab
cd /opt/go/src/github.com/freetaxii/testlab
go build testTAXII.go
```


## Tests ##

Below is a list of tests which have been implemented:

Basic Tests - These are run on every endpoint 
- [x] No Authentication Parameter
- [x] Incorrect Authentication Parameter
- [x] Correct Authentication Parameter
- [x] Missing Trailing Slash
- [x] Invalid Media Types for Accept
- [x] Valid Media Types for Accept

Discovery Endpoint Tests
- [x] All Basic Tests
- [x] Successful GET of Discovery Resource

API Root Endpoint Tests
- [x] All Basic Tests
- [x] Successful GET of API Root Resource


## License ##

This is free software, licensed under the Apache License, Version 2.0.


## Copyright ##

Copyright 2018 Bret Jordan, All rights reserved.

