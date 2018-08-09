// Copyright 2018 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package suite

import (
	"encoding/json"
	"io/ioutil"

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
	s.Logger.Println()
	s.Logger.Println("== Testing Read-Only Objects Service")
	if s.Verbose {
		s.Logger.Println("++ Calling Path:", s.Req.URL.Path)
	}
	s.basicTests()
	s.getROCollectionObjectsOutput()
}

func (s *Suite) getROCollectionObjectsOutput() {
	s.Logger.Println("== Test O1: Test successful response from read-only objects endpoint")
	if s.Verbose {
		s.Logger.Println("++ This test will check to see if a proper object resource is returned")
	}

	i := GenerateIndicatorData()
	var o objects.Bundle
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

	for index, _ := range i {
		if valid := s.compareCollections(*i[index], o.Objects[index]); valid != true {
			s.Logger.Println("ERROR: Returned indicator does not match expected")
		}
	}

	var data []byte
	data, _ = json.MarshalIndent(o, "", "    ")
	s.Logger.Println("++ Bundle Resource Returned:\n", string(data))

	s.printSummary()
	s.reset()
}

/*
compareIndicators - This method will compare two indicators to make sure they
are the same. Indicator i1 represent the correct data, i2 represents what was
retrieved from a server. So we need to make sure that i2 is the same as i1.
*/
func (s *Suite) compareIndicators(i1, i2 objects.Indicator) bool {

	// Check Type Value
	if i2.ObjectType != i1.ObjectType {
		s.ProblemsFound++
	} else {
		if s.Verbose {
			s.Logger.Println("++ Types Match:", i1.ObjectType, "|", i2.ObjectType)
		}
	}

	// Check Spec Version Value
	if i2.SpecVersion != i1.SpecVersion {
		s.ProblemsFound++
	} else {
		if s.Verbose {
			s.Logger.Println("++ Spec Versions Match:", i1.SpecVersion, "|", i2.SpecVersion)
		}
	}

	// Check ID Value
	if i2.ID != i1.ID {
		s.ProblemsFound++
	} else {
		if s.Verbose {
			s.Logger.Println("++ IDs Match:", i1.ID, "|", i2.ID)
		}
	}

	// Check Created Value
	if i2.Created != i1.Created {
		s.ProblemsFound++
	} else {
		if s.Verbose {
			s.Logger.Println("++ Created Dates Match:", i1.Created, "|", i2.Created)
		}
	}

	// Check Modified Value
	if i2.Modified != i1.Modified {
		s.ProblemsFound++
	} else {
		if s.Verbose {
			s.Logger.Println("++ Modified Dates Match:", i1.Modified, "|", i2.Modified)
		}
	}

	// Check Name Value
	if i2.Name != i1.Name {
		s.ProblemsFound++
	} else {
		if s.Verbose {
			s.Logger.Println("++ Names Match:", i1.Name, "|", i2.Name)
		}
	}

	// Check Description Value
	if i2.Description != i1.Description {
		s.ProblemsFound++
	} else {
		if s.Verbose {
			s.Logger.Println("++ Descriptions Match:", i1.Description, "|", i2.Description)
		}
	}

	// Check Indicator Types Property Length
	if len(i2.IndicatorTypes) != len(i1.IndicatorTypes) {
		s.ProblemsFound++
	} else {
		if s.Verbose {
			s.Logger.Println("++ Indicator Types Length Match:", len(i1.IndicatorTypes), "|", len(i2.IndicatorTypes))
		}
	}

	// Check Indicator Types values
	if i2.IndicatorTypes[0] != i1.IndicatorTypes[0] {
		s.ProblemsFound++
	} else {
		if s.Verbose {
			s.Logger.Println("++ Indicator Types Match:", i1.IndicatorTypes[0], "|", i2.IndicatorTypes[0])
		}
	}

	// Check Pattern Value
	if i2.Pattern != i1.Pattern {
		s.ProblemsFound++
	} else {
		if s.Verbose {
			s.Logger.Println("++ Patterns Match:", i1.Pattern, "|", i2.Pattern)
		}
	}

	// Check ValidFrom Value
	if i2.ValidFrom != i1.ValidFrom {
		s.ProblemsFound++
	} else {
		if s.Verbose {
			s.Logger.Println("++ ValidFrom Values Match:", i1.ValidFrom, "|", i2.ValidFrom)
		}
	}

	// Check ValidUntil Value
	if i2.ValidUntil != i1.ValidUntil {
		s.ProblemsFound++
	} else {
		if s.Verbose {
			s.Logger.Println("++ ValidUntil Values Match:", i1.ValidUntil, "|", i2.ValidUntil)
		}
	}

	// Check Kill Chain Phases Property Length
	if len(i2.KillChainPhases) != len(i1.KillChainPhases) {
		s.ProblemsFound++
	} else {
		if s.Verbose {
			s.Logger.Println("++ Kill Chain Phases Length Match:", len(i1.KillChainPhases), "|", len(i2.KillChainPhases))
		}
	}

	// Check Kill Chain Phases values
	if i2.KillChainPhases[0].KillChainName != i1.KillChainPhases[0].KillChainName {
		s.ProblemsFound++
	} else {
		if s.Verbose {
			s.Logger.Println("++ Kill Chain Phases Match:", i1.KillChainPhases[0].KillChainName, "|", i2.KillChainPhases[0].KillChainName)
		}
	}

	// Check Kill Chain Phases values
	if i2.KillChainPhases[0].PhaseName != i1.KillChainPhases[0].PhaseName {
		s.ProblemsFound++
	} else {
		if s.Verbose {
			s.Logger.Println("++ Kill Chain Phases Match:", i1.KillChainPhases[0].PhaseName, "|", i2.KillChainPhases[0].PhaseName)
		}
	}

	if s.ProblemsFound > 0 {
		s.Logger.Printf("ERROR: Returned indicator does not match expected value")
		s.Logger.Printf("ERROR: Expected %s", i1)
		s.Logger.Printf("ERROR: Got %s", i2)
		return false
	}

	return true
}
