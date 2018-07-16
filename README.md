# FreeTAXII/testlab #

[![Go Report Card](https://goreportcard.com/badge/github.com/freetaxii/testlab)](https://goreportcard.com/report/github.com/freetaxii/testlab) [![GoDoc](https://godoc.org/github.com/freetaxii/testlab?status.png)](https://godoc.org/github.com/freetaxii/testlab)

testlab contains a series of testing tools to automatically test a TAXII 2.x 
implementation. It was written in the Go (Golang) programming language.

## Test Tools ##

### basicTests.go ###
This tool will perform basic connectivity tests against every
endpoint. It will check media types for both the Accept and Content-Type headers
and will verify that the endpoint returns the right resource.

### getContentTests.go ###
This tool will perform various GET requests to the object
endpoints for read-only collections. It will test filtering of the data and 
ensure all of the object level endpoints return the right results. This test 
tool requires the following:
1) A read-only collection (22f763c1-e478-4765-8635-e4c32db665ea)
2) The provided STIX data (data/read-only.stix) will need to be loaded in to 
this read-only collection. 

### addContentTests.go ###
This tool will perform various POST requests to the object
endpoints for write-only and read-write collections. It will then perform 
various GET requests for the data on the read-write collection to ensure the 
added data is correctly preserved. This test tool requires the following:
1) A write-only collection (4f7327e2-f5b4-4269-b6e0-3564d174ce69)
2) A read-write collection (8c49f14d-8ea3-4f03-ab28-19dbca973dde)


## Installation ##

This package can be installed with the go get command:

```
go get github.com/freetaxii/testlab
cd /opt/go/src/github.com/freetaxii/testlab
go build basicTests.go
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

Basic Discovery Endpoint Tests
- [x] All Basic Tests
- [x] Successful GET of Discovery Resource

Basic API Root Endpoint Tests
- [x] All Basic Tests
- [x] Successful GET of API Root Resource


## License ##

This is free software, licensed under the Apache License, Version 2.0.


## Copyright ##

Copyright 2018 Bret Jordan, All rights reserved.

