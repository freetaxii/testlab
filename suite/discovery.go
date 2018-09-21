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
TestDiscoveryService - This method will perform all of the standard tests
against the Discovery endpoint. It will also check to make sure the output
from the GET request is correct and will echo the output to the logs.
*/
func (s *Suite) TestDiscoveryService() {
	s.Logger.Println("## ---------------------------------------------------------")
	s.Logger.Println("## Testing Discovery Service")
	s.Logger.Println("## ---------------------------------------------------------")

	s.setPath(s.Discovery)
	s.EndpointType = "taxii"

	s.basicEndpointTests()
	s.getDiscoveryOutput()
}

func (s *Suite) getDiscoveryOutput() {
	s.Logger.Println("## Test D1: Test successful response from discovery endpoint")
	if s.Verbose {
		s.Logger.Println("++ This test will check to see if a proper discovery resource is returned")
		s.Logger.Println("++ Calling Path:", s.Req.URL.Path)
	}

	media := s.TAXIIMediaType + s.TAXIIVersion
	s.setAccept(media)

	var o resources.Discovery
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
	s.Logger.Println("++ Discovery Resource Returned:\n", string(data))

	s.printSummary()
	s.reset()
}
