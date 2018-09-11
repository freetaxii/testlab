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

func (s *Suite) basicFilteringTestsObjectsRO() {
	// Test no filtering
	// Test version filtering using ALL
	// Test version filtering using FIRST
	// Test version filtering using LAST
	// Test version filtering using FIRST,LAST
	// Test version filtering using VERSION
	// Test version filtering using LAST,FIRST,VERSION
	// Test id filtering using single ID
	// Test id filtering using two IDs
	// Test type filtering using single Type
	allIndicators := GenerateIndicatorData()

	s.testFT01([]indicator.Indicator{allIndicators[4], allIndicators[5]})
	s.testFT02([]indicator.Indicator{allIndicators[0], allIndicators[1], allIndicators[2], allIndicators[3], allIndicators[4], allIndicators[5]})
	s.testFT03([]indicator.Indicator{allIndicators[0], allIndicators[5]})
	s.testFT04([]indicator.Indicator{allIndicators[4], allIndicators[5]})
	s.testFT05([]indicator.Indicator{allIndicators[0], allIndicators[4], allIndicators[5]})
	s.testFT06([]indicator.Indicator{allIndicators[1]})
	s.testFT07([]indicator.Indicator{allIndicators[0], allIndicators[2], allIndicators[4], allIndicators[5]})
	s.testFT08([]indicator.Indicator{allIndicators[4]})
	s.testFT09([]indicator.Indicator{allIndicators[4], allIndicators[5]})
	s.testFT10([]indicator.Indicator{allIndicators[4], allIndicators[5]})
}

func (s *Suite) basicFilteringTestsObjectRO() {
	// Test no filtering
	// Test version filtering using ALL
	// Test version filtering using FIRST
	// Test version filtering using LAST
	// Test version filtering using FIRST,LAST
	// Test version filtering using VERSION
	// Test version filtering using LAST,FIRST,VERSION
	allIndicators := GenerateIndicatorData()

	s.testFT01([]indicator.Indicator{allIndicators[4]})
	s.testFT02([]indicator.Indicator{allIndicators[0], allIndicators[1], allIndicators[2], allIndicators[3], allIndicators[4]})
	s.testFT03([]indicator.Indicator{allIndicators[0]})
	s.testFT04([]indicator.Indicator{allIndicators[4]})
	s.testFT05([]indicator.Indicator{allIndicators[0], allIndicators[4]})
	s.testFT06([]indicator.Indicator{allIndicators[1]})
	s.testFT07([]indicator.Indicator{allIndicators[0], allIndicators[2], allIndicators[4]})
}

/*
testFT01 - This method will ensure the correct indicators are returned from
the read-only collection. There should be two returned.
*/
func (s *Suite) testFT01(indicators []indicator.Indicator) {
	s.Logger.Println("== Test FT-01: Test No Filtering")
	if s.Verbose {
		s.Logger.Println("++ This test will not apply any filters to the read-only collection")
	}
	s.testFilteringResponse(indicators)
}

/*
testFT02 - This method will ensure the correct indicators are returned from
the read-only collection. There should be six returned.
*/
func (s *Suite) testFT02(indicators []indicator.Indicator) {
	s.Logger.Println("== Test FT-02: Test Version Filtering Using All")
	if s.Verbose {
		s.Logger.Println("++ This test will filter the read-only collection by versions using the all keyword")
	}

	values := s.Req.URL.Query()
	values.Set("match[version]", "all")
	s.Req.URL.RawQuery = values.Encode()
	s.testFilteringResponse(indicators)
}

/*
testFT03 - This method will ensure the correct indicators are returned from
the read-only collection. There should be two returned.
*/
func (s *Suite) testFT03(indicators []indicator.Indicator) {
	s.Logger.Println("== Test FT-03: Test Version Filtering Using First")
	if s.Verbose {
		s.Logger.Println("++ This test will filter the read-only collection by versions using the first keyword")
	}

	values := s.Req.URL.Query()
	values.Set("match[version]", "first")
	s.Req.URL.RawQuery = values.Encode()
	s.testFilteringResponse(indicators)
}

/*
testFT04 - This method will ensure the correct indicators are returned from
the read-only collection. There should be two returned.
*/
func (s *Suite) testFT04(indicators []indicator.Indicator) {
	s.Logger.Println("== Test FT-04: Test Version Filtering Using Last")
	if s.Verbose {
		s.Logger.Println("++ This test will filter the read-only collection by versions using the last keyword")
	}

	values := s.Req.URL.Query()
	values.Set("match[version]", "last")
	s.Req.URL.RawQuery = values.Encode()
	s.testFilteringResponse(indicators)
}

/*
testFT05 - This method will ensure the correct indicators are returned from
the read-only collection. There should be three returned.
*/
func (s *Suite) testFT05(indicators []indicator.Indicator) {
	s.Logger.Println("== Test FT-05: Test Version Filtering Using First,Last")
	if s.Verbose {
		s.Logger.Println("++ This test will filter the read-only collection by versions using the first and last keywords")
	}

	values := s.Req.URL.Query()
	values.Set("match[version]", "first,last")
	s.Req.URL.RawQuery = values.Encode()
	s.testFilteringResponse(indicators)
}

/*
testFT06 - This method will ensure the correct indicators are returned from
the read-only collection. There should be one returned.
*/
func (s *Suite) testFT06(indicators []indicator.Indicator) {
	s.Logger.Println("== Test FT-06: Test Version Filtering Using Specific Version")
	if s.Verbose {
		s.Logger.Println("++ This test will filter the read-only collection by version using the version 2018-08-08T01:52:01.234Z")
	}

	values := s.Req.URL.Query()
	values.Set("match[version]", "2018-08-08T01:52:01.234Z")
	s.Req.URL.RawQuery = values.Encode()
	s.testFilteringResponse(indicators)
}

/*
testFT07 - This method will ensure the correct indicators are returned from
the read-only collection. There should be four returned.
*/
func (s *Suite) testFT07(indicators []indicator.Indicator) {
	s.Logger.Println("== Test FT-07: Test Version Filtering Using Last,First,Version")
	if s.Verbose {
		s.Logger.Println("++ This test will filter the read-only collection by version using the last, first, and version")
	}

	values := s.Req.URL.Query()
	values.Set("match[version]", "last,first,2018-08-08T01:53:01.345Z")
	s.Req.URL.RawQuery = values.Encode()
	s.testFilteringResponse(indicators)
}

/*
testFT08 - This method will ensure the correct indicators
are returned from the read-only collection. There should be four returned.
*/
func (s *Suite) testFT08(indicators []indicator.Indicator) {
	s.Logger.Println("== Test FT-08: Test ID Filtering Using One ID")
	if s.Verbose {
		s.Logger.Println("++ This test will filter the read-only collection by ID using a single STIX ID")
	}

	values := s.Req.URL.Query()
	values.Set("match[id]", "indicator--1efc6673-9d95-46c3-a09c-c29f926da9af")
	s.Req.URL.RawQuery = values.Encode()
	s.testFilteringResponse(indicators)
}

/*
testFT09 - This method will ensure the correct indicators
are returned from the read-only collection. There should be four returned.
*/
func (s *Suite) testFT09(indicators []indicator.Indicator) {
	s.Logger.Println("== Test FT-09: Test ID Filtering Using Two IDs")
	if s.Verbose {
		s.Logger.Println("++ This test will filter the read-only collection by ID using two STIX IDs")
	}

	values := s.Req.URL.Query()
	values.Set("match[id]", "indicator--1efc6673-9d95-46c3-a09c-c29f926da9af,indicator--213dea46-8750-4b8b-b988-aae8f86a62d6")
	s.Req.URL.RawQuery = values.Encode()
	s.testFilteringResponse(indicators)
}

/*
testFT10 - This method will ensure the correct indicators are
returned from the read-only collection. There should be two returned.
*/
func (s *Suite) testFT10(indicators []indicator.Indicator) {
	s.Logger.Println("== Test FT-10: Test Type Filtering Using Indicator")
	if s.Verbose {
		s.Logger.Println("++ This test will filter the read-only collection by type using Indicator")
	}

	values := s.Req.URL.Query()
	values.Set("match[type]", "indicator")
	s.Req.URL.RawQuery = values.Encode()
	s.testFilteringResponse(indicators)
}

/*
testFilteringResponse - This method is used by other tests that will test filtering and
ensure that the correct objects are returned.
*/
func (s *Suite) testFilteringResponse(indicators []indicator.Indicator) {
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

	b, err := bundle.Decode(resp.Body)
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
			stixtype, err := bundle.DecodeObjectType(v)
			if err != nil {
				// We should probably log the error here
				continue
			}

			switch stixtype {
			case "indicator":
				o, err := indicator.Decode(v)
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
					s.Logger.Println("-- ERROR: Returned indicator", o.ID, "version", o.Modified, "does not match expected")

				} else {
					if s.Debug {
						for _, v := range details {
							s.Logger.Println(v)
						}
					}
					if s.Verbose {
						s.Logger.Println("++ Returned indicator", o.ID, "version", o.Modified, "matches expected")
					}
				}
			}

			count++
		}

		if s.Verbose {
			s.Logger.Println("++ Number objects returned:", count)
		}

		if s.Debug {
			data, _ := b.EncodeToString()
			s.Logger.Println("++ Bundle Resource Returned:\n", data)
		}
	}

	s.printSummary()
	s.reset()
}
