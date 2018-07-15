// Copyright 2018 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package testing

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"

	"github.com/gologme/log"
)

type Workbench struct {
	Verbose      bool
	Username     string
	Password     string
	URL          string
	Discovery    string
	APIRoot      string
	OldMediaType bool
}

func NewWorkbench() *Workbench {
	var wb Workbench
	return &wb
}

type Suite struct {
	Req            *http.Request
	Client         *http.Client
	ProblemsFound  int
	STIXMediaType  string
	TAXIIMediaType string
	MediaVersion   string
	Workbench
}

func NewSuite(wb *Workbench) *Suite {
	var s Suite
	var err error
	netTransport := &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	s.Client = &http.Client{
		Timeout:   time.Second * 10,
		Transport: netTransport,
	}

	s.URL = wb.URL
	s.Discovery = wb.Discovery
	s.APIRoot = wb.APIRoot
	s.Username = wb.Username
	s.Password = wb.Password
	s.Verbose = wb.Verbose
	s.OldMediaType = wb.OldMediaType

	if wb.OldMediaType {
		s.TAXIIMediaType = "application/vnd.oasis.taxii+json"
		s.STIXMediaType = "application/vnd/oasis.stix+json"
	} else {
		s.TAXIIMediaType = "application/taxii+json"
		s.STIXMediaType = "application/stix+json"
	}

	if s.OldMediaType {
		s.MediaVersion = ";version=2.0"
	} else {
		s.MediaVersion = ";version=2.1"
	}

	s.Req, err = http.NewRequest(http.MethodGet, s.URL, nil)
	s.testError(err)
	return &s
}

// ----------------------------------------------------------------------
//
// Private Functions
//
// ----------------------------------------------------------------------

func (s *Suite) setPath(p string) {
	s.Req.URL.Path = "/" + p + "/"
}

/*
setAccept - This function will set the accept header to the string provided
*/
func (s *Suite) setAccept(accept string) {
	s.Req.Header.Add("Accept", accept)
}

/*
resetHeader - This function will clear out all settings from the HTTP header
*/
func (s *Suite) resetHeader() {
	s.Req.Header = make(map[string][]string)
}

/*
reset - This function will clear out all settings from the HTTP header and the
problems found
*/
func (s *Suite) reset() {
	s.resetHeader()
	s.ProblemsFound = 0
}

/*
checkResponseCode - This function will verify the actual HTTP response code
against one more more possible expected response codes.
*/
func (s *Suite) checkResponseCode(actual int, expected ...int) int {
	if len(expected) >= 2 {
		if expected[0] != actual && expected[1] != actual {
			log.Printf("ERROR: Expected HTTP response code %d. Got %d\n", expected[0], actual)
			return 1
		}
	} else if len(expected) == 1 {
		if expected[0] != actual {
			log.Printf("ERROR: Expected HTTP response code %d. Got %d\n", expected[0], actual)
			return 1
		}
	} else {
		log.Fatalln("FATAL: Missing expected HTTP code")
	}
	return 0
}

/*
testError - This function will test the Go errors that come back from other
function calls. This prevents us from having to put the if statement everywhere.
*/
func (s *Suite) testError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

/*
printSummary - This function will print out a summary of the number of errors
found in a specific test.
*/
func (s *Suite) printSummary() {
	if s.ProblemsFound == 0 {
		//log.Println("SUCCESS: This test completed successfully")
	} else if s.ProblemsFound == 1 {
		log.Println("ERROR:", s.ProblemsFound, "problem found in this test")
	} else if s.ProblemsFound > 1 {
		log.Println("ERROR:", s.ProblemsFound, "problems found in this test")
	}
}
