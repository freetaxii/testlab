// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/freetaxii/libstix2/datastore/sqlite3"
	"github.com/freetaxii/libstix2/objects/bundle"
	"github.com/freetaxii/libstix2/resources/collections"
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
	defaultDatabaseFilename = "freetaxii.db"
	sOptDatabaseFilename    = getopt.StringLong("filename", 'f', defaultDatabaseFilename, "Database Filename", "string")
	bOptIndicatorsOnly      = getopt.BoolLong("indicator", 'i', "Only print indicators")
	bOptDatabase            = getopt.BoolLong("database", 0, "Add to database")
	bOptHelp                = getopt.BoolLong("help", 0, "Help")
	bOptVer                 = getopt.BoolLong("version", 0, "Version")
)

func main() {
	processCommandLineFlags()
	var data []byte
	var ds *sqlite3.Store
	var err error
	database := *bOptDatabase
	indicatorsOnly := *bOptIndicatorsOnly

	// --------------------------------------------------
	// Setup logger
	// --------------------------------------------------
	logger := log.New(os.Stderr, "", log.LstdFlags)
	logger.EnableLevel("info")

	colCache := make(map[string]collections.Collection)

	if database {
		ds = sqlite3.New(logger, *sOptDatabaseFilename, colCache)
		defer ds.Close()
	}

	// ----------------------------------------------------------------------
	//
	// Create TAXII Collections
	//
	// ----------------------------------------------------------------------
	fmt.Println("\n\nStart TAXII Collections Output")

	col := collections.New()

	// Create ReadOnly Collection
	c1 := suite.GenerateROCollection()
	if database {
		err = ds.AddTAXIIObject(c1)
		handleError(err)
	}
	col.AddCollection(c1)

	// Create WriteOnly Collection
	c2 := suite.GenerateWOCollection()
	if database {
		err = ds.AddTAXIIObject(c2)
		handleError(err)
	}
	col.AddCollection(c2)

	// Create ReadOnly Collection
	c3 := suite.GenerateRWCollection()
	if database {
		err = ds.AddTAXIIObject(c3)
		handleError(err)
	}
	col.AddCollection(c3)

	data, _ = json.MarshalIndent(col, "", "    ")
	fmt.Println(string(data))

	// ----------------------------------------------------------------------
	//
	// Create STIX Objects
	//
	// ----------------------------------------------------------------------

	// Need 2 indicators
	// 		one with multiple versions with different labels on each
	// Need 1 attack pattern
	// Need 1 threat actor
	// Need 1 campaign
	// Need 2 relationship objects
	// Need 1 sighting
	// Total 8 objects plus versions
	//
	// Tests Objects Endpoint
	// Get all objects correctly
	// Get all objects by type
	// Get all objects by "latest"
	// Get all objects by "first"
	// Get all objects by type and by latest
	// Get all objects by type and by first
	// Get a specific version of an object (test positive and negative case, where you provide a version that does not exist)
	// Get type that is not in database
	// Get an id that is not in the database
	//
	// Objects by ID Endpoint
	// Get object
	// Get object that is not there (error)
	// Get object latest
	// Get object first
	// Get object specific version (positive and negative case)
	//
	// Get manifests

	fmt.Println("\n\nStart STIX Object Output")
	b := bundle.New()
	b.SetID("bundle--e5214f9b-ae28-4692-9394-2fd2ed85d78a")

	counter := make(map[string]int)
	iData := suite.GenerateIndicatorData()
	for _, v := range iData {
		b.AddObject(v)
		counter[v.ID]++
		if database {
			err = ds.AddObject(&v)
			handleError(err)
			// First check to see if the STIX ID is already in the collection
			// we do this by checking this local map that contains a counter
			// since the collection should be new before running this command
			if counter[v.ID] == 1 {
				err = ds.AddToCollection(c1.ID, v.ID)
				handleError(err)
			}
		}
	}

	if indicatorsOnly == false {
		apData := suite.GenerateAttackPatternData()
		for _, v := range apData {
			b.AddObject(v)
			if database {
				ds.AddObject(v)
				err = ds.AddToCollection(c1.ID, v.ID)
				handleError(err)
			}
		}

		taData := suite.GenerateThreatActorData()
		for _, v := range taData {
			b.AddObject(v)
			if database {
				ds.AddObject(v)
				err = ds.AddToCollection(c1.ID, v.ID)
				handleError(err)
			}
		}

		cData := suite.GenerateCampaignData()
		for _, v := range cData {
			b.AddObject(v)
			if database {
				ds.AddObject(v)
				err = ds.AddToCollection(c1.ID, v.ID)
				handleError(err)
			}
		}
	}

	data, _ = json.MarshalIndent(b, "", "    ")
	fmt.Println(string(data))

}

func handleError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

// --------------------------------------------------
// Private functions
// --------------------------------------------------

// processCommandLineFlags - This function will process the command line flags
// and will print the version or help information as needed.
func processCommandLineFlags() {
	getopt.HelpColumn = 35
	getopt.DisplayWidth = 120
	getopt.SetParameters("")
	getopt.Parse()

	// Lets check to see if the version command line flag was given. If it is
	// lets print out the version infomration and exit.
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
}

// printOutputHeader - This function will print a header for all console output
func printOutputHeader() {
	fmt.Println("")
	fmt.Println("FreeTAXII TestLab - Test Data")
	fmt.Println("Copyright: Bret Jordan")
	fmt.Println("Version:", Version)
	if Build != "" {
		fmt.Println("Build:", Build)
	}
	fmt.Println("")
}
