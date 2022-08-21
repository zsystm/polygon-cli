package rpctypes

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	// ethtypes "github.com/ethereum/go-ethereum/core/types"
	ethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/rs/zerolog/log"
)

type (
	RawQuantityResponse string
	RawDataResponse     string
	RawData8Response    string
	RawData20Response   string
	RawData32Response   string
	RawData256Response  string

	RawTransactionResponse struct {
		// blockHash: DATA, 32 Bytes - hash of the block where this transaction was in. null when its pending.
		BlockHash RawData32Response `json:"blockHash"`

		// blockNumber: QUANTITY - block number where this transaction was in. null when its pending.
		BlockNumber RawQuantityResponse `json:"blockNumber"`

		// from: DATA, 20 Bytes - address of the sender.
		From RawData20Response `json:"from"`

		// gas: QUANTITY - gas provided by the sender.
		Gas RawQuantityResponse `json:"gas"`

		// gasPrice: QUANTITY - gas price provided by the sender in Wei.
		GasPrice RawQuantityResponse `json:"gasPrice"`

		// hash: DATA, 32 Bytes - hash of the transaction.
		Hash RawData32Response `json:"hash"`

		// input: DATA - the data send along with the transaction.
		Input RawDataResponse `json:"input"`

		// nonce: QUANTITY - the number of transactions made by the sender prior to this one.
		Nonce RawQuantityResponse `json:"nonce"`

		// to: DATA, 20 Bytes - address of the receiver. null when its a contract creation transaction.
		To RawData20Response `json:"to"`

		// transactionIndex: QUANTITY - integer of the transactions index position in the block. null when its pending.
		TransactionIndex RawQuantityResponse `json:"transactionIndex"`

		// value: QUANTITY - value transferred in Wei.
		Value RawQuantityResponse `json:"value"`

		// v: QUANTITY - ECDSA recovery id
		V RawQuantityResponse `json:"v"`

		// r: QUANTITY - ECDSA signature r
		R RawQuantityResponse `json:"r"`

		// s: QUANTITY - ECDSA signature s
		S RawQuantityResponse `json:"s"`

		// EIP 2718 Type field?
		Type RawQuantityResponse `json:"type"`
	}

	RawBlockResponse struct {
		// number: QUANTITY - the block number. null when its pending block.
		Number RawQuantityResponse `json:"number"`

		// hash: DATA, 32 Bytes - hash of the block. null when its pending block.
		Hash RawData32Response `json:"hash"`

		// parentHash: DATA, 32 Bytes - hash of the parent block.
		ParentHash RawData32Response `json:"parentHash"`

		// nonce: DATA, 8 Bytes - hash of the generated proof-of-work. null when its pending block.
		Nonce RawData8Response `json:"nonce"`

		// sha3Uncles: DATA, 32 Bytes - SHA3 of the uncles data in the block.
		SHA3Uncles RawData32Response `json:"sha3Uncles"`

		// logsBloom: DATA, 256 Bytes - the bloom filter for the logs of the block. null when its pending block.
		LogsBloom RawData256Response `json:"logsBloom"`

		// transactionsRoot: DATA, 32 Bytes - the root of the transaction trie of the block.
		TransactionsRoot RawData32Response `json:"transactionsRoot"`

		// stateRoot: DATA, 32 Bytes - the root of the final state trie of the block.
		StateRoot RawData32Response `json:"stateRoot"`

		// receiptsRoot: DATA, 32 Bytes - the root of the receipts trie of the block.
		ReceiptsRoot RawData32Response `json:"receiptsRoot"`

		// miner: DATA, 20 Bytes - the address of the beneficiary to whom the mining rewards were given.
		Miner RawData20Response `json:"miner"`

		// difficulty: QUANTITY - integer of the difficulty for this block.
		Difficulty RawQuantityResponse `json:"difficulty"`

		// totalDifficulty: QUANTITY - integer of the total difficulty of the chain until this block.
		TotalDifficulty RawQuantityResponse `json:"totalDifficulty"`

		// extraData: DATA - the "extra data" field of this block.
		ExtraData RawDataResponse `json:"extraData"`

		// size: QUANTITY - integer the size of this block in bytes.
		Size RawQuantityResponse `json:"size"`

		// gasLimit: QUANTITY - the maximum gas allowed in this block.
		GasLimit RawQuantityResponse `json:"gasLimit"`

		// gasUsed: QUANTITY - the total used gas by all transactions in this block.
		GasUsed RawQuantityResponse `json:"gasUsed"`

		// timestamp: QUANTITY - the unix timestamp for when the block was collated.
		Timestamp RawQuantityResponse `json:"timestamp"`

		// transactions: Array - Array of transaction objects, or 32 Bytes transaction hashes depending on the last given parameter.
		Transactions []RawTransactionResponse `json:"transactions"`

		// uncles: Array - Array of uncle hashes.
		Uncles []RawData32Response `json:"uncles"`

		// baseFeePerGas: QUANTITY - fixed per block fee
		BaseFeePerGas RawQuantityResponse `json:"baseFeePerGas"`
	}

	RawTxReceipt struct {
		// transactionHash: DATA, 32 Bytes - hash of the transaction.
		TransactionHash RawData32Response `json:"transactionHash"`

		// transactionIndex: QUANTITY - integer of the transactions index position in the block.
		TransactionIndex RawQuantityResponse `json:"transactionIndex"`

		// blockHash: DATA, 32 Bytes - hash of the block where this transaction was in.
		BlockHash RawData32Response `json:"blockHash"`

		// blockNumber: QUANTITY - block number where this transaction was in.
		BlockNumber RawQuantityResponse `json:"blockNumber"`

		// from: DATA, 20 Bytes - address of the sender.
		From RawData20Response `json:"from"`

		// to: DATA, 20 Bytes - address of the receiver. null when its a contract creation transaction.
		To RawDataResponse `json:"to"`

		// cumulativeGasUsed : QUANTITY - The total amount of gas used when this transaction was executed in the block.
		CumulativeGasUsed RawQuantityResponse `json:"cumulativeGasUsed"`

		// gasUsed : QUANTITY - The amount of gas used by this specific transaction alone.
		GasUsed RawQuantityResponse `json:"gasUsed"`

		// contractAddress : DATA, 20 Bytes - The contract address created, if the transaction was a contract creation, otherwise null.
		ContractAddress RawData20Response `json:"contractAddress"`

		// logs: Array - Array of log objects, which this transaction generated.
		Logs []RawDataResponse `json:"logs"`

		// logsBloom: DATA, 256 Bytes - Bloom filter for light clients to quickly retrieve related logs. It also returns either :
		LogsBloom RawData256Response `json:"logsBloom"`

		// root : DATA 32 bytes of post-transaction stateroot (pre Byzantium)
		Root RawData32Response `json:"root"`

		// status: QUANTITY either 1 (success) or 0 (failure)
		Status RawQuantityResponse `json:"status"`
	}

	PolyTransaction interface {
		GasPrice() *big.Int
		Hash() ethcommon.Hash
		To() ethcommon.Address
		From() ethcommon.Address
		Data() []byte
		Value() *big.Int
		Gas() uint64
		Nonce() uint64
		String() string
		MarshalJSON() ([]byte, error)
	}
	PolyTransactions []PolyTransaction
	PolyBlock        interface {
		Number() *big.Int
		Time() uint64
		Transactions() PolyTransactions
		Uncles() []RawData32Response
		Size() uint64
		GasUsed() uint64
		Miner() ethcommon.Address
		Hash() ethcommon.Hash
		Difficulty() *big.Int
		GasLimit() uint64
		BaseFee() *big.Int
		Extra() []byte
		ParentHash() ethcommon.Hash
		UncleHash() ethcommon.Hash
		Root() ethcommon.Hash
		TxHash() ethcommon.Hash
		Nonce() uint64
		String() string
		MarshalJSON() ([]byte, error)
	}

	implPolyBlock struct {
		inner *RawBlockResponse
	}
	implPolyTransaction struct {
		inner *RawTransactionResponse
	}
)

func NewPolyBlock(r *RawBlockResponse) PolyBlock {
	i := new(implPolyBlock)
	i.inner = r
	return i
}
func NewPolyTransaction(r *RawTransactionResponse) PolyTransaction {
	i := new(implPolyTransaction)
	i.inner = r
	return i
}

func (i *implPolyBlock) Number() *big.Int {
	return i.inner.Number.ToBigInt()
}
func (i *implPolyBlock) Difficulty() *big.Int {
	return i.inner.Difficulty.ToBigInt()
}
func (i *implPolyBlock) BaseFee() *big.Int {
	return i.inner.BaseFeePerGas.ToBigInt()
}

func (i *implPolyBlock) Time() uint64 {
	return i.inner.Timestamp.ToUint64()
}
func (i *implPolyBlock) Transactions() PolyTransactions {
	pt := make(PolyTransactions, len(i.inner.Transactions))
	for idx := range i.inner.Transactions {
		pt[idx] = NewPolyTransaction(&i.inner.Transactions[idx])
	}
	return pt
}
func (i *implPolyBlock) Uncles() []RawData32Response {
	return i.inner.Uncles
}
func (i *implPolyBlock) Size() uint64 {
	return i.inner.Size.ToUint64()
}
func (i *implPolyBlock) GasUsed() uint64 {
	return i.inner.GasUsed.ToUint64()
}
func (i *implPolyBlock) GasLimit() uint64 {
	return i.inner.GasLimit.ToUint64()
}
func (i *implPolyBlock) Nonce() uint64 {
	return i.inner.Nonce.ToUint64()
}
func (i *implPolyBlock) Miner() ethcommon.Address {
	return i.inner.Miner.ToAddress()
}
func (i *implPolyBlock) Hash() ethcommon.Hash {
	return i.inner.Hash.ToHash()
}
func (i *implPolyBlock) ParentHash() ethcommon.Hash {
	return i.inner.ParentHash.ToHash()
}
func (i *implPolyBlock) UncleHash() ethcommon.Hash {
	return i.inner.SHA3Uncles.ToHash()
}
func (i *implPolyBlock) Root() ethcommon.Hash {
	return i.inner.StateRoot.ToHash()
}
func (i *implPolyBlock) TxHash() ethcommon.Hash {
	return i.inner.TransactionsRoot.ToHash()
}
func (i *implPolyBlock) Extra() []byte {
	return i.inner.ExtraData.ToBytes()
}
func (i *implPolyBlock) String() string {
	d, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	return string(d)
}
func (i *implPolyBlock) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.inner)
}

func (i *implPolyTransaction) GasPrice() *big.Int {
	return i.inner.GasPrice.ToBigInt()
}
func (i *implPolyTransaction) Gas() uint64 {
	return i.inner.Gas.ToUint64()
}
func (i *implPolyTransaction) Nonce() uint64 {
	return i.inner.Nonce.ToUint64()
}
func (i *implPolyTransaction) Value() *big.Int {
	return i.inner.Value.ToBigInt()
}
func (i *implPolyTransaction) Hash() ethcommon.Hash {
	return i.inner.Hash.ToHash()
}
func (i *implPolyTransaction) To() ethcommon.Address {
	return i.inner.To.ToAddress()
}
func (i *implPolyTransaction) From() ethcommon.Address {
	return i.inner.From.ToAddress()
}
func (i *implPolyTransaction) Data() []byte {
	return i.inner.Input.ToBytes()
}
func (i *implPolyTransaction) String() string {
	d, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	return string(d)
}
func (i *implPolyTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.inner)
}

// HexToBigInt assumes that it's input is a hex encoded string and
// will try to convert it to a big int
func ConvHexToBigInt(raw any) (bi *big.Int, err error) {
	bi = big.NewInt(0)
	hexString, err := rawRespToString(raw)
	if err != nil {
		return nil, err
	}
	hexString = strings.Replace(hexString, "0x", "", -1)
	if len(hexString)%2 != 0 {
		hexString = "0" + hexString
	}

	rawGas, err := hex.DecodeString(hexString)
	if err != nil {
		log.Error().Err(err).Str("hex", hexString).Msg("Unable to decode hex string")
		return
	}
	bi.SetBytes(rawGas)
	return
}

func rawRespToString(raw any) (string, error) {
	var hexString string
	switch v := raw.(type) {
	case RawQuantityResponse:
		hexString = string(v)
	case RawDataResponse:
		hexString = string(v)
	case RawData8Response:
		hexString = string(v)
	case RawData20Response:
		hexString = string(v)
	case RawData32Response:
		hexString = string(v)
	case RawData256Response:
		hexString = string(v)
	case string:
		hexString = v
	default:
		return "", fmt.Errorf("Could not assert %v as a string", raw)
	}
	return hexString, nil
}

// HexToUint64 assumes that its input is a hex encoded string and it
// will attempt to convert this into a uint64
func ConvHexToUint64(raw any) (uint64, error) {
	hexString, err := rawRespToString(raw)
	if err != nil {
		return 0, err
	}

	hexString = strings.Replace(hexString, "0x", "", -1)
	if len(hexString)%2 != 0 {
		hexString = "0" + hexString
	}

	result, err := strconv.ParseUint(hexString, 16, 64)
	if err != nil {
		return 0, err
	}
	return uint64(result), nil
}

func MustConvHexToUint64(raw any) uint64 {
	num, err := ConvHexToUint64(raw)
	if err != nil {
		panic(fmt.Sprintf("Failed to covert Hex to uint64: %v", err))
	}
	return num
}

func NewRawBlockResponseFromAny(raw any) (*RawBlockResponse, error) {
	topMap, ok := raw.(map[string]any)
	if !ok {
		return nil, fmt.Errorf("Unable to map raw response")
	}
	_ = topMap
	return nil, nil

}

func normalizeHexString(s string) string {
	hexString := strings.Replace(s, "0x", "", -1)
	if len(hexString)%2 != 0 {
		hexString = "0" + hexString
	}
	return hexString
}
func (r RawData8Response) ToUint64() uint64 {
	hexString := normalizeHexString(string(r))
	result, err := strconv.ParseUint(hexString, 16, 64)
	if err != nil {
		return 0
	}
	return uint64(result)
}

func (r RawQuantityResponse) ToUint64() uint64 {
	hexString := normalizeHexString(string(r))
	result, err := strconv.ParseUint(hexString, 16, 64)
	if err != nil {
		return 0
	}
	return uint64(result)
}
func (r RawQuantityResponse) ToFloat64() float64 {
	hexString := normalizeHexString(string(r))
	result, err := strconv.ParseFloat(hexString, 64)
	if err != nil {
		return 0
	}
	return result
}

func (r RawQuantityResponse) ToInt64() int64 {
	hexString := normalizeHexString(string(r))
	result, err := strconv.ParseUint(hexString, 16, 64)
	if err != nil {
		return 0
	}
	return int64(result)

}

func (r *RawQuantityResponse) ToBigInt() *big.Int {
	hexString := normalizeHexString(string(*r))
	bi := new(big.Int)
	bi.SetString(hexString, 16)
	return bi
}
func (r *RawQuantityResponse) String() string {
	return r.ToBigInt().String()
}

func (r *RawData20Response) ToAddress() ethcommon.Address {
	return ethcommon.HexToAddress(string(*r))
}
func (r *RawData32Response) ToHash() ethcommon.Hash {
	return ethcommon.HexToHash(string(*r))
}
func (r *RawDataResponse) ToBytes() []byte {
	hexString := normalizeHexString(string(*r))
	data, err := hex.DecodeString(hexString)
	if err != nil {
		log.Error().Err(err).Msg("Unable to covert raw data to bytes")
		return nil
	}
	return data
}