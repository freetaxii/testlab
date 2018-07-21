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

func (s *Suite) TestAPIRootService() {
	s.setPath(s.APIRoot)
	s.Logger.Println()
	s.Logger.Println("== Testing API Root Service")
	s.basicTests()
	s.testAPIRootOutput()
}

func (s *Suite) testAPIRootOutput() {
	s.Logger.Println("== Test A1: Test successful response")
	if s.Verbose {
		s.Logger.Println("++ This test will check to see if a proper API root resource is returned")
	}

	var o resources.APIRoot
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

	var data []byte
	data, _ = json.MarshalIndent(o, "", "    ")
	s.Logger.Println("\n", string(data))

	s.printSummary()
	s.reset()
}
