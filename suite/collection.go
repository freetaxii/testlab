// Copyright 2018 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package suite

import (
	"encoding/json"
	"io/ioutil"

	"github.com/freetaxii/libstix2/resources"
)

/*
TestROCollectionService - This method will perform all of the standard tests
against the Read-Only Collection endpoint. It will also check to make sure the
output from the GET request is correct and will echo the output to the logs.
*/
func (s *Suite) TestROCollectionService() {
	s.Logger.Println("== Testing Read-Only Collection Service")

	path := s.APIRoot + "collections/" + s.ReadOnly + "/"
	s.setPath(path)
	s.EndpointType = "taxii"

	s.basicEndpointTests()

	s.Logger.Println("== Test C2: Test successful response from read-only collection endpoint")
	if s.Verbose {
		s.Logger.Println("++ This test will check to see if a proper read-only collection resource is returned")
	}
	c := GenerateROCollection()
	s.testCollectionResponse(c)
}

/*
TestWOCollectionService - This method will perform all of the standard tests
against the Write-Only Collection endpoint. It will also check to make sure the
output from the GET request is correct and will echo the output to the logs.
*/
func (s *Suite) TestWOCollectionService() {
	s.Logger.Println("== Testing Write-Only Collection Service")

	path := s.APIRoot + "collections/" + s.WriteOnly + "/"
	s.setPath(path)
	s.EndpointType = "taxii"

	s.basicEndpointTests()

	s.Logger.Println("== Test C3: Test successful response from write-only collection endpoint")
	if s.Verbose {
		s.Logger.Println("++ This test will check to see if a proper write-only collection resource is returned")
	}
	c := GenerateWOCollection()
	s.testCollectionResponse(c)
}

/*
TestRWCollectionService - This method will perform all of the standard tests
against the Read-Write Collection endpoint. It will also check to make sure the
output from the GET request is correct and will echo the output to the logs.
*/
func (s *Suite) TestRWCollectionService() {
	s.Logger.Println("== Testing Read-Write Collection Service")

	path := s.APIRoot + "collections/" + s.ReadWrite + "/"
	s.setPath(path)
	s.EndpointType = "taxii"

	s.basicEndpointTests()

	s.Logger.Println("== Test C4: Test successful response from read-write collection endpoint")
	if s.Verbose {
		s.Logger.Println("++ This test will check to see if a proper read-write collection resource is returned")
	}
	c := GenerateRWCollection()
	s.testCollectionResponse(c)
}

/*
testCollectionResponse - This method is used by other tests that will test
to ensure that the correct objects are returned.
*/
func (s *Suite) testCollectionResponse(c *resources.Collection) {
	if s.Verbose {
		s.Logger.Println("++ Calling Path:", s.Req.URL.Path)
	}

	media := s.TAXIIMediaType + s.TAXIIVersion
	s.setAccept(media)

	s.Req.SetBasicAuth(s.Username, s.Password)
	resp, err := s.Client.Do(s.Req)
	s.handleError(err)
	defer resp.Body.Close()
	s.ProblemsFound += s.checkResponseCode(resp.StatusCode, 200)

	body, err := ioutil.ReadAll(resp.Body)
	s.handleError(err)

	var o resources.Collection
	jerr := json.Unmarshal(body, &o)
	s.handleError(jerr)

	if valid, problems, details := c.Compare(&o); valid != true {
		s.ProblemsFound += problems
		if s.Debug {
			for _, v := range details {
				s.Logger.Println(v)
			}
		}
		s.Logger.Println("-- ERROR: Returned collection", c.ID, "does not match expected")

	} else {
		if s.Debug {
			for _, v := range details {
				s.Logger.Println(v)
			}
		}
		if s.Verbose {
			s.Logger.Println("++ Returned collection", c.ID, "matches expected")
		}
	}

	if s.Debug {
		var data []byte
		data, _ = json.MarshalIndent(o, "", "    ")
		s.Logger.Println("++ Collection Resource Returned:\n", string(data))
	}

	s.printSummary()
	s.reset()
}
