// Copyright 2018 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package testing

import (
	"strings"

	"github.com/gologme/log"
)

func (s *Suite) BasicTests() {
	// Test no authentication
	// Test failed authentication
	// Test successful authentication
	// Test missing trailing slash
	// Test each invalid accept header
	// Test each valid accept header
	// Test the data that is returned
	s.test1()
	s.test2()
	s.test3()
	s.test4()
	s.test5()
	s.test6()
}

func (s *Suite) test1() {
	log.Println("== Test B1: No authentication test")
	if s.Verbose {
		log.Println("++ This test will send an empty authentication parameter and will check to see if a 401 or 404 status code is returned")
	}
	s.setAccept(s.TAXIIMediaType)
	resp, err := s.Client.Do(s.Req)
	s.testError(err)
	defer resp.Body.Close()
	s.ProblemsFound += s.checkResponseCode(resp.StatusCode, 401, 404)
	s.printSummary()
	s.reset()
}

func (s *Suite) test2() {
	log.Println("== Test B2: Wrong authentication test")
	if s.Verbose {
		log.Println("++ This test will send an incorrect authentication parameter and will check to see if a 401 or 404 status code is returned")
	}
	s.setAccept(s.TAXIIMediaType)
	s.Req.SetBasicAuth(s.Username, "foo")
	resp, err := s.Client.Do(s.Req)
	s.testError(err)
	defer resp.Body.Close()
	s.ProblemsFound += s.checkResponseCode(resp.StatusCode, 401, 404)
	s.printSummary()
	s.reset()
}

func (s *Suite) test3() {
	log.Println("== Test B3: Test successful authentication")
	if s.Verbose {
		log.Println("++ This test will send a correct authentication parameter and will check to see if a 200 status code is returned")
	}
	s.setAccept(s.TAXIIMediaType)
	s.Req.SetBasicAuth(s.Username, s.Password)
	resp, err := s.Client.Do(s.Req)
	s.testError(err)
	defer resp.Body.Close()
	s.ProblemsFound += s.checkResponseCode(resp.StatusCode, 200)
	s.printSummary()
	s.reset()
}

func (s *Suite) test4() {
	log.Println("== Test B4: Test missing trailing slash")
	if s.Verbose {
		log.Println("++ This test will request a URL with a missing trailing slash and check to see if a 404 status code is returned")
	}
	// Save original path
	orig := s.Req.URL.Path

	s.Req.URL.Path = strings.TrimSuffix(s.Req.URL.Path, "/")
	s.setAccept(s.TAXIIMediaType)
	s.Req.SetBasicAuth(s.Username, s.Password)
	resp, err := s.Client.Do(s.Req)
	s.testError(err)
	defer resp.Body.Close()
	s.ProblemsFound += s.checkResponseCode(resp.StatusCode, 404)
	s.printSummary()

	// Set it back
	s.Req.URL.Path = orig
	s.reset()

}

func (s *Suite) test5() {
	log.Println("== Test B5: Test invalid media types")
	if s.Verbose {
		log.Println("++ This test will make a series of requests with invalid Accept media types and check to see if a 406 status code is returned")
	}
	invalidHeaders := []string{"", "application/foo"}

	for _, v := range invalidHeaders {
		s.setAccept(v)
		s.Req.SetBasicAuth(s.Username, s.Password)
		resp, err := s.Client.Do(s.Req)
		s.testError(err)
		defer resp.Body.Close()
		s.ProblemsFound += s.checkResponseCode(resp.StatusCode, 406)
		s.resetHeader()
	}
	s.printSummary()
	s.reset()
}

func (s *Suite) test6() {
	log.Println("== Test B6: Test valid media types")
	if s.Verbose {
		log.Println("++ This test will make a series of requests with valid Accept media types and check to see if a 200 status code is returned")
	}

	m1 := s.TAXIIMediaType
	m2 := s.TAXIIMediaType + s.MediaVersion

	validHeaders := []string{m1, m2}

	for _, v := range validHeaders {
		s.setAccept(v)
		s.Req.SetBasicAuth(s.Username, s.Password)
		resp, err := s.Client.Do(s.Req)
		s.testError(err)
		defer resp.Body.Close()
		s.ProblemsFound += s.checkResponseCode(resp.StatusCode, 200)
		s.resetHeader()
	}
	s.printSummary()
	s.reset()
}
