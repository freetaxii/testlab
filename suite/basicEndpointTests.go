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
	s.Logger.Println("## Start Basic Endpoint Tests\n")

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
	s.Logger.Println("## Test BE-01: No Authentication Test")
	s.Logger.Infoln("++ This test will send an empty authentication parameter and will check to see if a 401 or 404 status code is returned")
	s.Logger.Infoln("++ Calling Path:", s.Req.URL.Path)

	s.startTest()
	s.setAccept(s.FullMediaType)

	resp, err := s.Client.Do(s.Req)
	s.handleError(err)
	defer resp.Body.Close()
	s.ProblemsFound += s.checkResponseCode(resp.StatusCode, 401, 404)
	s.printTestSummary()
}

func (s *Suite) testBE02() {
	s.Logger.Println("## Test BE-02: Wrong Authentication Test")
	s.Logger.Infoln("++ This test will send an incorrect authentication parameter and will check to see if a 401 or 404 status code is returned")
	s.Logger.Infoln("++ Calling Path:", s.Req.URL.Path)

	s.startTest()
	s.setAccept(s.FullMediaType)
	s.enableAuth(s.Settings.Username, "foo")

	resp, err := s.Client.Do(s.Req)
	s.handleError(err)
	defer resp.Body.Close()
	s.ProblemsFound += s.checkResponseCode(resp.StatusCode, 401, 404)
	s.printTestSummary()
}

func (s *Suite) testBE03() {
	s.Logger.Println("## Test BE-03: Test Successful Authentication")
	s.Logger.Infoln("++ This test will send a correct authentication parameter and will check to see if a 200 status code is returned")
	s.Logger.Infoln("++ Calling Path:", s.Req.URL.Path)

	s.startTest()
	s.setAccept(s.FullMediaType)
	s.enableAuth(s.Settings.Username, s.Settings.Password)

	resp, err := s.Client.Do(s.Req)
	s.handleError(err)
	defer resp.Body.Close()
	s.ProblemsFound += s.checkResponseCode(resp.StatusCode, 200)

	s.printTestSummary()
}

func (s *Suite) testBE04() {
	s.Logger.Println("## Test BE-04: Test Missing Trailing Slash")
	s.Logger.Infoln("++ This test will request a URL with a missing trailing slash and check to see if a 404 status code is returned")
	s.Logger.Infoln("++ Calling Path:", s.Req.URL.Path)

	s.startTest()
	s.setAccept(s.FullMediaType)
	s.enableAuth(s.Settings.Username, s.Settings.Password)

	// Save original path
	orig := s.Req.URL.Path
	s.Req.URL.Path = strings.TrimSuffix(s.Req.URL.Path, "/")

	resp, err := s.Client.Do(s.Req)
	s.handleError(err)
	defer resp.Body.Close()
	s.ProblemsFound += s.checkResponseCode(resp.StatusCode, 404)

	// Set it back
	s.Req.URL.Path = orig

	s.printTestSummary()
}

func (s *Suite) testBE05() {
	s.Logger.Println("## Test BE-05: Test Invalid Accept Media Types")
	s.Logger.Infoln("++ This test will make a series of requests with invalid Accept media types and check to see if a 406 status code is returned")
	s.Logger.Infoln("++ Calling Path:", s.Req.URL.Path)

	invalidAcceptHeaders := []string{"", "application/foo"}

	for _, v := range invalidAcceptHeaders {
		s.startTest()
		s.setAccept(v)
		s.enableAuth(s.Settings.Username, s.Settings.Password)

		resp, err := s.Client.Do(s.Req)
		s.handleError(err)
		defer resp.Body.Close()
		s.ProblemsFound += s.checkResponseCode(resp.StatusCode, 406)
	}

	s.printTestSummary()
}

func (s *Suite) testBE06() {
	s.Logger.Println("## Test BE-06: Test Valid Accept Media Types")
	s.Logger.Infoln("++ This test will make a series of requests with valid Accept media types and check to see if a 200 status code is returned")
	s.Logger.Infoln("++ Calling Path:", s.Req.URL.Path)

	m1 := s.TAXIIMediaType
	m2 := s.FullMediaType
	validHeaders := []string{m1, m2}

	for _, v := range validHeaders {
		s.startTest()
		s.setAccept(v)
		s.enableAuth(s.Settings.Username, s.Settings.Password)

		resp, err := s.Client.Do(s.Req)
		s.handleError(err)
		defer resp.Body.Close()
		s.ProblemsFound += s.checkResponseCode(resp.StatusCode, 200)
	}

	s.printTestSummary()
}

func (s *Suite) testBE07() {
	s.Logger.Println("## Test BE-07: Test Valid Content-Type Media Type")
	s.Logger.Infoln("++ This test will make a series of requests with valid Accept media types and check to see if the correct media type is returned")
	s.Logger.Infoln("++ Calling Path:", s.Req.URL.Path)

	m1 := s.TAXIIMediaType
	m2 := s.FullMediaType
	validHeaders := []string{m1, m2}

	for _, v := range validHeaders {
		s.startTest()
		s.setAccept(v)
		s.enableAuth(s.Settings.Username, s.Settings.Password)

		resp, err := s.Client.Do(s.Req)
		s.handleError(err)
		defer resp.Body.Close()
		s.ProblemsFound += s.checkContentType(resp.Header.Get("Content-type"), m2)
	}

	s.printTestSummary()
}
