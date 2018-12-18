// Copyright 2018 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package suite

import (
	"github.com/freetaxii/libstix2/objects"
	"github.com/freetaxii/libstix2/objects/indicator"
	"github.com/freetaxii/libstix2/resources/envelope"
)

func (s *Suite) basicIndicatorFilteringTestsObjectsRO() {
	s.Logger.Println("## Start Objects Filtering Tests for RO Collections\n")
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

	s.testFilter01([]indicator.Indicator{allIndicators[4], allIndicators[5]})
	s.testFilter02([]indicator.Indicator{allIndicators[0], allIndicators[1], allIndicators[2], allIndicators[3], allIndicators[4], allIndicators[5]})
	s.testFilter03([]indicator.Indicator{allIndicators[0], allIndicators[5]})
	s.testFilter04([]indicator.Indicator{allIndicators[4], allIndicators[5]})
	s.testFilter05([]indicator.Indicator{allIndicators[0], allIndicators[4], allIndicators[5]})
	s.testFilter06([]indicator.Indicator{allIndicators[1]})
	s.testFilter07([]indicator.Indicator{allIndicators[0], allIndicators[2], allIndicators[4], allIndicators[5]})
	s.testFilter08([]indicator.Indicator{allIndicators[4]})
	s.testFilter09([]indicator.Indicator{allIndicators[4], allIndicators[5]})
	s.testFilter10([]indicator.Indicator{allIndicators[4], allIndicators[5]})
}

func (s *Suite) basicIndicatorFilteringTestsObjectRO() {
	s.Logger.Println("## Start Object By ID Filtering Tests for RO Collections\n")
	// Test no filtering
	// Test version filtering using ALL
	// Test version filtering using FIRST
	// Test version filtering using LAST
	// Test version filtering using FIRST,LAST
	// Test version filtering using VERSION
	// Test version filtering using LAST,FIRST,VERSION
	allIndicators := GenerateIndicatorData()

	s.testFilter01([]indicator.Indicator{allIndicators[4]})
	s.testFilter02([]indicator.Indicator{allIndicators[0], allIndicators[1], allIndicators[2], allIndicators[3], allIndicators[4]})
	s.testFilter03([]indicator.Indicator{allIndicators[0]})
	s.testFilter04([]indicator.Indicator{allIndicators[4]})
	s.testFilter05([]indicator.Indicator{allIndicators[0], allIndicators[4]})
	s.testFilter06([]indicator.Indicator{allIndicators[1]})
	s.testFilter07([]indicator.Indicator{allIndicators[0], allIndicators[2], allIndicators[4]})
}

/*
testFilter01 - This method will ensure the correct indicators are returned from
the read-only collection. There should be two returned.
*/
func (s *Suite) testFilter01(indicators []indicator.Indicator) {
	s.Logger.Println("## Test Filter-01: Test No Filtering")
	s.Logger.Infoln("++ This test will not apply any filters to the read-only collection")

	s.startTest()
	s.setAccept(s.FullMediaType)
	s.enableAuth(s.Settings.Username, s.Settings.Password)
	s.testFilteringResponse(indicators)
}

/*
testFilter02 - This method will ensure the correct indicators are returned from
the read-only collection. There should be six returned.
*/
func (s *Suite) testFilter02(indicators []indicator.Indicator) {
	s.Logger.Println("## Test Filter-02: Test Version Filtering Using All")
	s.Logger.Infoln("++ This test will filter the read-only collection by versions using the all keyword")

	s.startTest()
	s.setAccept(s.FullMediaType)
	s.enableAuth(s.Settings.Username, s.Settings.Password)

	values := s.Req.URL.Query()
	values.Set("match[version]", "all")
	s.Req.URL.RawQuery = values.Encode()
	s.testFilteringResponse(indicators)
}

/*
testFilter03 - This method will ensure the correct indicators are returned from
the read-only collection. There should be two returned.
*/
func (s *Suite) testFilter03(indicators []indicator.Indicator) {
	s.Logger.Println("## Test Filter-03: Test Version Filtering Using First")
	s.Logger.Infoln("++ This test will filter the read-only collection by versions using the first keyword")

	s.startTest()
	s.setAccept(s.FullMediaType)
	s.enableAuth(s.Settings.Username, s.Settings.Password)

	values := s.Req.URL.Query()
	values.Set("match[version]", "first")
	s.Req.URL.RawQuery = values.Encode()
	s.testFilteringResponse(indicators)
}

/*
testFilter04 - This method will ensure the correct indicators are returned from
the read-only collection. There should be two returned.
*/
func (s *Suite) testFilter04(indicators []indicator.Indicator) {
	s.Logger.Println("## Test Filter-04: Test Version Filtering Using Last")
	s.Logger.Infoln("++ This test will filter the read-only collection by versions using the last keyword")

	s.startTest()
	s.setAccept(s.FullMediaType)
	s.enableAuth(s.Settings.Username, s.Settings.Password)

	values := s.Req.URL.Query()
	values.Set("match[version]", "last")
	s.Req.URL.RawQuery = values.Encode()
	s.testFilteringResponse(indicators)
}

/*
testFilter05 - This method will ensure the correct indicators are returned from
the read-only collection. There should be three returned.
*/
func (s *Suite) testFilter05(indicators []indicator.Indicator) {
	s.Logger.Println("## Test Filter-05: Test Version Filtering Using First,Last")
	s.Logger.Infoln("++ This test will filter the read-only collection by versions using the first and last keywords")

	s.startTest()
	s.setAccept(s.FullMediaType)
	s.enableAuth(s.Settings.Username, s.Settings.Password)

	values := s.Req.URL.Query()
	values.Set("match[version]", "first,last")
	s.Req.URL.RawQuery = values.Encode()
	s.testFilteringResponse(indicators)
}

/*
testFilter06 - This method will ensure the correct indicators are returned from
the read-only collection. There should be one returned.
*/
func (s *Suite) testFilter06(indicators []indicator.Indicator) {
	s.Logger.Println("## Test Filter-06: Test Version Filtering Using Specific Version")
	s.Logger.Infoln("++ This test will filter the read-only collection by version using the version 2018-08-08T01:52:01.234Z")

	s.startTest()
	s.setAccept(s.FullMediaType)
	s.enableAuth(s.Settings.Username, s.Settings.Password)

	values := s.Req.URL.Query()
	values.Set("match[version]", "2018-08-08T01:52:01.234Z")
	s.Req.URL.RawQuery = values.Encode()
	s.testFilteringResponse(indicators)
}

/*
testFilter07 - This method will ensure the correct indicators are returned from
the read-only collection. There should be four returned.
*/
func (s *Suite) testFilter07(indicators []indicator.Indicator) {
	s.Logger.Println("## Test Filter-07: Test Version Filtering Using Last,First,Version")
	s.Logger.Infoln("++ This test will filter the read-only collection by version using the last, first, and version")

	s.startTest()
	s.setAccept(s.FullMediaType)
	s.enableAuth(s.Settings.Username, s.Settings.Password)

	values := s.Req.URL.Query()
	values.Set("match[version]", "last,first,2018-08-08T01:53:01.345Z")
	s.Req.URL.RawQuery = values.Encode()
	s.testFilteringResponse(indicators)
}

/*
testFilter08 - This method will ensure the correct indicators
are returned from the read-only collection. There should be four returned.
*/
func (s *Suite) testFilter08(indicators []indicator.Indicator) {
	s.Logger.Println("## Test Filter-08: Test ID Filtering Using One ID")
	s.Logger.Infoln("++ This test will filter the read-only collection by ID using a single STIX ID")

	s.startTest()
	s.setAccept(s.FullMediaType)
	s.enableAuth(s.Settings.Username, s.Settings.Password)

	values := s.Req.URL.Query()
	values.Set("match[id]", "indicator--1efc6673-9d95-46c3-a09c-c29f926da9af")
	s.Req.URL.RawQuery = values.Encode()
	s.testFilteringResponse(indicators)
}

/*
testFilter09 - This method will ensure the correct indicators
are returned from the read-only collection. There should be four returned.
*/
func (s *Suite) testFilter09(indicators []indicator.Indicator) {
	s.Logger.Println("## Test Filter-09: Test ID Filtering Using Two IDs")
	s.Logger.Infoln("++ This test will filter the read-only collection by ID using two STIX IDs")

	s.startTest()
	s.setAccept(s.FullMediaType)
	s.enableAuth(s.Settings.Username, s.Settings.Password)

	values := s.Req.URL.Query()
	values.Set("match[id]", "indicator--1efc6673-9d95-46c3-a09c-c29f926da9af,indicator--213dea46-8750-4b8b-b988-aae8f86a62d6")
	s.Req.URL.RawQuery = values.Encode()
	s.testFilteringResponse(indicators)
}

/*
testFilter10 - This method will ensure the correct indicators are
returned from the read-only collection. There should be two returned.
*/
func (s *Suite) testFilter10(indicators []indicator.Indicator) {
	s.Logger.Println("## Test Filter-10: Test Type Filtering Using Indicator")
	s.Logger.Infoln("++ This test will filter the read-only collection by type using Indicator")

	s.startTest()
	s.setAccept(s.FullMediaType)
	s.enableAuth(s.Settings.Username, s.Settings.Password)

	values := s.Req.URL.Query()
	values.Set("match[type]", "indicator")
	s.Req.URL.RawQuery = values.Encode()
	s.testFilteringResponse(indicators)
}

/*
testFilteringResponse - This method is used by other tests that will test filtering and
ensure that the correct objects are returned.
*/
func (s *Suite) testFilteringResponse(correctIndicators []indicator.Indicator) {
	s.Logger.Infoln("++ Calling Path:", s.Req.URL.Path)
	s.Logger.Infoln("++ Query Params:", s.makePrettyQueryParams())

	// Make HTTP Request
	resp, err := s.Client.Do(s.Req)
	s.handleError(err)
	defer resp.Body.Close()

	// Check HTTP response code first
	s.ProblemsFound += s.checkResponseCode(resp.StatusCode, 200)

	envelopeFromResponse, err := envelope.DecodeRaw(resp.Body)
	if err != nil {
		s.Logger.Println("-- ERROR: Invalid envelope returned", err)
		s.ProblemsFound++
		s.printTestSummary()
		return
	}

	count := 0

	if len(envelopeFromResponse.Objects) > 0 {
		for index, v := range envelopeFromResponse.Objects {

			// Make a first pass to decode just the object type value. Once we have this
			// value we can easily make a second pass and decode the rest of the object.
			stixtype, err := objects.DecodeType(v)
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

				if valid, problems, details := correctIndicators[index].Compare(o); valid != true {
					s.ProblemsFound += problems
					if s.Debug {
						for _, v := range details {
							s.Logger.Debugln(v)
						}
					}
					s.Logger.Println("-- ERROR: Returned indicator", o.ID, "version", o.Modified, "does not match expected")

				} else {
					if s.Debug {
						for _, v := range details {
							s.Logger.Debugln(v)
						}
					}
					s.Logger.Infoln("++ Returned indicator", o.ID, "version", o.Modified, "matches expected")
				}
			}

			count++
		}
		s.Logger.Infoln("++ Number objects returned:", count)

		if s.Debug {
			data, _ := envelopeFromResponse.EncodeToString()
			s.Logger.Debugln("++ Envelope Resource Returned:\n", data)
		}
	}

	s.printTestSummary()
}
