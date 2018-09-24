// Copyright 2018 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package suite

import (
	"github.com/freetaxii/libstix2/objects/bundle"
	"github.com/freetaxii/libstix2/objects/indicator"
)

/*
TestObjectsServiceROCollection - This method will perform all of the standard tests
against the Read-Only Objects endpoint. It will also check to make sure the
output from the GET request is correct and will echo the output to the logs.
*/
func (s *Suite) TestObjectsServiceROCollection() {
	s.Logger.Println("## ---------------------------------------------------------")
	s.Logger.Println("## Testing Objects Service Read-Only Collection")
	s.Logger.Println("## ---------------------------------------------------------")

	path := s.APIRoot + "collections/" + s.ReadOnly + "/objects/"
	s.setPath(path)
	s.EndpointType = "stix"

	s.basicEndpointTests()
	s.basicIndicatorFilteringTestsObjectsRO()
}

/*
TestObjectServiceROCollection - This method will perform all of the standard tests
against the Read-Only Objects endpoint. It will also check to make sure the
output from the GET request is correct and will echo the output to the logs.
*/
func (s *Suite) TestObjectServiceROCollection() {
	s.Logger.Println("## ---------------------------------------------------------")
	s.Logger.Println("## Testing Object Service Object By ID Read-Only Collection")
	s.Logger.Println("## ---------------------------------------------------------")

	allIndicators := GenerateIndicatorData()
	path := s.APIRoot + "collections/" + s.ReadOnly + "/objects/" + allIndicators[0].ID + "/"
	s.setPath(path)
	s.EndpointType = "stix"

	s.basicEndpointTests()
	s.basicIndicatorFilteringTestsObjectRO()
}

/*
testSortOrder01 - This method will get all indicators from the read-only
collection and make sure they are all correct.
*/
func (s *Suite) testSortOrder01() {
	s.Logger.Println("## Test SO-01: Test Sort Order")
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

	b, err := bundle.DecodeRaw(resp.Body)
	if err != nil {
		s.Logger.Println("-- ERROR: Invalid bundle returned", err)
		s.ProblemsFound++
		s.printTestSummary()
		s.reset()
		return
	}

	allIndicators := GenerateIndicatorData()
	// This first test will only have 2 indicators
	indicators := []indicator.Indicator{allIndicators[4], allIndicators[5]}

	for index, v := range b.Objects {

		// Make a first pass to decode just the object type value. Once we have this
		// value we can easily make a second pass and decode the rest of the object.
		stixtype, err := bundle.DecodeObjectType(v)
		if err != nil {
			// We should probably log the error here
			continue
		}

		switch stixtype {
		case "indicator":
			o, _, err := indicator.Decode(v)
			if err != nil {
				// We should probably log the error here
				continue
			}

			// Test sort order.
			if o.ID != indicators[index].ID {
				s.Logger.Println("-- ERROR: Sort order for returned data is wrong needs to be ascending")
				s.ProblemsFound++
				continue
			}
		}
	}
	s.printTestSummary()
	s.reset()
}
