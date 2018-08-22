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
TestObjectsServiceROCollection - This method will perform all of the standard tests
against the Read-Only Objects endpoint. It will also check to make sure the
output from the GET request is correct and will echo the output to the logs.
*/
func (s *Suite) TestObjectsServiceROCollection() {
	s.Logger.Println()
	s.Logger.Println("== Testing Objects Service Read-Only Collection")

	path := s.APIRoot + "collections/" + s.ReadOnly + "/objects/"
	s.setPath(path)
	s.EndpointType = "stix"

	s.basicEndpointTests()
	s.basicFilteringTestsObjectsRO()
}

/*
testSortOrder01 - This method will get all indicators from the read-only
collection and make sure they are all correct.
*/
func (s *Suite) testSortOrder01() {
	s.Logger.Println("== Test SO-01: Test Sort Order")
	if s.Verbose {
		s.Logger.Println("++ This test will check to see if the sort order is correct for indicators returned from the read-only collection")
		s.Logger.Println("++ Calling Path:", s.Req.URL.Path)
	}

	media := s.STIXMediaType + s.STIXVersion
	s.setAccept(media)

	s.Req.SetBasicAuth(s.Username, s.Password)
	resp, err := s.Client.Do(s.Req)
	s.handleError(err)
	defer resp.Body.Close()
	s.ProblemsFound += s.checkResponseCode(resp.StatusCode, 200)

	b, err := objects.DecodeBundle(resp.Body)
	if err != nil {
		s.Logger.Println("-- ERROR: Invalid bundle returned", err)
		s.ProblemsFound++
		s.printSummary()
		s.reset()
		return
	}

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
