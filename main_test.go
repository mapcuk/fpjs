package main

import (
	"testing"
)

func Test_prioritize(t *testing.T) {
	trscs := []Transaction{{
		ID:              "001",
		Amount:          190,
		BankName:        "Street24",
		BankCountryCode: "au", //250 ms latency
	}, {
		ID:              "002",
		Amount:          250,
		BankName:        "Street24",
		BankCountryCode: "ae", //80 ms latency
	}, {
		ID:              "003",
		Amount:          170,
		BankName:        "Street24",
		BankCountryCode: "vn", //129 ms latency
	}, {
		ID:              "004",
		Amount:          150,
		BankName:        "Street24",
		BankCountryCode: "fj", //360 ms latency
	}}
	totalTimeMS := 500
	numExpectedRecords := 3

	result := prioritize(trscs, totalTimeMS)
	resultTimeMS := 0
	for _, t := range result {
		resultTimeMS += ApiLatencies[t.BankCountryCode]
	}
	if resultTimeMS > totalTimeMS {
		t.Errorf("Got transactions which takes %d ms (more than %d)", resultTimeMS, totalTimeMS)
	}
	if len(result) != numExpectedRecords {
		t.Errorf("Got %d records, want %d", len(result), numExpectedRecords)
	}
}
