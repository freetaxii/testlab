// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package suite

import (
	"github.com/freetaxii/libstix2/objects"
)

func GenerateIndicatorData() []objects.Indicator {
	var indicators []objects.Indicator

	i1 := objects.NewIndicator()
	i1.SetID("indicator--1efc6673-9d95-46c3-a09c-c29f926da9af")
	i1.SetCreated("2018-08-08T01:51:01.123Z")
	i1.SetModified("2018-08-08T01:51:01.123Z")
	i1.SetCreatedByRef("identity--abd090f7-5ada-4506-b6d0-5feae5ff90bc")
	i1.SetLang("en-us")
	i1.SetConfidence(99)
	i1.SetName("TestLab Indicator 1")
	i1.SetDescription("This is indicator 1 for Read-Only TestLab Collection")
	i1.AddType("compromised")
	i1.SetValidFrom("2018-08-08T01:51:01.123Z")
	i1.SetValidUntil("2018-09-09T01:51:01.123Z")
	pattern1 := "[ ipv4-addr:value = '192.168.100.100' ]"
	i1.SetPattern(pattern1)
	i1.CreateKillChainPhase("lockheed-martin-cyber-kill-chain", "delivery")
	indicators = append(indicators, *i1)

	i11 := i1
	i11.SetModified("2018-08-08T01:52:01.234Z")
	i11.AddLabel("a")
	indicators = append(indicators, *i11)

	i12 := objects.NewIndicator()
	i12 = i11
	i12.SetModified("2018-08-08T01:53:01.345Z")
	i12.AddLabel("b")
	i12.AddLabel("c")
	indicators = append(indicators, *i12)

	i13 := i12
	i13.SetModified("2018-08-08T01:54:01.456Z")
	indicators = append(indicators, *i13)

	i14 := i1
	i14.SetModified("2018-08-08T01:55:01.567Z")
	i14.Labels = nil
	i14.AddLabel("a")
	i14.AddLabel("b")
	i14.AddLabel("d")
	indicators = append(indicators, *i14)

	i2 := objects.NewIndicator()
	i2.SetID("indicator--213dea46-8750-4b8b-b988-aae8f86a62d6")
	i2.SetCreated("2018-08-08T02:51:02.123Z")
	i2.SetModified("2018-08-08T02:51:02.123Z")
	i2.SetCreatedByRef("identity--abd090f7-5ada-4506-b6d0-5feae5ff90bc")
	i2.SetLang("en-us")
	i2.SetConfidence(99)
	i2.SetName("TestLab Indicator 2")
	i2.SetDescription("This is indicator 2 for Read-Only TestLab Collection")
	i2.AddType("anonymization")
	i2.AddType("compromised")
	i2.SetValidFrom("2018-08-08T02:51:02.123Z")
	i2.SetValidUntil("2018-09-09T02:51:02.123Z")
	pattern2 := "[ ipv4-addr:value = '192.168.200.200' ]"
	i2.SetPattern(pattern2)
	i2.CreateKillChainPhase("lockheed-martin-cyber-kill-chain", "delivery")
	extref, _ := i2.NewExternalReference()
	extref.SetSourceName("TestLab Indicators")
	extref.SetDescription("This is from the TestLab")
	extref.SetExternalID("2")
	extref.SetURL("https://github.com/freetaxii/testlab")
	indicators = append(indicators, *i2)

	return indicators
}

func GenerateAttackPatternData() []objects.AttackPattern {
	var ap []objects.AttackPattern

	a1 := objects.NewAttackPattern()
	a1.SetID("attack-pattern--9a624a80-ac52-49e2-b4ef-6b5e5f26a50d")
	a1.SetCreated("2018-08-08T03:51:03.123Z")
	a1.SetModified("2018-08-08T03:51:03.123Z")
	a1.SetName("TestLab Attack Pattern 1")
	a1.SetDescription("This is attack pattern 1 for Read-Only TestLab Collection")
	a1.CreateKillChainPhase("lockheed-martin-cyber-kill-chain", "delivery")
	er, _ := a1.NewExternalReference()
	er.SetSourceName("capec")
	er.SetExternalID("CAPEC-163")
	ap = append(ap, *a1)

	return ap
}

func GenerateThreatActorData() []objects.ThreatActor {
	var ta []objects.ThreatActor

	t1 := objects.NewThreatActor()
	t1.SetID("threat-actor--a6036137-f757-482e-bf63-fcb5e25efdd8")
	t1.SetCreated("2018-08-08T04:51:04.123Z")
	t1.SetModified("2018-08-08T04:51:04.123Z")
	t1.SetName("TestLab Threat Actor 1")
	t1.SetDescription("This is threat actor 1 for Read-Only TestLab Collection")
	t1.AddType("activist")
	ta = append(ta, *t1)

	return ta
}

func GenerateCampaignData() []objects.Campaign {
	var c []objects.Campaign

	c1 := objects.NewCampaign()
	c1.SetID("campaign--bba8b6c6-fa62-4767-8303-58390db33a19")
	c1.SetCreated("2018-08-08T05:51:05.123Z")
	c1.SetModified("2018-08-08T05:51:05.123Z")
	c1.SetName("TestLab Campaign 1")
	c1.SetDescription("This is campaign 1 for Read-Only TestLab Collection")
	c = append(c, *c1)

	return c
}
