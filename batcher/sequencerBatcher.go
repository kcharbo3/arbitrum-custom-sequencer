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
	"reflect"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
)

const SequencerInboxAddressString = "0xDF634a2801b75A8a2d6FBA57BE478A2111db44d3"

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

func (b *SequencerBatcher) BatchTransactions(txs []*types.Transaction, prevAcc common.Hash, prevMsgCount *big.Int) (SequencerBatchItem, int) {

	var l2BatchContents []message.AbstractL2Message
	var batchDataSize int
	finalTransactionProcessed := 0

	// make sure each tx is under limit
	for index, tx := range txs {
		if len(tx.Data()) > maxTxDataSize {
			fmt.Println("Transaction too large, skipping.")
			continue
		}
		if batchDataSize+len(tx.Data()) > maxTxDataSize {
			finalTransactionProcessed = index - 1
			fmt.Println("Transaction batch too large, ending batch at tx (inclusive) at position: " + strconv.Itoa(finalTransactionProcessed))
			break
		}

		l2BatchContents = append(l2BatchContents, message.NewCompressedECDSAFromEth(tx))
		batchDataSize += len(tx.Data())
	}

	batch, err := message.NewTransactionBatchFromMessages(l2BatchContents)

	if err != nil {
		fmt.Println("Error converting l2BatchContents into tx batch")
	}

	l2Message := message.NewSafeL2Message(batch)

	// TODO: pull chain time or take in chain time as parameter
	chainTime := inbox.NewRandomChainTime()

	msgCount := prevMsgCount.Add(prevMsgCount, big.NewInt(1))

	// convert l2Message to seqMsg
	seqMsg := message.NewInboxMessage(l2Message, b.SequencerAddress, new(big.Int).Set(msgCount), big.NewInt(0), chainTime)
	txBatchItem := NewSequencerItem(big.NewInt(0), seqMsg, prevAcc)

	return txBatchItem, finalTransactionProcessed
}

func (b *SequencerBatcher) CreateMultipleBatches(txs []*types.Transaction) []SequencerBatchItem {
	// get sequencer inbox length
	seqInboxLength, err := b.contract.GetInboxAccsLength(nil)
	if err != nil {
		fmt.Println("Error retrieving sequencer inbox length from sequencer inbox contract")
		return nil
	}
	fmt.Println("Sequencer Inbox Length: " + seqInboxLength.String())

	// get prev acc
	var prevAcc common.Hash = common.HexToHash("0x00")
	if seqInboxLength.Cmp(big.NewInt(0)) > 0 {
		prevAcc, err = b.contract.InboxAccs(nil, seqInboxLength.Sub(seqInboxLength, big.NewInt(1)))
		if err != nil {
			fmt.Println("Error retrieving last inbox acc from sequencer inbox contract")
			return nil
		}
	}
	fmt.Println("Previous acc: " + prevAcc.String())

	// get prev msg count
	prevMsgCount, err := b.contract.MessageCount(nil)
	if err != nil {
		fmt.Println("Error retrieving message count from the sequencer inbox contract")
		return nil
	}
	fmt.Println("Previous message count: " + prevMsgCount.String())

	var seqBatchItems []SequencerBatchItem
	numTxs := len(txs)
	lastTxProcessed := 0
	for lastTxProcessed < numTxs {
		seqBatchItem, txCount := b.BatchTransactions(txs[lastTxProcessed:], prevAcc, prevMsgCount)
		if txCount == -1 {
			fmt.Println("BatchTransactions returned negative tx processed count. " +
				"This can happen if first tx is larger than max batch size.")
			break
		}

		lastTxProcessed++
		seqBatchItems = append(seqBatchItems, seqBatchItem)
		prevAcc = seqBatchItem.Accumulator
		prevMsgCount = seqBatchItem.LastSeqNum
	}

	fmt.Println("Final acc: " + prevAcc.String())
	fmt.Println("Final seq num (total batches): " + prevMsgCount.String())

	return seqBatchItems
}

const gasCostBase int = 70292
const gasCostPerMessage int = 1431
const gasCostPerMessageByte int = 16

func (b *SequencerBatcher) PublishBatchToInbox(seqBatchItems []SequencerBatchItem) (bool, error) {
	var transactionsData []byte
	var transactionsLengths []*big.Int
	var metadata []*big.Int
	var startDelayedMessagesRead *big.Int
	var l1BlockNumber *big.Int
	var l1Timestamp *big.Int
	var lastAcc common.Hash
	var lastSeqNum *big.Int
	estimatedGasCost := gasCostBase
	lastMetadataEnd := 0

	// TODO: figure out startDelayedMessagesRead
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
		}

		transactionsData = append(transactionsData, seqMsg.Data...)
		transactionsLengths = append(transactionsLengths, big.NewInt(int64(len(seqMsg.Data))))

		lastAcc = batch.Accumulator
		lastSeqNum = batch.LastSeqNum
	}

	lastSectionCount := len(transactionsLengths) - lastMetadataEnd
	if lastSectionCount > 0 {
		metadata = append(metadata, big.NewInt(int64(lastSectionCount)), l1BlockNumber, l1Timestamp, startDelayedMessagesRead, big.NewInt(0))
	}

	fmt.Println("Last acc: " + lastAcc.String())
	fmt.Println("Last seq num: " + lastSeqNum.String())
	/*
	fmt.Println("Number of batches included: " + strconv.Itoa(len(transactionsLengths)))
	fmt.Println("Metadata length: " + strconv.Itoa(len(metadata)))
	fmt.Println("Printing metadata...")
	for i := 0; i < len(metadata); i += 5 {
		fmt.Println("Section: " + strconv.Itoa(i + 1))

		fmt.Println("Num items: " + metadata[i].String())
		fmt.Println("L1 block number: " + metadata[i + 1].String())
		fmt.Println("L1 timestamp: " + metadata[i + 2].String())
		fmt.Println("NewTotalDelayedMessagesRead: " + metadata[i + 3].String())
		fmt.Println("newDelayedAcc: " + metadata[i + 4].String())
	} */

	b.prepareAuthForTransaction()
	var lastAccBytes [32]byte
	lastAccBytes = lastAcc
	testReflection(transactionsData, transactionsLengths, metadata, lastAccBytes)
	/*
	tx, err := b.contract.AddSequencerL2BatchFromOrigin(b.auth, transactionsData, transactionsLengths, metadata, lastAccBytes)

	if err != nil {
		fmt.Println("Error sending batches to sequencer inbox")
		return false, err
	}
	fmt.Println("Transaction hash: " + tx.Hash().Hex()) */

	return true, nil


	/*
	// Call transaction
	arbTx, err := ethbridge.AddSequencerL2BatchFromOriginCustomNonce(ctx, b.sequencerInbox, b.auth, nonce, transactionsData, transactionsLengths, metadata, lastAcc, b.gasRefunderAddress, b.config.Node.Sequencer.GasRefunderExtraGas)
	if err != nil {
		return false, err
	}

	// Update prevMsgCount for the next iteration if we're not publishingAllBatchItems
	// AddSequencerL2BatchFromOriginCustomNonce will have already updated the nonce
	prevMsgCount.Set(newMsgCount)*/
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

// use godot package to load/read the .env file and
// return the value of the key
func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load("../../.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func testReflection(params ...interface{}) {

	for _, a := range params {
		value := reflect.ValueOf(a)
		elemType := value.Elem().Type()
		fmt.Println(elemType.String())
	}
}

/*
func (b *SequencerBatcher) handleBatchReceiptLog(rawLog *types.Log) {
	if rawLog.Address == b.gasRefunderAddress && rawLog.Topics[0] == refundGasCostsDeniedEventID {
		parsedLog, err := b.gasRefunder.ParseRefundGasCostsDenied(*rawLog)
		if err != nil {
			logger.Warn().Err(err).Msg("failed to parse RefundGasCostsDenied log")
			return
		}
		var loggingLog *zerolog.Event
		if parsedLog.Reason == 0 || parsedLog.Reason == 1 {
			// CONTRACT_NOT_ALLOWED or REFUNDEE_NOT_ALLOWED
			loggingLog = logger.Error()
		} else {
			loggingLog = logger.Warn()
		}
		loggingLog.Int("reason", int(parsedLog.Reason)).Str("txHash", rawLog.TxHash.String()).Msg("batch posting gas costs refund denied")
	}
} */
