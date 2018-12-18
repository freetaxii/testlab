// Copyright 2015-2018 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source tree.

package suite

import (
	"crypto/tls"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/gologme/log"
)

type Suite struct {
	Logger         *log.Logger
	Req            *http.Request
	Client         *http.Client
	Verbose        bool
	Debug          bool
	ProblemsFound  int
	TAXIIMediaType string
	TAXIIVersion   string
	FullMediaType  string
	Settings       struct {
		Username  string
		Password  string
		URL       string
		Proxy     string
		Discovery string
		APIRoot   string
	}
	CollectionIDs struct {
		ReadOnly  string
		WriteOnly string
		ReadWrite string
	}
}

/*
New - This function will create a new test suite object and assign a logger
*/
func New(logger *log.Logger) *Suite {
	var s Suite

	// ------------------------------------------------------------
	// Setup Logging
	// ------------------------------------------------------------
	if logger == nil {
		s.Logger = log.New(os.Stderr, "", log.LstdFlags)
	} else {
		s.Logger = logger
	}

	return &s
}

/*
Setup - This method will setup the test suite
*/
func (s *Suite) Setup() {

	// ------------------------------------------------------------
	// Setup Logging Levels
	// ------------------------------------------------------------
	if s.Verbose {
		s.Logger.EnableLevel("info")
	}

	if s.Debug {
		s.Logger.EnableLevel("debug")
	}

	// ------------------------------------------------------------
	// Setup Media Types
	// ------------------------------------------------------------
	s.TAXIIMediaType = "application/taxii+json"
	s.TAXIIVersion = "version=2.1"
	s.FullMediaType = s.TAXIIMediaType + ";" + s.TAXIIVersion

	// ------------------------------------------------------------
	// Verify Endpoints
	// ------------------------------------------------------------
	if !strings.HasPrefix(s.Settings.Discovery, "/") {
		s.Settings.Discovery = "/" + s.Settings.Discovery
	}
	if !strings.HasSuffix(s.Settings.Discovery, "/") {
		s.Settings.Discovery = s.Settings.Discovery + "/"
	}

	if !strings.HasPrefix(s.Settings.APIRoot, "/") {
		s.Settings.APIRoot = "/" + s.Settings.APIRoot
	}
	if !strings.HasSuffix(s.Settings.APIRoot, "/") {
		s.Settings.APIRoot = s.Settings.APIRoot + "/"
	}

	// ------------------------------------------------------------
	// Setup HTTP Client and Proxy if defined
	// ------------------------------------------------------------
	var err error
	var netTransport *http.Transport
	if s.Settings.Proxy != "" {
		proxyURL, err := url.Parse(s.Settings.Proxy)
		if err != nil {
			s.Logger.Fatalln(err)
		}

		netTransport = &http.Transport{
			Dial: (&net.Dialer{
				Timeout: 5 * time.Second,
			}).Dial,
			TLSHandshakeTimeout: 5 * time.Second,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
			Proxy: http.ProxyURL(proxyURL),
		}
	} else {
		netTransport = &http.Transport{
			Dial: (&net.Dialer{
				Timeout: 5 * time.Second,
			}).Dial,
			TLSHandshakeTimeout: 5 * time.Second,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}
	}

	s.Client = &http.Client{
		Timeout:   time.Second * 10,
		Transport: netTransport,
	}

	s.Req, err = http.NewRequest(http.MethodGet, s.Settings.URL, nil)
	s.handleError(err)
}
