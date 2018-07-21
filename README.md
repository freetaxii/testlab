# FreeTAXII/testlab #

[![Go Report Card](https://goreportcard.com/badge/github.com/freetaxii/testlab)](https://goreportcard.com/report/github.com/freetaxii/testlab) [![GoDoc](https://godoc.org/github.com/freetaxii/testlab?status.png)](https://godoc.org/github.com/freetaxii/testlab)

testlab contains a series of testing tools to automatically test a TAXII 2.x 
implementation. It was written in the Go (Golang) programming language and as
such each tool can be compiled into a standalone statically compiled native 
executable. 

## Test Setup ##
The three test tools listed below all require that an existing API Root be 
pre-configured on the TAXII Server. This API root will be used for all tests in
this suite. The name and path of this API root will need to be fed in to the 
test tools via a command line parameter. For example:

```
https://somesite.com/api1/
./basicTests -a api1

https://somesite.com/taxii/api1/
./basicTests -a taxii/api1
```

This API Root should be configured in one of two ways with the collections 
listed below. Configuration 2 the Read-Write Implementation is the most common 
and should be done by the majority of implementations.


### Configuration 1: Read Only Implementation ###
```
Read-only  Test Collection ID 22f763c1-e478-4765-8635-e4c32db665ea

{
    "id": "22f763c1-e478-4765-8635-e4c32db665ea",
    "title": "Read-Only TestLab Collection",
    "description": "This is a Read-Only collection for use with the FreeTAXII TestLab tool",
    "can_read": true,
    "can_write": false,
    "media_types": [
        "application/stix+json;version=2.1"
    ]
}
```

### Configuration 2: Read and Write Implementation ###
```
Read-only  Test Collection ID 22f763c1-e478-4765-8635-e4c32db665ea
Write-only Test Collection ID 4f7327e2-f5b4-4269-b6e0-3564d174ce69
Read-Write Test Collection ID 8c49f14d-8ea3-4f03-ab28-19dbca973dde

{
    "id": "22f763c1-e478-4765-8635-e4c32db665ea",
    "title": "Read-Only TestLab Collection",
    "description": "This is a Read-Only collection for use with the FreeTAXII TestLab tool",
    "can_read": true,
    "can_write": false,
    "media_types": [
        "application/stix+json;version=2.1"
    ]
}
{
    "id": "4f7327e2-f5b4-4269-b6e0-3564d174ce69",
    "title": "Write-Only TestLab Collection",
    "description": "This is a Write-Only collection for use with the FreeTAXII TestLab tool",
    "can_read": false,
    "can_write": true,
    "media_types": [
        "application/stix+json;version=2.1"
    ]
}
{
    "id": "8c49f14d-8ea3-4f03-ab28-19dbca973dde",
    "title": "Read-Write TestLab Collection",
    "description": "This is a Read-Write collection for use with the FreeTAXII TestLab tool",
    "can_read": true,
    "can_write": true,
    "media_types": [
        "application/stix+json;version=2.1"
    ]
}
```


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
- [x] Valid Media Type for Content-type

Basic Discovery Endpoint Tests
- [x] All Basic Tests
- [x] Successful GET of Discovery Resource

Basic API Root Endpoint Tests
- [x] All Basic Tests
- [x] Successful GET of API Root Resource

Basic Collections Endpoint Tests
- [x] All Basic Tests
- [x] Successful GET of Collections Resource


## License ##

This is free software, licensed under the Apache License, Version 2.0.


## Copyright ##

Copyright 2018 Bret Jordan, All rights reserved.

