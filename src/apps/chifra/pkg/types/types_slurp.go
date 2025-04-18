// Copyright 2016, 2024 The TrueBlocks Authors. All rights reserved.
// Use of this source code is governed by a license that can
// be found in the LICENSE file.
/*
 * Parts of this file were auto generated. Edit only those parts of
 * the code inside of 'EXISTING_CODE' tags.
 */

package types

// EXISTING_CODE
import (
	"encoding/json"
	"fmt"
	"io"
	"path/filepath"
	"strings"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/cache"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// EXISTING_CODE

type Slurp struct {
	ArticulatedTx     *Function      `json:"articulatedTx"`
	BlockHash         base.Hash      `json:"blockHash"`
	BlockNumber       base.Blknum    `json:"blockNumber"`
	ContractAddress   base.Address   `json:"contractAddress"`
	CumulativeGasUsed string         `json:"cumulativeGasUsed"`
	From              base.Address   `json:"from"`
	FunctionName      string         `json:"functionName"`
	Gas               base.Gas       `json:"gas"`
	GasPrice          base.Gas       `json:"gasPrice"`
	GasUsed           base.Gas       `json:"gasUsed"`
	HasToken          bool           `json:"hasToken"`
	Hash              base.Hash      `json:"hash"`
	Input             string         `json:"input"`
	IsError           bool           `json:"-"`
	MethodId          string         `json:"methodId"`
	Nonce             base.Value     `json:"nonce"`
	Timestamp         base.Timestamp `json:"timestamp"`
	To                base.Address   `json:"to"`
	TransactionIndex  base.Txnum     `json:"transactionIndex"`
	TxReceiptStatus   string         `json:"txReceiptStatus"`
	ValidatorIndex    base.Value     `json:"validatorIndex"`
	Value             base.Wei       `json:"value"`
	WithdrawalIndex   base.Value     `json:"withdrawalIndex"`
	// EXISTING_CODE
	Amount base.Wei `json:"amount"`
	// EXISTING_CODE
}

func (s Slurp) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

func (s *Slurp) Model(chain, format string, verbose bool, extraOpts map[string]any) Model {
	_ = chain
	_ = format
	_ = verbose
	_ = extraOpts
	var model = map[string]any{}
	var order = []string{}

	// EXISTING_CODE
	to := hexutil.Encode(s.To.Bytes())
	if to == "0x0000000000000000000000000000000000000000" {
		to = "0x0" // weird special case to preserve what RPC does
	}

	model = map[string]any{
		"blockNumber": s.BlockNumber,
		"from":        s.From,
		"timestamp":   s.Timestamp,
		"date":        s.Date(),
		"to":          s.To,
		"value":       s.Value.String(),
	}

	if s.From == base.BlockRewardSender || s.From == base.UncleRewardSender {
		model["from"] = s.From.Hex()
		s.Input = ""
		order = []string{
			"blockNumber",
			"timestamp",
			"date",
			"from",
			"to",
			"value",
		}

	} else if s.From == base.WithdrawalSender {
		model["from"] = s.From.Hex()
		s.Input = ""
		model["withdrawalIndex"] = s.WithdrawalIndex
		model["validatorIndex"] = s.ValidatorIndex
		order = []string{
			"blockNumber",
			"validatorIndex",
			"withdrawalIndex",
			"timestamp",
			"date",
			"from",
			"to",
			"value",
		}

	} else {
		order = []string{
			"blockNumber",
			"transactionIndex",
			"timestamp",
			"date",
			"from",
			"to",
			"hasToken",
			"isError",
			"hash",
			"gasPrice",
			"gasUsed",
			"gasCost",
			"value",
			"input",
		}

		model["gas"] = s.Gas
		model["gasCost"] = s.GasCost()
		model["gasPrice"] = s.GasPrice
		model["gasUsed"] = s.GasUsed
		model["hash"] = s.Hash
	}

	if s.HasToken {
		model["hasToken"] = s.HasToken
	}
	if s.IsError {
		model["isError"] = s.IsError
	}
	if s.BlockHash != base.HexToHash("0xdeadbeef") && !s.BlockHash.IsZero() {
		model["blockHash"] = s.BlockHash
	}
	model["transactionIndex"] = s.TransactionIndex

	// TODO: Turn this back on
	var articulatedTx map[string]any
	isArticulated := extraOpts["articulate"] == true && s.ArticulatedTx != nil
	if isArticulated && format != "json" {
		order = append(order, "compressedTx")
	}

	// TODO: ARTICULATE SLURP
	if isArticulated {
		articulatedTx = map[string]any{
			"name": s.ArticulatedTx.Name,
		}
		inputModels := parametersToMap(s.ArticulatedTx.Inputs)
		if inputModels != nil {
			articulatedTx["inputs"] = inputModels
		}
		outputModels := parametersToMap(s.ArticulatedTx.Outputs)
		if outputModels != nil {
			articulatedTx["outputs"] = outputModels
		}
		sm := s.ArticulatedTx.StateMutability
		if sm != "" && sm != "nonpayable" && sm != "view" {
			articulatedTx["stateMutability"] = sm
		}
	}

	if format == "json" {
		a := s.ContractAddress.Hex()
		if strings.HasPrefix(a, "0x") && len(a) == 42 {
			model["contractAddress"] = a
		}
		// etherscan sometimes returns this word instead of a hex address
		if len(s.Input) > 2 && s.Input != "deprecated" {
			model["input"] = s.Input
		}

		// TODO: ARTICULATE SLURP
		if isArticulated {
			model["articulatedTx"] = articulatedTx
			// } else {
			// 	if s.Message != "" {
			// 		model["message"] = s.Message
			// 	}
		}

	} else {
		model["hasToken"] = s.HasToken
		model["isError"] = s.IsError
		// etherscan sometimes returns this word instead of a hex address
		if s.Input == "deprecated" {
			s.Input = "0x"
		}
		model["input"] = s.Input

		// TODO: ARTICULATE SLURP
		model["compressedTx"] = ""
		enc := s.Input
		if len(s.Input) >= 10 {
			enc = s.Input[:10]
		}
		model["encoding"] = enc

		if isArticulated {
			model["compressedTx"] = MakeCompressed(articulatedTx)
			// } else if s.Message != "" {
			// 	model["encoding"] = ""
			// 	model["compressedTx"] = s.Message
		}
	}

	asEther := true // like transactions, we always export ether for slurps -- extraOpts["ether"] == true
	if asEther {
		model["ether"] = s.Value.ToFloatString(18)
		order = append(order, "ether")
	}

	items := []namer{
		{addr: s.From, name: "fromName"},
		{addr: s.To, name: "toName"},
		{addr: s.ContractAddress, name: "contractName"},
	}
	for _, item := range items {
		if name, loaded, found := nameAddress(extraOpts, item.addr); found {
			model[item.name] = name.Name
			order = append(order, item.name)
		} else if loaded && format != "json" {
			model[item.name] = ""
			order = append(order, item.name)
		}
	}
	order = reorderOrdering(order)
	// EXISTING_CODE

	return Model{
		Data:  model,
		Order: order,
	}
}

func (s *Slurp) Date() string {
	return base.FormattedDate(s.Timestamp)
}

func (s *SlurpGroup) CacheLocations() (string, string, string) {
	paddedId := fmt.Sprintf("%s-%09d-%05d", s.Address.Hex()[2:], s.BlockNumber, s.TransactionIndex)
	parts := make([]string, 3)
	parts[0] = paddedId[:2]
	parts[1] = paddedId[2:4]
	parts[2] = paddedId[4:6]
	subFolder := strings.ToLower("Slurp") + "s"
	directory := filepath.Join(subFolder, filepath.Join(parts...))
	return directory, paddedId, "bin"
}

type SlurpGroup struct {
	BlockNumber      base.Blknum
	TransactionIndex base.Txnum
	Address          base.Address
	Slurps           []Slurp
}

func (s *SlurpGroup) MarshalCache(writer io.Writer) (err error) {
	return base.WriteValue(writer, s.Slurps)
}

func (s *SlurpGroup) UnmarshalCache(fileVersion uint64, reader io.Reader) (err error) {
	return base.ReadValue(reader, &s.Slurps, fileVersion)
}

func (s *Slurp) MarshalCache(writer io.Writer) (err error) {
	// ArticulatedTx
	optArticulatedTx := &cache.Optional[Function]{
		Value: s.ArticulatedTx,
	}
	if err = base.WriteValue(writer, optArticulatedTx); err != nil {
		return err
	}

	// BlockHash
	if err = base.WriteValue(writer, &s.BlockHash); err != nil {
		return err
	}

	// BlockNumber
	if err = base.WriteValue(writer, s.BlockNumber); err != nil {
		return err
	}

	// ContractAddress
	if err = base.WriteValue(writer, s.ContractAddress); err != nil {
		return err
	}

	// CumulativeGasUsed
	if err = base.WriteValue(writer, s.CumulativeGasUsed); err != nil {
		return err
	}

	// From
	if err = base.WriteValue(writer, s.From); err != nil {
		return err
	}

	// FunctionName
	if err = base.WriteValue(writer, s.FunctionName); err != nil {
		return err
	}

	// Gas
	if err = base.WriteValue(writer, s.Gas); err != nil {
		return err
	}

	// GasPrice
	if err = base.WriteValue(writer, s.GasPrice); err != nil {
		return err
	}

	// GasUsed
	if err = base.WriteValue(writer, s.GasUsed); err != nil {
		return err
	}

	// HasToken
	if err = base.WriteValue(writer, s.HasToken); err != nil {
		return err
	}

	// Hash
	if err = base.WriteValue(writer, &s.Hash); err != nil {
		return err
	}

	// Input
	if err = base.WriteValue(writer, s.Input); err != nil {
		return err
	}

	// IsError
	if err = base.WriteValue(writer, s.IsError); err != nil {
		return err
	}

	// MethodId
	if err = base.WriteValue(writer, s.MethodId); err != nil {
		return err
	}

	// Nonce
	if err = base.WriteValue(writer, s.Nonce); err != nil {
		return err
	}

	// Timestamp
	if err = base.WriteValue(writer, s.Timestamp); err != nil {
		return err
	}

	// To
	if err = base.WriteValue(writer, s.To); err != nil {
		return err
	}

	// TransactionIndex
	if err = base.WriteValue(writer, s.TransactionIndex); err != nil {
		return err
	}

	// TxReceiptStatus
	if err = base.WriteValue(writer, s.TxReceiptStatus); err != nil {
		return err
	}

	// ValidatorIndex
	if err = base.WriteValue(writer, s.ValidatorIndex); err != nil {
		return err
	}

	// Value
	if err = base.WriteValue(writer, &s.Value); err != nil {
		return err
	}

	// WithdrawalIndex
	if err = base.WriteValue(writer, s.WithdrawalIndex); err != nil {
		return err
	}

	return nil
}

func (s *Slurp) UnmarshalCache(fileVersion uint64, reader io.Reader) (err error) {
	// Check for compatibility and return cache.ErrIncompatibleVersion to invalidate this item (see #3638)
	// EXISTING_CODE
	// EXISTING_CODE

	// ArticulatedTx
	optArticulatedTx := &cache.Optional[Function]{
		Value: s.ArticulatedTx,
	}
	if err = base.ReadValue(reader, optArticulatedTx, fileVersion); err != nil {
		return err
	}
	s.ArticulatedTx = optArticulatedTx.Get()

	// BlockHash
	if err = base.ReadValue(reader, &s.BlockHash, fileVersion); err != nil {
		return err
	}

	// BlockNumber
	if err = base.ReadValue(reader, &s.BlockNumber, fileVersion); err != nil {
		return err
	}

	// ContractAddress
	if err = base.ReadValue(reader, &s.ContractAddress, fileVersion); err != nil {
		return err
	}

	// CumulativeGasUsed
	if err = base.ReadValue(reader, &s.CumulativeGasUsed, fileVersion); err != nil {
		return err
	}

	// From
	if err = base.ReadValue(reader, &s.From, fileVersion); err != nil {
		return err
	}

	// FunctionName
	if err = base.ReadValue(reader, &s.FunctionName, fileVersion); err != nil {
		return err
	}

	// Gas
	if err = base.ReadValue(reader, &s.Gas, fileVersion); err != nil {
		return err
	}

	// GasPrice
	if err = base.ReadValue(reader, &s.GasPrice, fileVersion); err != nil {
		return err
	}

	// GasUsed
	if err = base.ReadValue(reader, &s.GasUsed, fileVersion); err != nil {
		return err
	}

	// HasToken
	if err = base.ReadValue(reader, &s.HasToken, fileVersion); err != nil {
		return err
	}

	// Hash
	if err = base.ReadValue(reader, &s.Hash, fileVersion); err != nil {
		return err
	}

	// Input
	if err = base.ReadValue(reader, &s.Input, fileVersion); err != nil {
		return err
	}

	// IsError
	if err = base.ReadValue(reader, &s.IsError, fileVersion); err != nil {
		return err
	}

	// MethodId
	if err = base.ReadValue(reader, &s.MethodId, fileVersion); err != nil {
		return err
	}

	// Nonce
	if err = base.ReadValue(reader, &s.Nonce, fileVersion); err != nil {
		return err
	}

	// Timestamp
	if err = base.ReadValue(reader, &s.Timestamp, fileVersion); err != nil {
		return err
	}

	// To
	if err = base.ReadValue(reader, &s.To, fileVersion); err != nil {
		return err
	}

	// TransactionIndex
	if err = base.ReadValue(reader, &s.TransactionIndex, fileVersion); err != nil {
		return err
	}

	// TxReceiptStatus
	if err = base.ReadValue(reader, &s.TxReceiptStatus, fileVersion); err != nil {
		return err
	}

	// ValidatorIndex
	if err = base.ReadValue(reader, &s.ValidatorIndex, fileVersion); err != nil {
		return err
	}

	// Value
	if err = base.ReadValue(reader, &s.Value, fileVersion); err != nil {
		return err
	}

	// WithdrawalIndex
	if err = base.ReadValue(reader, &s.WithdrawalIndex, fileVersion); err != nil {
		return err
	}

	s.FinishUnmarshal(fileVersion)

	return nil
}

// FinishUnmarshal is used by the cache. It may be unused depending on auto-code-gen
func (s *Slurp) FinishUnmarshal(fileVersion uint64) {
	_ = fileVersion
	// EXISTING_CODE
	// EXISTING_CODE
}

// EXISTING_CODE
//

func (s *Slurp) GasCost() base.Gas {
	return s.GasPrice * s.GasUsed
}

// EXISTING_CODE
