// Copyright 2018 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package main

import (
	"fmt"
	"os"

	"github.com/freetaxii/testlab/suite"
	"github.com/gologme/log"
	"github.com/pborman/getopt"
)

// These global variables hold build information. The Build variable will be
// populated by the Makefile and uses the Git Head hash as its identifier.
// These variables are used in the console output for --version and --help.
var (
	Version = "0.5.1"
	Build   string
)

// These global variables are for dealing with command line options
var (
	sOptURL       = getopt.StringLong("url", 'u', "https://127.0.0.1:8000/", "TAXII Server Address", "string")
	sOptProxy     = getopt.StringLong("proxy", 'x', "", "Proxy Server Address", "string")
	sOptDiscovery = getopt.StringLong("discovery", 'd', "taxii2", "Name of Discovery Service", "string")
	sOptAPIRoot   = getopt.StringLong("apiroot", 'a', "api1", "Name of API Root", "string")
	sOptReadOnly  = getopt.StringLong("readonly", 'r', "22f763c1-e478-4765-8635-e4c32db665ea", "The read-only collection ID", "string")
	sOptWriteOnly = getopt.StringLong("writeonly", 'w', "4f7327e2-f5b4-4269-b6e0-3564d174ce69", "The write-only collection ID", "string")
	sOptReadWrite = getopt.StringLong("readwrite", 'z', "8c49f14d-8ea3-4f03-ab28-19dbca973dde", "The read-write collection ID", "string")
	sOptUsername  = getopt.StringLong("username", 'n', "", "Username", "string")
	sOptPassword  = getopt.StringLong("password", 'p', "", "Password", "string")
	bOptVerbose   = getopt.BoolLong("verbose", 0, "Enable verbose output")
	bOptDebug     = getopt.BoolLong("debug", 0, "Enable debug output")
	bOptHelp      = getopt.BoolLong("help", 0, "Help")
	bOptVer       = getopt.BoolLong("version", 0, "Version")
)

func main() {
	// --------------------------------------------------
	// Setup logger
	// --------------------------------------------------
	logger := log.New(os.Stderr, "", log.LstdFlags)

	s := suite.New(logger)
	processCommandLineFlags(s)

	logger.Println("## ---------------------------------------------------------")
	logger.Println("## Starting FreeTAXII Testing Suite...")
	logger.Println("## ---------------------------------------------------------\n")

	s.Setup()
	s.TestDiscoveryService()
	s.TestAPIRootService()
	s.TestCollectionsService()
}

// --------------------------------------------------
// Private functions
// --------------------------------------------------

/*
processCommandLineFlags - This function will process the command line flags
and will print the version or help information as needed.
*/
func processCommandLineFlags(s *suite.Suite) {
	getopt.HelpColumn = 35
	getopt.DisplayWidth = 120
	getopt.SetParameters("")
	getopt.Parse()

	// Lets check to see if the version command line flag was given. If it is
	// lets print out the version information and exit.
	if *bOptVer {
		printOutputHeader()
		os.Exit(0)
	}

	// Lets check to see if the help command line flag was given. If it is lets
	// print out the help information and exit.
	if *bOptHelp {
		printOutputHeader()
		getopt.Usage()
		os.Exit(0)
	}

	// ------------------------------------------------------------
	// Map command line parameters to struct values
	// ------------------------------------------------------------
	s.Verbose = *bOptVerbose
	s.Debug = *bOptDebug

	s.Settings.URL = *sOptURL
	s.Settings.Proxy = *sOptProxy
	s.Settings.Discovery = *sOptDiscovery
	s.Settings.APIRoot = *sOptAPIRoot
	s.Settings.Username = *sOptUsername
	s.Settings.Password = *sOptPassword

	s.CollectionIDs.ReadOnly = *sOptReadOnly
	s.CollectionIDs.WriteOnly = *sOptWriteOnly
	s.CollectionIDs.ReadWrite = *sOptReadWrite
}

/*
printOutputHeader - This function will print a header for all console output
*/
func printOutputHeader() {
	fmt.Println("")
	fmt.Println("FreeTAXII TestLab - Basic Connectivity Tests")
	fmt.Println("Copyright: Bret Jordan")
	fmt.Println("Version:", Version)
	if Build != "" {
		fmt.Println("Build:", Build)
	}
	fmt.Println("")
}
