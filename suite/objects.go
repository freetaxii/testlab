// Copyright 2018 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package suite

import (
	"github.com/freetaxii/libstix2/objects"
)

/*
TestROCollectionObjectsService - This method will perform all of the standard tests
against the Read-Only Objects endpoint. It will also check to make sure the
output from the GET request is correct and will echo the output to the logs.
*/
func (s *Suite) TestROCollectionObjectsService() {
	path := s.APIRoot + "/collections/" + s.ReadOnly + "/objects"
	s.setPath(path)
	s.EndpointType = "stix"
	s.Logger.Println()
	s.Logger.Println("== Testing Read-Only Objects Service")
	if s.Verbose {
		s.Logger.Println("++ Calling Path:", s.Req.URL.Path)
	}

	s.basicEndpointTests()
	s.testSortOrder()
	s.getROCollectionObjectsOutput()
}

/*
testSortOrder - This method will get all indicators from the
read-only collection and make sure they are all correct.
*/
func (s *Suite) testSortOrder() {
	s.Logger.Println("== Test O1: Test sort order of response from read-only objects endpoint")
	if s.Verbose {
		s.Logger.Println("++ This test will check to see if the objects are in the right order")
	}

	media := s.STIXMediaType + s.STIXVersion
	s.setAccept(media)

	s.Req.SetBasicAuth(s.Username, s.Password)
	resp, err := s.Client.Do(s.Req)
	s.handleError(err)
	defer resp.Body.Close()
	s.ProblemsFound += s.checkResponseCode(resp.StatusCode, 200)

	b, err := objects.DecodeBundle(resp.Body)
	s.handleError(err)

	allIndicators := GenerateIndicatorData()
	// This first test will only have 2 indicators
	indicators := []objects.Indicator{allIndicators[4], allIndicators[5]}

	for index, v := range b.Objects {

		// Make a first pass to decode just the object type value. Once we have this
		// value we can easily make a second pass and decode the rest of the object.
		stixtype, err := objects.DecodeObjectType(v)
		if err != nil {
			// We should probably log the error here
			continue
		}

		switch stixtype {
		case "indicator":
			o, err := objects.DecodeIndicator(v)
			if err != nil {
				// We should probably log the error here
				continue
			}

			// Test sort order.
			if o.ID != indicators[index].ID {
				s.Logger.Println("ERROR: Sort order for returned data is wrong needs to be ascending")
				s.ProblemsFound++
				continue
			}
		}
	}
	s.printSummary()
	s.reset()
}

/*
getROCollectionObjectsOutput - This method will get all indicators from the
read-only collection and make sure they are all correct.
*/
func (s *Suite) getROCollectionObjectsOutput() {
	s.Logger.Println("== Test O2: Test successful response from read-only objects endpoint")
	if s.Verbose {
		s.Logger.Println("++ This test will check to see if a proper object resource is returned")
	}

	media := s.STIXMediaType + s.STIXVersion
	s.setAccept(media)

	s.Req.SetBasicAuth(s.Username, s.Password)
	resp, err := s.Client.Do(s.Req)
	s.handleError(err)
	defer resp.Body.Close()
	s.ProblemsFound += s.checkResponseCode(resp.StatusCode, 200)

	b, err := objects.DecodeBundle(resp.Body)
	s.handleError(err)

	count := 0
	allIndicators := GenerateIndicatorData()

	// This first test will only have 2 indicators, the newest of each version
	indicators := []objects.Indicator{allIndicators[4], allIndicators[5]}

	for index, v := range b.Objects {

		// Make a first pass to decode just the object type value. Once we have this
		// value we can easily make a second pass and decode the rest of the object.
		stixtype, err := objects.DecodeObjectType(v)
		if err != nil {
			// We should probably log the error here
			continue
		}

		switch stixtype {
		case "indicator":
			o, err := objects.DecodeIndicator(v)
			if err != nil {
				// We should probably log the error here
				continue
			}

			if valid, problems, details := indicators[index].Compare(o); valid != true {
				s.ProblemsFound += problems
				if s.Verbose {
					for _, v := range details {
						s.Logger.Println(v)
					}
				}
				s.Logger.Println("ERROR: Returned indicator", o.ID, "does not match expected")
			} else {
				if s.Verbose {
					for _, v := range details {
						s.Logger.Println(v)
					}
				}
				s.Logger.Println("SUCCESS: Returned indicator", o.ID, "matches expected")
			}
		}

		count++
	}

	s.Logger.Println("++ Number objects returned:", count)

	data, _ := b.EncodeToString()
	s.Logger.Println("++ Bundle Resource Returned:\n", data)

	s.printSummary()
	s.reset()
}
