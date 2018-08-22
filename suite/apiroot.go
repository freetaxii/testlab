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
TestAPIRootService - This method will perform all of the standard tests
against the API Root endpoint. It will also check to make sure the output
from the GET request is correct and will echo the output to the logs.
*/
func (s *Suite) TestAPIRootService() {
	s.setPath(s.APIRoot)
	s.EndpointType = "taxii"
	s.Logger.Println()
	s.Logger.Println("== Testing API Root Service")
	if s.Verbose {
		s.Logger.Println("++ Calling Path:", s.Req.URL.Path)
	}

	s.basicTests()
	s.getAPIRootOutput()
}

func (s *Suite) getAPIRootOutput() {
	s.Logger.Println("== Test A1: Test successful response from api root endpoint")
	if s.Verbose {
		s.Logger.Println("++ This test will check to see if a proper API root resource is returned")
	}

	media := s.TAXIIMediaType + s.TAXIIVersion
	s.setAccept(media)

	var o resources.APIRoot
	s.Req.SetBasicAuth(s.Username, s.Password)
	resp, err := s.Client.Do(s.Req)
	s.handleError(err)
	defer resp.Body.Close()
	s.ProblemsFound += s.checkResponseCode(resp.StatusCode, 200)

	body, err := ioutil.ReadAll(resp.Body)
	s.handleError(err)

	jerr := json.Unmarshal(body, &o)
	s.handleError(jerr)

	var data []byte
	data, _ = json.MarshalIndent(o, "", "    ")
	s.Logger.Println("++ API Root Resource Returned:\n", string(data))

	s.printSummary()
	s.reset()
}
