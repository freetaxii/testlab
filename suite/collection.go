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
	path := s.APIRoot + "/collections/" + s.ReadOnly
	s.setPath(path)
	s.Logger.Println()
	s.Logger.Println("== Testing Read-Only Collection Service")
	if s.Verbose {
		s.Logger.Println("++ Calling Path:", s.Req.URL.Path)
	}
	s.basicTests()

	s.Logger.Println("== Test C2: Test successful response from read-only collection endpoint")
	if s.Verbose {
		s.Logger.Println("++ This test will check to see if a proper read-only collection resource is returned")
	}

	ro := GenerateROCollection()
	var o resources.Collection
	media := s.TAXIIMediaType + s.MediaVersion
	s.setAccept(media)
	s.Req.SetBasicAuth(s.Username, s.Password)
	resp, err := s.Client.Do(s.Req)
	s.testError(err)
	defer resp.Body.Close()
	s.ProblemsFound += s.checkResponseCode(resp.StatusCode, 200)

	body, err := ioutil.ReadAll(resp.Body)
	s.testError(err)

	jerr := json.Unmarshal(body, &o)
	s.testError(jerr)

	if valid := s.compareCollections(*ro, o); valid != true {
		s.Logger.Println("ERROR: Returned collection does not match expected read-only collection")
	}

	var data []byte
	data, _ = json.MarshalIndent(o, "", "    ")
	s.Logger.Println("++ Collection Resource Returned:\n", string(data))

	s.printSummary()
	s.reset()
}

/*
TestWOCollectionService - This method will perform all of the standard tests
against the Write-Only Collection endpoint. It will also check to make sure the
output from the GET request is correct and will echo the output to the logs.
*/
func (s *Suite) TestWOCollectionService() {
	path := s.APIRoot + "/collections/" + s.WriteOnly
	s.setPath(path)
	s.Logger.Println()
	s.Logger.Println("== Testing Write-Only Collection Service")
	if s.Verbose {
		s.Logger.Println("++ Calling Path:", s.Req.URL.Path)
	}
	s.basicTests()

	s.Logger.Println("== Test C3: Test successful response from write-only collection endpoint")
	if s.Verbose {
		s.Logger.Println("++ This test will check to see if a proper write-only collection resource is returned")
	}

	wo := GenerateWOCollection()
	var o resources.Collection
	media := s.TAXIIMediaType + s.MediaVersion
	s.setAccept(media)
	s.Req.SetBasicAuth(s.Username, s.Password)
	resp, err := s.Client.Do(s.Req)
	s.testError(err)
	defer resp.Body.Close()
	s.ProblemsFound += s.checkResponseCode(resp.StatusCode, 200)

	body, err := ioutil.ReadAll(resp.Body)
	s.testError(err)

	jerr := json.Unmarshal(body, &o)
	s.testError(jerr)

	if valid := s.compareCollections(*wo, o); valid != true {
		s.Logger.Println("ERROR: Returned collection does not match expected write-only collection")
	}

	var data []byte
	data, _ = json.MarshalIndent(o, "", "    ")
	s.Logger.Println("++ Collection Resource Returned:\n", string(data))

	s.printSummary()
	s.reset()
}

/*
TestRWCollectionService - This method will perform all of the standard tests
against the Read-Write Collection endpoint. It will also check to make sure the
output from the GET request is correct and will echo the output to the logs.
*/
func (s *Suite) TestRWCollectionService() {
	path := s.APIRoot + "/collections/" + s.ReadWrite
	s.setPath(path)
	s.Logger.Println()
	s.Logger.Println("== Testing Read-Write Collection Service")
	if s.Verbose {
		s.Logger.Println("++ Calling Path:", s.Req.URL.Path)
	}

	s.basicTests()

	s.Logger.Println("== Test C4: Test successful response from read-write collection endpoint")
	if s.Verbose {
		s.Logger.Println("++ This test will check to see if a proper read-write collection resource is returned")
	}

	rw := GenerateRWCollection()
	var o resources.Collection
	media := s.TAXIIMediaType + s.MediaVersion
	s.setAccept(media)
	s.Req.SetBasicAuth(s.Username, s.Password)
	resp, err := s.Client.Do(s.Req)
	s.testError(err)
	defer resp.Body.Close()
	s.ProblemsFound += s.checkResponseCode(resp.StatusCode, 200)

	body, err := ioutil.ReadAll(resp.Body)
	s.testError(err)

	jerr := json.Unmarshal(body, &o)
	s.testError(jerr)

	if valid := s.compareCollections(*rw, o); valid != true {
		s.Logger.Println("ERROR: Returned collection does not match expected read-write collection")
	}

	var data []byte
	data, _ = json.MarshalIndent(o, "", "    ")
	s.Logger.Println("++ Collection Resource Returned:\n", string(data))

	s.printSummary()
	s.reset()
}

/*
compareCollections - This method will compare two collections to make sure they
are the same. Collection c1 represent the correct data, c2 represents what was
retrieved from a server. So we need to make sure that c2 is the same as c1.
*/
func (s *Suite) compareCollections(c1, c2 resources.Collection) bool {

	// Check ID Value
	if c2.ID != c1.ID {
		s.ProblemsFound++
	} else {
		if s.Verbose {
			s.Logger.Println("++ IDs Match:", c1.ID, "|", c2.ID)
		}
	}

	// Check Title Value
	if c2.Title != c1.Title {
		s.ProblemsFound++
	} else {
		if s.Verbose {
			s.Logger.Println("++ Titles Match:", c1.Title, "|", c2.Title)
		}
	}

	// Check Description Value
	if c2.Description != c1.Description {
		s.ProblemsFound++
	} else {
		if s.Verbose {
			s.Logger.Println("++ Descriptions Match:", c1.Description, "|", c2.Description)
		}
	}

	// Check Can Read Value
	if c2.CanRead != c1.CanRead {
		s.ProblemsFound++
	} else {
		if s.Verbose {
			s.Logger.Println("++ Can Read Values Match:", c1.CanRead, "|", c2.CanRead)
		}
	}

	// Check Can Write Value
	if c2.CanWrite != c1.CanWrite {
		s.ProblemsFound++
	} else {
		if s.Verbose {
			s.Logger.Println("++ Can Write Values Match:", c1.CanWrite, "|", c2.CanWrite)
		}
	}

	// Check Media Type Property Length
	if len(c2.MediaTypes) != len(c1.MediaTypes) {
		s.ProblemsFound++
	} else {
		if s.Verbose {
			s.Logger.Println("++ Media Type Length Match:", len(c1.MediaTypes), "|", len(c2.MediaTypes))
		}
	}

	// Check Media Type values
	if c2.MediaTypes[0] != c1.MediaTypes[0] {
		s.ProblemsFound++
	} else {
		if s.Verbose {
			s.Logger.Println("++ Media Types Match:", c1.MediaTypes[0], "|", c2.MediaTypes[0])
		}
	}

	if s.ProblemsFound > 0 {
		s.Logger.Printf("ERROR: Returned collection does not match expected value")
		s.Logger.Printf("ERROR: Expected %s", c1)
		s.Logger.Printf("ERROR: Got %s", c2)
		return false
	}

	return true
}
