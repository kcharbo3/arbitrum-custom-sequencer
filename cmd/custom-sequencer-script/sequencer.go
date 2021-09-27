package main

import (
	"fmt"
	"strconv"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"

	"github/kcharbo3/custom-sequencer-script/batcher"
)

const InfuraEndpoint string = "https://rinkeby.infura.io/v3/acbb136e15fd40f38ea36b3e11469079"
const OwnerAddressString string = "0x43f8D77300Cf8283D0FB1919CDCF9a0402e9c02A"

func main() {
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

	randomTx := message.NewRandomTransaction()
	randomEthTx := randomTx.AsEthTx()

	transactions := make([]*types.Transaction, 0, 5)
	transactions = append(transactions, randomEthTx)

	fmt.Println("\nBatching transactions...")
	sequencerBatchItems := sequencer.CreateMultipleBatches(transactions)
	fmt.Println("Number of batches created: " + strconv.Itoa(len(sequencerBatchItems)))
	//fmt.Println("Total delayed count: " + sequencerBatchItems[0].TotalDelayedCount.String())

	fmt.Println("\nPublishing Batches...")
	didPublish, err := sequencer.PublishBatchToInbox(sequencerBatchItems)
	if err != nil {
		fmt.Println("Error publishing batch to inbox!")
	}
	fmt.Println("Batch published: " + strconv.FormatBool(didPublish))
}