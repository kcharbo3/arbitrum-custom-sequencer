package main

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github/kcharbo3/custom-sequencer-script/batcher"
	"io/ioutil"
	"os"
	"strconv"
)
// InfuraEndpoint - endpoint to use. Should be for rinkeby since the custom sequencerinbox contract is deployed there
const InfuraEndpoint string = "https://rinkeby.infura.io/v3/acbb136e15fd40f38ea36b3e11469079"

// TransactionJson - path to transaction json to replicate
const TransactionJson string = "../../json/updatedSingleTx.json"

// OwnerAddressString - this address owns the sequencerinbox contract. Can be any address that is registered as a sequencer on the inbox smart contract
const OwnerAddressString string = "0x43f8D77300Cf8283D0FB1919CDCF9a0402e9c02A"

// MultipleSeqMsgs - determines whether each transaction will be separated into their own sequencer msg
const MultipleSeqMsgs bool = false

// MultipleSections - determines whether each sequencer msg will be separated into their own section
const MultipleSections bool = false

// NumTxsToSend - number of transactions to send in one call to the sequencer inbox
const NumTxsToSend int = 25

func main() {
	tx := readTransactionJson()

	client, err := ethclient.Dial(InfuraEndpoint)
	if err != nil {
		fmt.Println("Error creating ethclient")
		return
	}

	sequencer, err := batcher.NewSequencerBatcher(OwnerAddressString, client)
	if err != nil {
		fmt.Println("Error creating new seq batcher")
		return
	}

	transactions := make([]*types.Transaction, 0)
	for i := 0; i < NumTxsToSend; i++ {
		/*
		randomTx := message.NewRandomTransaction()
		randomEthTx := randomTx.AsEthTx()
		transactions = append(transactions, randomEthTx) */
		transactions = append(transactions, tx)
	}

	fmt.Println("\nBatching " + strconv.Itoa(len(transactions)) + " transactions...")
	sequencerBatchItems := sequencer.CreateMultipleBatches(transactions, MultipleSeqMsgs, MultipleSections)
	fmt.Println("Number of batches created: " + strconv.Itoa(len(sequencerBatchItems)))

	fmt.Println("\nPublishing Batches...")
	didPublish, err := sequencer.PublishBatchToInbox(sequencerBatchItems)
	if err != nil {
		fmt.Println("Error publishing batch to inbox!")
	}
	fmt.Println("Batch published: " + strconv.FormatBool(didPublish))
}

func readTransactionJson() *types.Transaction {
	jsonFile, err := os.Open(TransactionJson)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened test.json")

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var tx types.Transaction

	err = json.Unmarshal(byteValue, &tx)
	if err != nil {
		fmt.Println("Error unmarshaling json: " + err.Error())
		return nil
	}

	defer jsonFile.Close()

	return &tx
}
