// Copyright 2018 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package testing

import (
	"encoding/json"
	"io/ioutil"

	"github.com/freetaxii/libstix2/resources"
	"github.com/gologme/log"
)

func (s *Suite) TestDiscoveryService() {
	s.setPath(s.Discovery)
	log.Println()
	log.Println("== Testing Discovery Service")
	s.BasicTests()
	s.DiscoveryOutput()
}

func (s *Suite) DiscoveryOutput() {
	log.Println("== Test D1: Test successful response")
	if s.Verbose {
		log.Println("++ This test will send a correct authentication parameter and will check to see if a proper discovery resource is returned")
	}

	var o resources.Discovery
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
	log.Println("\n", string(data))

	s.printSummary()
	s.reset()
}
