package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/btcsuite/btcd/rpcclient"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get RPC configuration from environment variables
	rpcUser := os.Getenv("RPC_USER")
	rpcPass := os.Getenv("RPC_PASS")
	rpcHost := os.Getenv("RPC_HOST")

	// Connect to Bitcoin Core RPC server
	client, err := rpcclient.New(&rpcclient.ConnConfig{
		HTTPPostMode: true,
		DisableTLS:   true,
		Host:         rpcHost,
		User:         rpcUser,
		Pass:         rpcPass,
	}, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Shutdown()

	// Get the current block count
	blockCount, err := client.GetBlockCount()
	if err != nil {
		log.Fatal(err)
	}

	// Iterate through blocks and find the longest UTXO chain
	var longestChainLength int
	var longestChainBlockHash string
	for i := int64(0); i < blockCount; i++ {
		blockHash, err := client.GetBlockHash(i)
		if err != nil {
			log.Fatal(err)
		}

		blockData, err := client.GetBlockVerboseTx(blockHash)
		if err != nil {
			log.Fatal(err)
		}

		// Convert block data to JSON-encoded byte slice
		blockJSON, err := json.Marshal(blockData)
		if err != nil {
			log.Fatal(err)
		}

		var block map[string]interface{}
		err = json.Unmarshal(blockJSON, &block)
		if err != nil {
			log.Fatal(err)
		}

		utxoChainLength := countUTXOChainLength(block)
		if utxoChainLength > longestChainLength {
			longestChainLength = utxoChainLength
			longestChainBlockHash = blockHash.String()
		}
	}

	fmt.Printf("Longest UTXO chain length: %d\n", longestChainLength)
	fmt.Printf("Block hash of the block with the longest UTXO chain: %s\n", longestChainBlockHash)
}

func countUTXOChainLength(block map[string]interface{}) int {
	chainLength := 0
	transactions, ok := block["tx"].([]interface{})
	if !ok {
		return chainLength
	}
	for _, tx := range transactions {
		txMap, ok := tx.(map[string]interface{})
		if !ok {
			continue
		}
		vin, ok := txMap["vin"].([]interface{})
		if !ok {
			continue
		}
		vout, ok := txMap["vout"].([]interface{})
		if !ok {
			continue
		}
		if len(vin) == 1 && len(vout) == 2 {
			chainLength++
		}
	}
	return chainLength
}
