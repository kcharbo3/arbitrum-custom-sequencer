/*
 * Copyright 2020-2021, Offchain Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package batcher

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	common2 "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"math/big"
	"os"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
)

// CUSTOM SEQUENCER - don't need delayedMessagesRead > 1 (can be 0)
const SequencerInboxAddressString = "0x140318708284EF8fF2387995bb5AB9f349bAD04E"

type SequencerBatcher struct {
	SequencerAddress common.Address
	EthClient *ethclient.Client
	privateKey *ecdsa.PrivateKey
	auth *bind.TransactOpts
	contract *Batcher
}

func NewSequencerBatcher(address string, ethClient *ethclient.Client) (*SequencerBatcher, error) {
	batcher := &SequencerBatcher{
		SequencerAddress: common.HexToAddress(address),
		EthClient: ethClient,
	}

	privKeyString := goDotEnvVariable("SEQUENCER_PRIVATE_KEY")
	privateKey, err := crypto.HexToECDSA(privKeyString)
	if err != nil {
		fmt.Println("Error converting private key string into private key")
		return nil, err
	}
	batcher.privateKey = privateKey

	chainID, err := batcher.EthClient.NetworkID(context.Background())
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		fmt.Println("Error creating auth from private key")
		return nil, err
	}

	batcher.auth = auth

	batcher.contract, err = NewBatcher(common2.HexToAddress(SequencerInboxAddressString), batcher.EthClient)
	if err != nil {
		fmt.Println("Error creating instance of sequencer inbox contract")
		return nil, err
	}

	return batcher, nil
}

const maxTxDataSize int = 100_000


func (b *SequencerBatcher) BatchTransactions(txs []*types.Transaction, prevAcc common.Hash, prevMsgCount big.Int, l1BlockNumber *big.Int, l1Timestamp *big.Int, useMultipleSeqMsgs bool) (SequencerBatchItem, int) {
	var l2BatchContents []message.AbstractL2Message
	var batchDataSize int
	numTxsProcessed := 0

	// make sure each tx is under limit
	for index, tx := range txs {
		if len(tx.Data()) > maxTxDataSize {
			fmt.Println("Transaction too large, skipping.")
			continue
		}
		if batchDataSize+len(tx.Data()) > maxTxDataSize {
			numTxsProcessed = index
			fmt.Println("Transaction batch too large, ending batch after adding (inclusive) " + strconv.Itoa(numTxsProcessed) + " transactions")
			break
		}
		numTxsProcessed = index + 1
		l2BatchContents = append(l2BatchContents, message.NewCompressedECDSAFromEth(tx))
		batchDataSize += len(tx.Data())

		if useMultipleSeqMsgs {
			break
		}
	}

	batch, err := message.NewTransactionBatchFromMessages(l2BatchContents)

	if err != nil {
		fmt.Println("Error converting l2BatchContents into tx batch")
	}

	l2Message := message.NewSafeL2Message(batch)

	chainTime := inbox.NewRandomChainTime()
	chainTime.BlockNum = common.NewTimeBlocks(l1BlockNumber)
	chainTime.Timestamp = l1Timestamp

	// convert l2Message to seqMsg
	seqMsg := message.NewInboxMessage(l2Message, b.SequencerAddress, &prevMsgCount, big.NewInt(0), chainTime)
	txBatchItem := NewSequencerItem(big.NewInt(0), seqMsg, prevAcc)

	return txBatchItem, numTxsProcessed
}

// random constants that work
const blockNumber int64 = 9364214
const timestamp int64 = 1231231231
func (b *SequencerBatcher) CreateMultipleBatches(txs []*types.Transaction, useMultipleSeqMsgs bool, useMultipleSections bool) []SequencerBatchItem {
	// get current sequencer inbox length
	prevAcc, err := b.getPrevAcc()
	if err != nil {
		fmt.Println("Error getting prev acc: " + err.Error())
		return nil
	}
	fmt.Println("Previous acc: " + prevAcc.String())

	// get prev msg count
	prevMsgCount, err := b.contract.MessageCount(nil)
	if err != nil {
		fmt.Println("Error retrieving message count from the sequencer inbox contract")
		return nil
	}
	fmt.Println("Previous message count: " + prevMsgCount.String())

	l1BlockNumber := big.NewInt(blockNumber)
	l1Timestamp := big.NewInt(timestamp)

	var seqBatchItems []SequencerBatchItem
	numTxs := len(txs)
	nextTxToProcessIndex := 0
	for nextTxToProcessIndex < numTxs {
		seqBatchItem, numTxsProcessed := b.BatchTransactions(txs[nextTxToProcessIndex:], prevAcc, *prevMsgCount, l1BlockNumber, l1Timestamp, useMultipleSeqMsgs)
		if numTxsProcessed == 0 {
			fmt.Println("BatchTransactions did not batch any txs. " +
				"This can happen if first tx is larger than max batch size.")
			break
		}

		nextTxToProcessIndex += numTxsProcessed
		seqBatchItems = append(seqBatchItems, seqBatchItem)
		prevAcc = seqBatchItem.Accumulator
		prevMsgCount = prevMsgCount.Add(prevMsgCount, big.NewInt(1))

		if useMultipleSections {
			l1BlockNumber = l1BlockNumber.Add(l1BlockNumber, big.NewInt(1))
			l1Timestamp = l1Timestamp.Add(l1Timestamp, big.NewInt(1))
		}
	}

	return seqBatchItems
}

const gasCostBase int = 70292
const gasCostPerMessage int = 1431
const gasCostPerMessageByte int = 16

func (b *SequencerBatcher) PublishBatchToInbox(seqBatchItems []SequencerBatchItem) (bool, error) {
	var transactionsData []byte
	var transactionsLengths []*big.Int
	var metadata []*big.Int
	var startDelayedMessagesRead *big.Int = big.NewInt(0)
	var l1BlockNumber *big.Int
	var l1Timestamp *big.Int
	var lastAcc common.Hash
	var lastSeqNum *big.Int
	estimatedGasCost := gasCostBase
	lastMetadataEnd := 0

	for _, batch := range seqBatchItems {
		if len(batch.SequencerMessage) >= 128*1024 {
			// batch too large
			return false, errors.New("batch too big")
		}

		seqMsg, err := inbox.NewInboxMessageFromData(batch.SequencerMessage)
		if err != nil {
			fmt.Println("Error converting batch.SequencerMessage into an inbox message")
			break
		}

		estimatedGasCost += gasCostPerMessage + gasCostPerMessageByte*len(seqMsg.Data)
		if l1BlockNumber == nil {
			l1BlockNumber = seqMsg.ChainTime.BlockNum.AsInt()
			l1Timestamp = seqMsg.ChainTime.Timestamp
		} else if l1BlockNumber.Cmp(seqMsg.ChainTime.BlockNum.AsInt()) != 0 || l1Timestamp.Cmp(seqMsg.ChainTime.Timestamp) != 0 {
			sectionCount := len(transactionsLengths) - lastMetadataEnd
			// [numItems, l1BlockNumber, l1Timestamp, newTotalDelayedMessagesRead, newDelayedAcc]
			metadata = append(metadata, big.NewInt(int64(sectionCount)), l1BlockNumber, l1Timestamp, startDelayedMessagesRead, big.NewInt(0))
			lastMetadataEnd = len(transactionsLengths)

			l1BlockNumber = seqMsg.ChainTime.BlockNum.AsInt()
			l1Timestamp = seqMsg.ChainTime.Timestamp
		}

		transactionsData = append(transactionsData, seqMsg.Data...)
		transactionsLengths = append(transactionsLengths, big.NewInt(int64(len(seqMsg.Data))))

		lastAcc = batch.Accumulator
		lastSeqNum = batch.LastSeqNum
	}

	fmt.Println("Seq num/ message count after adding these batches: " + lastSeqNum.String())
	fmt.Println("Estimated gas cost: " + strconv.Itoa(estimatedGasCost))

	lastSectionCount := len(transactionsLengths) - lastMetadataEnd
	if lastSectionCount > 0 {
		metadata = append(metadata, big.NewInt(int64(lastSectionCount)), l1BlockNumber, l1Timestamp, startDelayedMessagesRead, big.NewInt(0))
	}

	b.prepareAuthForTransaction()
	tx, err := b.contract.AddSequencerL2BatchFromOrigin(b.auth, transactionsData, transactionsLengths, metadata, lastAcc)
	if err != nil {
		fmt.Println("Error sending batches to sequencer inbox: " + err.Error())
		return false, err
	}

	txSuccess, err := b.logTransactionResults(tx)
	if err != nil {
		fmt.Println("Error logging transaction results")
		return false, err
	}

	// compare accs
	fmt.Println("\nChecking if accs match...")
	fmt.Println("Locally calculated acc: " + lastAcc.String())
	contractAcc, err := b.getPrevAcc()
	if err != nil {
		fmt.Println("Error calling getPrevAcc for the contract acc")
	}
	fmt.Println("Contract calculated acc: " + contractAcc.String())

	return txSuccess, nil
}

func (b SequencerBatcher) prepareAuthForTransaction() {
	nonce, err := b.EthClient.PendingNonceAt(context.Background(), common2.Address(b.SequencerAddress))
	if err != nil {
		fmt.Println("Error retrieving sequencer nonce")
		return
	}

	gasPrice, err := b.EthClient.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Println("Error retrieving suggested gas price")
		return
	}

	b.auth.Nonce = big.NewInt(int64(nonce))
	b.auth.Value = big.NewInt(0)     // in wei
	//b.auth.GasLimit = uint64(300000) // in units
	b.auth.GasPrice = gasPrice
}

func (b SequencerBatcher) getPrevAcc() (common.Hash, error) {
	seqInboxLength, err := b.contract.GetInboxAccsLength(nil)
	if err != nil {
		fmt.Println("Error retrieving sequencer inbox length from sequencer inbox contract")
		return *new([32]byte), err
	}

	// get prev acc
	var prevAcc common.Hash = common.HexToHash("0x00")
	if seqInboxLength.Cmp(big.NewInt(0)) > 0 {
		prevAcc, err = b.contract.InboxAccs(nil, seqInboxLength.Sub(seqInboxLength, big.NewInt(1)))
		if err != nil {
			fmt.Println("Error retrieving last inbox acc from sequencer inbox contract")
			return *new([32]byte), err
		}
	}
	return prevAcc, nil
}

func (b SequencerBatcher) logTransactionResults(tx *types.Transaction) (bool, error){
	receipt, err := b.EthClient.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		fmt.Println("Error getting transaction receipt")
		fmt.Println("Error: " + err.Error())
	}

	counter := 0
	const maxTries int = 5
	for receipt == nil && counter < maxTries {
		time.Sleep(10 * time.Second)
		receipt, err = b.EthClient.TransactionReceipt(context.Background(), tx.Hash())
		counter++
		if err != nil {
			fmt.Println("Error getting transaction receipt on attempt: " + strconv.Itoa(counter))
			fmt.Println("Error: " + err.Error())

			if counter == maxTries {
				fmt.Println("Error getting transaction receipt. Max attempts reached.")
				return false, err
			}
		}
	}

	fmt.Println("\nTransaction sent!")
	fmt.Println("Transaction hash: " + tx.Hash().Hex())
	fmt.Println("Transaction status: " + strconv.Itoa(int(receipt.Status)))
	fmt.Println("Gas used: " + strconv.Itoa(int(receipt.GasUsed)))
	return true, nil
}

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load("../../.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
