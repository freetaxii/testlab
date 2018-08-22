// Copyright 2018 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package suite

import (
	"crypto/tls"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/gologme/log"
)

type Workbench struct {
	Verbose      bool
	Debug        bool
	Username     string
	Password     string
	URL          string
	Proxy        string
	Discovery    string
	APIRoot      string
	OldMediaType bool
	ReadOnly     string
	WriteOnly    string
	ReadWrite    string
}

func NewWorkbench() *Workbench {
	var wb Workbench
	return &wb
}

type Suite struct {
	Logger          *log.Logger
	Req             *http.Request
	Client          *http.Client
	ProblemsFound   int
	EndpointType    string
	STIXMediaType   string
	TAXIIMediaType  string
	STIXVersion     string
	TAXIIVersion    string
	AcceptMediaType string
	Workbench
}

func NewSuite(logger *log.Logger, wb *Workbench) *Suite {
	var s Suite
	var err error
	var netTransport *http.Transport

	if logger == nil {
		s.Logger = log.New(os.Stderr, "", log.LstdFlags)
	} else {
		s.Logger = logger
	}

	if wb.Proxy != "" {
		proxyURL, err := url.Parse(wb.Proxy)
		if err != nil {
			s.Logger.Fatalln(err)
		}

		netTransport = &http.Transport{
			Dial: (&net.Dialer{
				Timeout: 5 * time.Second,
			}).Dial,
			TLSHandshakeTimeout: 5 * time.Second,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
			Proxy: http.ProxyURL(proxyURL),
		}
	} else {
		netTransport = &http.Transport{
			Dial: (&net.Dialer{
				Timeout: 5 * time.Second,
			}).Dial,
			TLSHandshakeTimeout: 5 * time.Second,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}
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
	s.Debug = wb.Debug
	s.OldMediaType = wb.OldMediaType
	s.ReadOnly = wb.ReadOnly
	s.WriteOnly = wb.WriteOnly
	s.ReadWrite = wb.ReadWrite

	if wb.OldMediaType {
		s.TAXIIMediaType = "application/vnd.oasis.taxii+json"
		s.STIXMediaType = "application/vnd/oasis.stix+json"
	} else {
		s.TAXIIMediaType = "application/taxii+json"
		s.STIXMediaType = "application/stix+json"
	}

	if s.OldMediaType {
		s.STIXVersion = ";version=2.0"
		s.TAXIIVersion = ";version=2.0"
	} else {
		s.STIXVersion = ";version=2.1"
		s.TAXIIVersion = ";version=2.1"
	}

	s.Req, err = http.NewRequest(http.MethodGet, s.URL, nil)
	s.handleError(err)
	return &s
}

// ----------------------------------------------------------------------
//
// Private Functions
//
// ----------------------------------------------------------------------

func (s *Suite) setPath(p string) {
	if !strings.HasPrefix(s.Req.URL.Path, "/") {
		p = "/" + p
	}

	if !strings.HasSuffix(s.Req.URL.Path, "/") {
		p = p + "/"
	}
	s.Req.URL.Path = p
}

/*
setAccept - This function will set the accept header to the string provided
*/
func (s *Suite) setAccept(accept string) {
	s.Req.Header.Add("Accept", accept)
	s.AcceptMediaType = accept
}

/*
resetHeader - This function will clear out all settings from the HTTP header
*/
func (s *Suite) resetHeader() {
	s.Req.Header = make(map[string][]string)
}

/*
resetQueryParams - This method will clear out all query parameters from the
HTTP header
*/
func (s *Suite) resetQueryParams() {
	s.Req.URL.RawQuery = ""
}

/*

 */
func (s *Suite) makePrettyQueryParams() string {
	pretty, _ := url.QueryUnescape(s.Req.URL.RawQuery)
	return pretty
}

/*
reset - This function will clear out all settings from the HTTP header and the
problems found
*/
func (s *Suite) reset() {
	s.resetHeader()
	s.resetQueryParams()
	s.ProblemsFound = 0
}

/*
checkResponseCode - This function will verify the actual HTTP response code
against one more more possible expected response codes. It will return an integer
representing the number of problems found.
*/
func (s *Suite) checkResponseCode(actual int, expected ...int) int {
	if len(expected) >= 2 {
		if expected[0] != actual && expected[1] != actual {
			s.Logger.Printf("-- ERROR: Expected HTTP response code %d. Got %d\n", expected[0], actual)
			return 1
		}
	} else if len(expected) == 1 {
		if expected[0] != actual {
			s.Logger.Printf("-- ERROR: Expected HTTP response code %d. Got %d\n", expected[0], actual)
			return 1
		}
	} else {
		s.Logger.Fatalln("-- FATAL: Missing expected HTTP code")
	}
	return 0
}

/*
checkContentType - This function will verify the actual HTTP response
content-type is correct. It will return an integer representing the number of
problems found.
*/
func (s *Suite) checkContentType(actual string, expected string) int {
	if expected != actual {
		s.Logger.Printf("-- ERROR: Expected HTTP content type %s. Got %s\n", expected, actual)
		return 1
	}
	return 0
}

/*
handleError - This function will test the Go errors that come back from other
function calls. This prevents us from having to put the if statement everywhere.
*/
func (s *Suite) handleError(err error) {
	if err != nil {
		s.Logger.Fatalln(err)
	}
}

/*
printSummary - This function will print out a summary of the number of errors
found in a specific test.
*/
func (s *Suite) printSummary() {
	if s.ProblemsFound == 0 {
		s.Logger.Println("== SUCCESS: This test completed successfully\n")
	} else if s.ProblemsFound == 1 {
		s.Logger.Println("== FAILURE:", s.ProblemsFound, "problem found in this test\n")
	} else if s.ProblemsFound > 1 {
		s.Logger.Println("== FAILURE:", s.ProblemsFound, "problems found in this test\n")
	}
}
