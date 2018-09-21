// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package suite

import "github.com/freetaxii/libstix2/resources/collections"

func GenerateROCollection() *collections.Collection {
	c1 := collections.NewCollection()
	c1.SetID("22f763c1-e478-4765-8635-e4c32db665ea")
	c1.SetTitle("Read-Only TestLab Collection")
	c1.SetDescription("This is a Read-Only collection for use with the FreeTAXII TestLab tool")
	c1.SetCanRead()
	c1.AddMediaType("application/stix+json;version=2.1")
	return c1
}

func GenerateWOCollection() *collections.Collection {
	c2 := collections.NewCollection()
	c2.SetID("4f7327e2-f5b4-4269-b6e0-3564d174ce69")
	c2.SetTitle("Write-Only TestLab Collection")
	c2.SetDescription("This is a Write-Only collection for use with the FreeTAXII TestLab tool")
	c2.SetCanWrite()
	c2.AddMediaType("application/stix+json;version=2.1")
	return c2
}

func GenerateRWCollection() *collections.Collection {
	c3 := collections.NewCollection()
	c3.SetID("8c49f14d-8ea3-4f03-ab28-19dbca973dde")
	c3.SetTitle("Read-Write TestLab Collection")
	c3.SetDescription("This is a Read-Write collection for use with the FreeTAXII TestLab tool")
	c3.SetCanRead()
	c3.SetCanWrite()
	c3.AddMediaType("application/stix+json;version=2.1")
	return c3
}
