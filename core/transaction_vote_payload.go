// Copyright (C) 2017 go-nebulas authors
//
// This file is part of the go-nebulas library.
//
// the go-nebulas library is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// the go-nebulas library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with the go-nebulas library.  If not, see <http://www.gnu.org/licenses/>.
//

package core

import (
	"encoding/json"
	"errors"

	"github.com/nebulasio/go-nebulas/util/byteutils"
)

// Action types
const (
	PrepareAction  = "prepare"
	CommitAction   = "commit"
	ChangeAction   = "change"
	AbdicateAction = "abdicate"
)

// Errors constants
var (
	ErrInvalidVoteAction = errors.New("invalid vote action")
)

// VotePayload carry vote information
// 1. action: prepare, block: current block hash, view: based block hash
// 2. action: commit, block: current block hash
// 3. action: change, block: parent block, times: change times
// 4. action: abdicate, block: parent block
type VotePayload struct {
	Action    string
	BlockHash byteutils.Hash
	ViewHash  byteutils.Hash
	Times     int
}

// LoadVotePayload from bytes
func LoadVotePayload(bytes []byte) (*VotePayload, error) {
	payload := &VotePayload{}
	if err := json.Unmarshal(bytes, payload); err != nil {
		return nil, err
	}
	return payload, nil
}

// NewPrepareVotePayload create a new prepare vote payload
func NewPrepareVotePayload(action string, block byteutils.Hash, view byteutils.Hash) *VotePayload {
	return &VotePayload{
		Action:    action,
		BlockHash: block,
		ViewHash:  view,
	}
}

// NewCommitVotePayload create a new commit vote payload
func NewCommitVotePayload(action string, block byteutils.Hash) *VotePayload {
	return &VotePayload{
		Action:    action,
		BlockHash: block,
	}
}

// NewChangeVotePayload create a new change vote payload
func NewChangeVotePayload(action string, block byteutils.Hash, times int) *VotePayload {
	return &VotePayload{
		Action:    action,
		BlockHash: block,
		Times:     times,
	}
}

// NewAbdicateVotePayload create a new abdicate vote payload
func NewAbdicateVotePayload(action string, block byteutils.Hash) *VotePayload {
	return &VotePayload{
		Action:    action,
		BlockHash: block,
	}
}

// ToBytes serialize payload
func (payload *VotePayload) ToBytes() ([]byte, error) {
	return json.Marshal(payload)
}

func (payload *VotePayload) prepare(from []byte, block *Block) error {
	return nil
}

func (payload *VotePayload) commit(from []byte, block *Block) error {
	return nil
}

func (payload *VotePayload) change(from []byte, block *Block) error {
	return nil
}

func (payload *VotePayload) abdicate(from []byte, block *Block) error {
	return nil
}

// Execute the call payload in tx, call a function
func (payload *VotePayload) Execute(tx *Transaction, block *Block) error {
	from := tx.from.Bytes()
	switch payload.Action {
	case PrepareAction:
		return payload.prepare(from, block)
	case CommitAction:
		return payload.commit(from, block)
	case ChangeAction:
		return payload.change(from, block)
	case AbdicateAction:
		return payload.abdicate(from, block)
	default:
		return ErrInvalidVoteAction
	}
}
