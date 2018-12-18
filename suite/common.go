// Copyright 2015-2018 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source tree.

package suite

import (
	"net/url"
	"strings"
)

// ----------------------------------------------------------------------
//
// Private Functions
//
// ----------------------------------------------------------------------

/*
startTest - This method will prepare the environment for a test
*/
func (s *Suite) startTest() {
	s.resetHeader()
	s.Req.URL.RawQuery = ""
}

/*
resetHeader - This method will reset the HTTP headers back to a clean state
*/
func (s *Suite) resetHeader() {
	s.Req.Header = make(map[string][]string)
}

func (s *Suite) enableAuth(u, p string) {
	s.Req.SetBasicAuth(u, p)
}

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
}

/*

 */
func (s *Suite) makePrettyQueryParams() string {
	pretty, _ := url.QueryUnescape(s.Req.URL.RawQuery)
	return pretty
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
printTestSummary - This function will print out a summary of the number of errors
found in a specific test.
*/
func (s *Suite) printTestSummary() {
	if s.ProblemsFound == 0 {
		s.Logger.Println("== SUCCESS: This test completed successfully\n")
	} else if s.ProblemsFound == 1 {
		s.Logger.Println("== FAILURE:", s.ProblemsFound, "problem found in this test\n")
	} else if s.ProblemsFound > 1 {
		s.Logger.Println("== FAILURE:", s.ProblemsFound, "problems found in this test\n")
	}
	s.ProblemsFound = 0
}
