package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strconv"
	"time"
)

type Transaction struct {
	TxID   string
	TxSize int
	TxFee  int
}

type ByFeePerByte []Transaction

func (a ByFeePerByte) Len() int      { return len(a) }
func (a ByFeePerByte) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByFeePerByte) Less(i, j int) bool {
	return float64(a[i].TxFee)/float64(a[i].TxSize) > float64(a[j].TxFee)/float64(a[j].TxSize)
}

func constructBlock(transactions []Transaction) ([]Transaction, int, int, int, float64, int) {
	startTime := time.Now()
	blockSizeLimit := 1000000 // 1MB limit
	currentBlockSize := 0
	totalFee := 0
	var includedTransactions []Transaction

	sort.Sort(ByFeePerByte(transactions))

	for _, transaction := range transactions {
		if currentBlockSize+transaction.TxSize <= blockSizeLimit {
			includedTransactions = append(includedTransactions, transaction)
			currentBlockSize += transaction.TxSize
			totalFee += transaction.TxFee
		}
	}

	endTime := time.Now()
	constructionTime := endTime.Sub(startTime).Seconds()
	memoryUsage := len(includedTransactions) * 16

	return includedTransactions, len(includedTransactions), currentBlockSize, totalFee, constructionTime, memoryUsage
}

func main() {
	transactions := readTransactionsFromFile("transactions.csv")

	block, numTransactions, blockSize, totalFee, constructionTime, memoryUsage := constructBlock(transactions)

	fmt.Println("Constructed Block:")
	for _, transaction := range block {
		fmt.Printf("%s, Size: %d, Fee: %d\n", transaction.TxID, transaction.TxSize, transaction.TxFee)
	}
	fmt.Printf("Number of Transactions: %d\n", numTransactions)
	fmt.Printf("Block Size: %d bytes\n", blockSize)
	fmt.Printf("Total Extracted Value: %d satoshis\n", totalFee)
	fmt.Printf("Construction Time: %.6f seconds\n", constructionTime)
	fmt.Printf("Memory Usage: %d bytes\n", memoryUsage)
}

func readTransactionsFromFile(filename string) []Transaction {
	var transactions []Transaction

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return transactions
	}
	defer file.Close()

	reader := csv.NewReader(file)
	_, err = reader.Read() // Skip header
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return transactions
	}

	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading records:", err)
		return transactions
	}

	for _, record := range records {
		txSize, _ := strconv.Atoi(record[1])
		txFee, _ := strconv.Atoi(record[2])
		transactions = append(transactions, Transaction{
			TxID:   record[0],
			TxSize: txSize,
			TxFee:  txFee,
		})
	}

	return transactions
}
