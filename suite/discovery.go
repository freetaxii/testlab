// Copyright 2018 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package suite

import (
	"encoding/json"
	"io/ioutil"

	"github.com/freetaxii/libstix2/resources/discovery"
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

	s.setPath(s.Settings.Discovery)

	s.basicEndpointTests()
	s.getDiscoveryOutput()
}

func (s *Suite) getDiscoveryOutput() {
	s.Logger.Println("## Test D1: Test Discovery Endpoint")
	s.Logger.Infoln("++ This test will check to see if a proper discovery resource is returned")
	s.Logger.Infoln("++ Calling Path:", s.Req.URL.Path)

	s.startTest()
	s.setAccept(s.FullMediaType)
	s.enableAuth(s.Settings.Username, s.Settings.Password)

	var o discovery.Discovery
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

	s.printTestSummary()
}
