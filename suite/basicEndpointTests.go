// Copyright 2018 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package suite

import (
	"strings"
)

func (s *Suite) basicEndpointTests() {
	// Test no authentication
	// Test failed authentication
	// Test successful authentication
	// Test missing trailing slash
	// Test each invalid accept header
	// Test each valid accept header
	// Test content-type header
	s.testBE01()
	s.testBE02()
	s.testBE03()
	s.testBE04()
	s.testBE05()
	s.testBE06()
	s.testBE07()
}

func (s *Suite) testBE01() {
	s.Logger.Println("== Test BE-01: No authentication test")
	if s.Verbose {
		s.Logger.Println("++ This test will send an empty authentication parameter and will check to see if a 401 or 404 status code is returned")
		s.Logger.Println("++ Calling Path:", s.Req.URL.Path)
	}

	if s.EndpointType == "stix" {
		media := s.STIXMediaType + s.STIXVersion
		s.setAccept(media)
	} else {
		media := s.TAXIIMediaType + s.TAXIIVersion
		s.setAccept(media)
	}
	resp, err := s.Client.Do(s.Req)
	s.handleError(err)
	defer resp.Body.Close()
	s.ProblemsFound += s.checkResponseCode(resp.StatusCode, 401, 404)

	s.printSummary()
	s.reset()
}

func (s *Suite) testBE02() {
	s.Logger.Println("== Test BE-02: Wrong authentication test")
	if s.Verbose {
		s.Logger.Println("++ This test will send an incorrect authentication parameter and will check to see if a 401 or 404 status code is returned")
		s.Logger.Println("++ Calling Path:", s.Req.URL.Path)
	}

	if s.EndpointType == "stix" {
		media := s.STIXMediaType + s.STIXVersion
		s.setAccept(media)
	} else {
		media := s.TAXIIMediaType + s.TAXIIVersion
		s.setAccept(media)
	}

	s.Req.SetBasicAuth(s.Username, "foo")
	resp, err := s.Client.Do(s.Req)
	s.handleError(err)
	defer resp.Body.Close()
	s.ProblemsFound += s.checkResponseCode(resp.StatusCode, 401, 404)

	s.printSummary()
	s.reset()
}

func (s *Suite) testBE03() {
	s.Logger.Println("== Test BE-03: Test successful authentication")
	if s.Verbose {
		s.Logger.Println("++ This test will send a correct authentication parameter and will check to see if a 200 status code is returned")
		s.Logger.Println("++ Calling Path:", s.Req.URL.Path)
	}

	if s.EndpointType == "stix" {
		media := s.STIXMediaType + s.STIXVersion
		s.setAccept(media)
	} else {
		media := s.TAXIIMediaType + s.TAXIIVersion
		s.setAccept(media)
	}

	s.Req.SetBasicAuth(s.Username, s.Password)
	resp, err := s.Client.Do(s.Req)
	s.handleError(err)
	defer resp.Body.Close()
	s.ProblemsFound += s.checkResponseCode(resp.StatusCode, 200)

	s.printSummary()
	s.reset()
}

func (s *Suite) testBE04() {
	s.Logger.Println("== Test BE-04: Test missing trailing slash")
	if s.Verbose {
		s.Logger.Println("++ This test will request a URL with a missing trailing slash and check to see if a 404 status code is returned")
		s.Logger.Println("++ Calling Path:", s.Req.URL.Path)
	}

	if s.EndpointType == "stix" {
		media := s.STIXMediaType + s.STIXVersion
		s.setAccept(media)
	} else {
		media := s.TAXIIMediaType + s.TAXIIVersion
		s.setAccept(media)
	}

	// Save original path
	orig := s.Req.URL.Path

	s.Req.URL.Path = strings.TrimSuffix(s.Req.URL.Path, "/")

	s.Req.SetBasicAuth(s.Username, s.Password)
	resp, err := s.Client.Do(s.Req)
	s.handleError(err)
	defer resp.Body.Close()
	s.ProblemsFound += s.checkResponseCode(resp.StatusCode, 404)

	// Set it back
	s.Req.URL.Path = orig

	s.printSummary()
	s.reset()
}

func (s *Suite) testBE05() {
	s.Logger.Println("== Test BE-05: Test invalid media types in Accept")
	if s.Verbose {
		s.Logger.Println("++ This test will make a series of requests with invalid Accept media types and check to see if a 406 status code is returned")
		s.Logger.Println("++ Calling Path:", s.Req.URL.Path)
	}

	invalidAcceptHeaders := []string{"", "application/foo"}

	for _, v := range invalidAcceptHeaders {
		s.setAccept(v)
		s.Req.SetBasicAuth(s.Username, s.Password)

		resp, err := s.Client.Do(s.Req)
		s.handleError(err)
		defer resp.Body.Close()
		s.ProblemsFound += s.checkResponseCode(resp.StatusCode, 406)
		s.resetHeader()
	}

	s.printSummary()
	s.reset()
}

func (s *Suite) testBE06() {
	s.Logger.Println("== Test BE-06: Test valid media types in Accept")
	if s.Verbose {
		s.Logger.Println("++ This test will make a series of requests with valid Accept media types and check to see if a 200 status code is returned")
		s.Logger.Println("++ Calling Path:", s.Req.URL.Path)
	}

	var m1, m2 string
	if s.EndpointType == "stix" {
		m1 = s.STIXMediaType
		m2 = s.STIXMediaType + s.STIXVersion
	} else {
		m1 = s.TAXIIMediaType
		m2 = s.TAXIIMediaType + s.TAXIIVersion
	}

	validHeaders := []string{m1, m2}

	for _, v := range validHeaders {
		s.setAccept(v)
		s.Req.SetBasicAuth(s.Username, s.Password)

		resp, err := s.Client.Do(s.Req)
		s.handleError(err)
		defer resp.Body.Close()
		s.ProblemsFound += s.checkResponseCode(resp.StatusCode, 200)
		s.resetHeader()
	}

	s.printSummary()
	s.reset()
}

func (s *Suite) testBE07() {
	s.Logger.Println("== Test BE-07: Test valid media type in Content-type")
	if s.Verbose {
		s.Logger.Println("++ This test will make a series of requests with valid Accept media types and check to see if the correct media type is returned")
		s.Logger.Println("++ Calling Path:", s.Req.URL.Path)
	}

	var m1, m2 string
	if s.EndpointType == "stix" {
		m1 = s.STIXMediaType
		m2 = s.STIXMediaType + s.STIXVersion
	} else {
		m1 = s.TAXIIMediaType
		m2 = s.TAXIIMediaType + s.TAXIIVersion
	}

	validHeaders := []string{m1, m2}

	for _, v := range validHeaders {
		s.setAccept(v)
		s.Req.SetBasicAuth(s.Username, s.Password)

		resp, err := s.Client.Do(s.Req)
		s.handleError(err)
		defer resp.Body.Close()
		s.ProblemsFound += s.checkContentType(resp.Header.Get("Content-type"), m2)
		s.resetHeader()
	}

	s.printSummary()
	s.reset()
}
