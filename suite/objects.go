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
	s.Logger.Println()
	s.Logger.Println("== Testing Objects Service Read-Only Collection")

	path := s.APIRoot + "collections/" + s.ReadOnly + "/objects/"
	s.setPath(path)
	s.EndpointType = "stix"

	s.basicEndpointTests()
	s.testSortOrder01()
	s.testFilterVer01()
	s.testFilterVer02()
	s.testFilterVer03()
	s.testFilterVer04()
	s.testFilterVer05()
	s.testFilterVer06()
	s.testFilterVer07()
	s.testFilterID01()
	s.testFilterID02()
	s.testFilterType01()
}

/*
testSortOrder01 - This method will get all indicators from the read-only
collection and make sure they are all correct.
*/
func (s *Suite) testSortOrder01() {
	s.Logger.Println("== Test O-01: Test Sort Order")
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

/*
testFilterVer01 - This method will ensure the correct indicators are returned from
the read-only collection. There should be two returned.
*/
func (s *Suite) testFilterVer01() {
	s.Logger.Println("== Test O-02: Test No Filtering")
	if s.Verbose {
		s.Logger.Println("++ This test will not apply any filters to the read-only collection")
	}

	allIndicators := GenerateIndicatorData()
	indicators := []objects.Indicator{allIndicators[4], allIndicators[5]}
	s.testResponse(indicators)
}

/*
testFilterVer02 - This method will ensure the correct indicators are returned from
the read-only collection. There should be six returned.
*/
func (s *Suite) testFilterVer02() {
	s.Logger.Println("== Test O-03: Test Version Filtering Using All")
	if s.Verbose {
		s.Logger.Println("++ This test will filter the read-only collection by versions using the all keyword")
	}

	path := s.APIRoot + "collections/" + s.ReadOnly + "/objects/"
	s.setPath(path)

	values := s.Req.URL.Query()
	values.Set("match[version]", "all")
	s.Req.URL.RawQuery = values.Encode()

	allIndicators := GenerateIndicatorData()
	indicators := []objects.Indicator{allIndicators[0], allIndicators[1], allIndicators[2], allIndicators[3], allIndicators[4], allIndicators[5]}
	s.testResponse(indicators)
}

/*
testFilterVer03 - This method will ensure the correct indicators are returned from
the read-only collection. There should be two returned.
*/
func (s *Suite) testFilterVer03() {
	s.Logger.Println("== Test O-04: Test Version Filtering Using First")
	if s.Verbose {
		s.Logger.Println("++ This test will filter the read-only collection by versions using the first keyword")
	}

	path := s.APIRoot + "collections/" + s.ReadOnly + "/objects/"
	s.setPath(path)

	values := s.Req.URL.Query()
	values.Set("match[version]", "first")
	s.Req.URL.RawQuery = values.Encode()

	allIndicators := GenerateIndicatorData()
	indicators := []objects.Indicator{allIndicators[0], allIndicators[5]}
	s.testResponse(indicators)
}

/*
testFilterVer04 - This method will ensure the correct indicators are returned from
the read-only collection. There should be two returned.
*/
func (s *Suite) testFilterVer04() {
	s.Logger.Println("== Test O-05: Test Version Filtering Using Last")
	if s.Verbose {
		s.Logger.Println("++ This test will filter the read-only collection by versions using the last keyword")
	}

	path := s.APIRoot + "collections/" + s.ReadOnly + "/objects/"
	s.setPath(path)

	values := s.Req.URL.Query()
	values.Set("match[version]", "last")
	s.Req.URL.RawQuery = values.Encode()

	allIndicators := GenerateIndicatorData()
	indicators := []objects.Indicator{allIndicators[4], allIndicators[5]}
	s.testResponse(indicators)
}

/*
testFilterVer05 - This method will ensure the correct indicators are returned from
the read-only collection. There should be three returned.
*/
func (s *Suite) testFilterVer05() {
	s.Logger.Println("== Test O-06: Test Version Filtering Using First,Last")
	if s.Verbose {
		s.Logger.Println("++ This test will filter the read-only collection by versions using the first and last keywords")
	}

	path := s.APIRoot + "collections/" + s.ReadOnly + "/objects/"
	s.setPath(path)

	values := s.Req.URL.Query()
	values.Set("match[version]", "first,last")
	s.Req.URL.RawQuery = values.Encode()

	allIndicators := GenerateIndicatorData()
	indicators := []objects.Indicator{allIndicators[0], allIndicators[4], allIndicators[5]}
	s.testResponse(indicators)
}

/*
testFilterVer06 - This method will ensure the correct indicators are returned from
the read-only collection. There should be one returned.
*/
func (s *Suite) testFilterVer06() {
	s.Logger.Println("== Test O-07: Test Version Filtering Using Specific Version")
	if s.Verbose {
		s.Logger.Println("++ This test will filter the read-only collection by version using the version 2018-08-08T01:52:01.234Z")
	}

	path := s.APIRoot + "collections/" + s.ReadOnly + "/objects/"
	s.setPath(path)

	values := s.Req.URL.Query()
	values.Set("match[version]", "2018-08-08T01:52:01.234Z")
	s.Req.URL.RawQuery = values.Encode()

	allIndicators := GenerateIndicatorData()
	indicators := []objects.Indicator{allIndicators[1]}
	s.testResponse(indicators)
}

/*
testFilterVer07 - This method will ensure the correct indicators are returned from
the read-only collection. There should be four returned.
*/
func (s *Suite) testFilterVer07() {
	s.Logger.Println("== Test O-08: Test Version Filtering Using Last,First,Version")
	if s.Verbose {
		s.Logger.Println("++ This test will filter the read-only collection by version using the last, first, and version")
	}

	path := s.APIRoot + "collections/" + s.ReadOnly + "/objects/"
	s.setPath(path)

	values := s.Req.URL.Query()
	values.Set("match[version]", "last,first,2018-08-08T01:53:01.345Z")
	s.Req.URL.RawQuery = values.Encode()

	allIndicators := GenerateIndicatorData()
	indicators := []objects.Indicator{allIndicators[0], allIndicators[2], allIndicators[4], allIndicators[5]}
	s.testResponse(indicators)
}

/*
testFilterID01 - This method will ensure the correct indicators
are returned from the read-only collection. There should be four returned.
*/
func (s *Suite) testFilterID01() {
	s.Logger.Println("== Test O-09: Test ID Filtering Using One ID")
	if s.Verbose {
		s.Logger.Println("++ This test will filter the read-only collection by ID using a single STIX ID")
	}

	path := s.APIRoot + "collections/" + s.ReadOnly + "/objects/"
	s.setPath(path)

	values := s.Req.URL.Query()
	values.Set("match[id]", "indicator--1efc6673-9d95-46c3-a09c-c29f926da9af")
	s.Req.URL.RawQuery = values.Encode()

	allIndicators := GenerateIndicatorData()
	indicators := []objects.Indicator{allIndicators[4]}
	s.testResponse(indicators)
}

/*
testFilterID02 - This method will ensure the correct indicators
are returned from the read-only collection. There should be four returned.
*/
func (s *Suite) testFilterID02() {
	s.Logger.Println("== Test O-10: Test ID Filtering Using Two IDs")
	if s.Verbose {
		s.Logger.Println("++ This test will filter the read-only collection by ID using two STIX IDs")
	}

	path := s.APIRoot + "collections/" + s.ReadOnly + "/objects/"
	s.setPath(path)

	values := s.Req.URL.Query()
	values.Set("match[id]", "indicator--1efc6673-9d95-46c3-a09c-c29f926da9af,indicator--213dea46-8750-4b8b-b988-aae8f86a62d6")
	s.Req.URL.RawQuery = values.Encode()

	allIndicators := GenerateIndicatorData()
	indicators := []objects.Indicator{allIndicators[4], allIndicators[5]}
	s.testResponse(indicators)
}

/*
testFilterType01 - This method will ensure the correct indicators are
returned from the read-only collection. There should be two returned.
*/
func (s *Suite) testFilterType01() {
	s.Logger.Println("== Test O-11: Test Type Filtering Using Indicator")
	if s.Verbose {
		s.Logger.Println("++ This test will filter the read-only collection by type using Indicator")
	}

	path := s.APIRoot + "collections/" + s.ReadOnly + "/objects/"
	s.setPath(path)

	values := s.Req.URL.Query()
	values.Set("match[type]", "indicator")
	s.Req.URL.RawQuery = values.Encode()

	allIndicators := GenerateIndicatorData()
	indicators := []objects.Indicator{allIndicators[4], allIndicators[5]}
	s.testResponse(indicators)
}

/*
testResponse - This method is used by other tests that will test filtering and
ensure that the correct objects are returned.
*/
func (s *Suite) testResponse(indicators []objects.Indicator) {
	if s.Verbose {
		s.Logger.Println("++ Calling Path:", s.Req.URL.Path)
		s.Logger.Println("++ Query Params:", s.makePrettyQueryParams())
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

	count := 0

	if len(b.Objects) > 0 {
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
					if s.Debug {
						for _, v := range details {
							s.Logger.Println(v)
						}
					}
					s.Logger.Println("-- Returned indicator", o.ID, "version", o.Modified, "does not match expected")
				} else {
					if s.Debug {
						for _, v := range details {
							s.Logger.Println(v)
						}
					}
					s.Logger.Println("++ Returned indicator", o.ID, "version", o.Modified, "matches expected")
				}
			}

			count++
		}

		s.Logger.Println("++ Number objects returned:", count)

		if s.Debug {
			data, _ := b.EncodeToString()
			s.Logger.Println("++ Bundle Resource Returned:\n", data)
		}
	}

	s.printSummary()
	s.reset()
}
