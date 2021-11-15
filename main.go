package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

var ApiLatencies = map[string]int{
	"ae": 80,
	"ar": 87,
	"au": 250,
	"be": 46,
	"bh": 82,
	"br": 37,
	"ca": 12,
	"ch": 55,
	"cl": 83,
	"cn": 115,
	"cy": 77,
	"de": 48,
	"es": 56,
	"fi": 50,
	"fj": 360,
	"fr": 53,
	"gi": 61,
	"gr": 66,
	"hk": 130,
	"id": 227,
	"ie": 42,
	"il": 79,
	"it": 62,
	"jp": 122,
	"ky": 30,
	"ma": 88,
	"mx": 14,
	"ng": 102,
	"nl": 47,
	"no": 46,
	"nz": 350,
	"pl": 49,
	"ro": 51,
	"ru": 55,
	"sa": 78,
	"se": 47,
	"sg": 130,
	"th": 133,
	"tr": 99,
	"ua": 52,
	"uk": 45,
	"us": 10,
	"vn": 129,
	"za": 105}

type Transaction struct {
	// a UUID of transaction
	ID string
	// in USD, typically a value betwen 0.01 and 1000 USD.
	Amount float32
	// bank name, e.g. "Bank of Scotland"
	BankName string
	// a 2-letter country code of where the bank is located
	BankCountryCode string
}

type TransactionStatus struct {
	// a UUID of transaction
	ID string
	// transaction is fraud
	isFraud bool
}

// method is called every second and receives an array of transactions
func processTransactions(transactions []Transaction) []TransactionStatus {
	results := make([]TransactionStatus, len(transactions))
	var totalAmount float32
	for _, transaction := range transactions {
		status := TransactionStatus{ID: transaction.ID}
		status.isFraud = processTransaction(transaction)
		results = append(results, status)
		totalAmount += transaction.Amount
	}
	log.Printf("Total amount: %.2f", totalAmount)
	return results
}

// function that will internally call the bank API and verify the transaction using this API
func processTransaction(transaction Transaction) bool {
	return false
}

// Common Knapsack problem 0-1, dynamic programming approach
func prioritize(transactions []Transaction, totalTimeMs int) []Transaction {
	count := len(transactions)
	table := make([][]float32, count+1)
	for index := range table {
		table[index] = make([]float32, totalTimeMs+1)
	}

	latencies := make([]int, count+1)
	for i := 0; i < count; i++ {
		transaction := transactions[i]
		latency, exists := ApiLatencies[transaction.BankCountryCode]
		if !exists {
			log.Printf("unknown latency for country %s", transaction.BankCountryCode)
			latency = totalTimeMs // let's consider this transaction takes max time
		}
		latencies[i] = latency
		for j := 0; j <= totalTimeMs; j++ {
			if latency > j {
				table[i+1][j] = table[i][j]
			} else {
				maxAmount := table[i][j-latency] + transaction.Amount

				if maxAmount > table[i][j] {
					table[i+1][j] = maxAmount
				} else {
					table[i+1][j] = table[i][j]
				}
			}
		}
	}

	var prioritized []Transaction
	j := totalTimeMs
	for i := count; i > 0; i-- {
		if table[i][j] > table[i-1][j] {
			prioritized = append(prioritized, transactions[i-1])
			j = j - latencies[i-1]
		}
	}
	return prioritized
}

func recordToTransaction(record []string) Transaction {
	newTransaction := Transaction{ID: record[0], BankCountryCode: record[2]}

	amount, err := strconv.ParseFloat(record[1], 32)
	if err != nil {
		log.Fatalf("bad amount in record: %s", strings.Join(record, ","))
	}
	newTransaction.Amount = float32(amount)
	return newTransaction
}

func main() {
	transactionsFile, err := os.Open("transactions.csv")
	if err != nil {
		log.Fatalf("Error opening file %v", err)
	}
	var transactions []Transaction
	reader := csv.NewReader(transactionsFile)
	_, _ = reader.Read() // skip header
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error reading CSV file: %v", err)
		}

		if len(record) != 3 {
			log.Fatalf("bad format of record %s", strings.Join(record, ", "))
		}
		transactions = append(transactions, recordToTransaction(record))
	}
	processWithPriority(transactions, 1000)
	processWithPriority(transactions, 50)
	processWithPriority(transactions, 60)
	processWithPriority(transactions, 90)
}

func processWithPriority(transactions []Transaction, totalTimeMs int) {
	var transactionsCopy = make([]Transaction, len(transactions))
	// copy slices to avoid side effects
	copy(transactionsCopy, transactions)
	log.Printf("Process transactions limited by total time: %d ms", totalTimeMs)
	sortedTransactions := prioritize(transactionsCopy, totalTimeMs)
	_ = processTransactions(sortedTransactions)
}
