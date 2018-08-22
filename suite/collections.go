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
TestCollectionsService - This method will perform all of the standard tests
against the Collections endpoint. It will also check to make sure the output
from the GET request is correct and will echo the output to the logs.
*/
func (s *Suite) TestCollectionsService() {
	s.Logger.Println()
	s.Logger.Println("== Testing Collections Service")

	path := s.APIRoot + "collections/"
	s.setPath(path)
	s.EndpointType = "taxii"

	s.basicEndpointTests()
	s.getCollectionsOutput()
}

func (s *Suite) getCollectionsOutput() {
	s.Logger.Println("== Test C1: Test successful response from collections endpoint")
	if s.Verbose {
		s.Logger.Println("++ This test will check to see if a proper collections resource is returned")
		s.Logger.Println("++ Calling Path:", s.Req.URL.Path)
	}

	media := s.TAXIIMediaType + s.TAXIIVersion
	s.setAccept(media)

	var o resources.Collections
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
	s.Logger.Println("++ Collections Resource Returned:\n", string(data))

	s.printSummary()
	s.reset()
}
