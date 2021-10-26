package main

import (
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func genTransactions(n int) []Transaction {
	results := make([]Transaction, n)
	for x := 0; x < n; x++ {
		results[x] = Transaction{
			ID:              strconv.Itoa(x),
			Amount:          rand.Float32() * 1000,
			BankName:        "Street24",
			BankCountryCode: "us", //10 ms latency
		}
	}
	return results
}

func Test_prioritize(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	totalTimeMS := 200
	numExpectedRecords := totalTimeMS / 10 //each transaction makes API call to US with latency 10ms
	trscs := genTransactions(numExpectedRecords + 1)

	result := prioritize(trscs, totalTimeMS)
	resultTimeMS := 0
	for _, t := range result {
		resultTimeMS += ApiLatencies[t.BankCountryCode]
	}
	t.Logf("Total time %d", resultTimeMS)
	if len(result) != numExpectedRecords {
		t.Errorf("Got %d records, want %d", len(result), numExpectedRecords)
	}
}
