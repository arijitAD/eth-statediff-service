// Copyright 2019 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

// Contains a batch of utility type declarations used by the tests. As the node
// operates on unique types, a lot of them are needed to check various features.

package statediff

import (
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// Subscription struct holds our subscription channels
type Subscription struct {
	PayloadChan chan<- Payload
	QuitChan    chan<- bool
}

// Params is used to carry in parameters from subscribing/requesting clients configuration
type Params struct {
	IntermediateStateNodes   bool
	IntermediateStorageNodes bool
	IncludeBlock             bool
	IncludeReceipts          bool
	IncludeTD                bool
	WatchedAddresses         []common.Address
	WatchedStorageSlots      []common.Hash
	Workers uint
}

// Args bundles the arguments for the state diff builder
type Args struct {
	OldStateRoot, NewStateRoot, BlockHash common.Hash
	BlockNumber                           *big.Int
}

// Payload packages the data to send to statediff subscriptions
type Payload struct {
	BlockRlp        []byte   `json:"blockRlp"`
	TotalDifficulty *big.Int `json:"totalDifficulty"`
	ReceiptsRlp     []byte   `json:"receiptsRlp"`
	StateObjectRlp  []byte   `json:"stateObjectRlp"    gencodec:"required"`

	encoded []byte
	err     error
}

func (sd *Payload) ensureEncoded() {
	if sd.encoded == nil && sd.err == nil {
		sd.encoded, sd.err = json.Marshal(sd)
	}
}

// Length to implement Encoder interface for Payload
func (sd *Payload) Length() int {
	sd.ensureEncoded()
	return len(sd.encoded)
}

// Encode to implement Encoder interface for Payload
func (sd *Payload) Encode() ([]byte, error) {
	sd.ensureEncoded()
	return sd.encoded, sd.err
}

// StateObject is the final output structure from the builder
type StateObject struct {
	BlockHash   common.Hash `json:"blockHash"       gencodec:"required"`
	BlockNumber *big.Int    `json:"blockNumber"     gencodec:"required"`
	Nodes       []StateNode `json:"nodes"           gencodec:"required"`
}