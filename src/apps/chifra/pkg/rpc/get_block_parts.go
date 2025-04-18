// Copyright 2021 The TrueBlocks Authors. All rights reserved.
// Use of this source code is governed by a license that can
// be found in the LICENSE file.

package rpc

import (
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/walk"
)

// GetBlockTimestamp returns the timestamp associated with a given block
func (conn *Connection) GetBlockTimestamp(bn base.Blknum) base.Timestamp {
	block, _ := conn.GetBlockHeaderByNumber(bn)
	return block.Timestamp
}

// GetBlockHashByHash returns a block's hash if it's a valid block
func (conn *Connection) GetBlockHashByHash(hash string) (base.Hash, error) {
	if block, err := conn.getLightBlockFromRpc(base.NOPOSN, base.HexToHash(hash)); err != nil {
		return base.Hash{}, err
	} else {
		err = conn.WriteToCache(block, walk.Cache_Blocks, block.Timestamp)
		return block.Hash, err
	}
}

// GetBlockNumberByHash returns a block's hash if it's a valid block
func (conn *Connection) GetBlockNumberByHash(hash string) (base.Blknum, error) {
	if block, err := conn.getLightBlockFromRpc(base.NOPOSN, base.HexToHash(hash)); err != nil {
		return 0, err
	} else {
		err = conn.WriteToCache(block, walk.Cache_Blocks, block.Timestamp)
		return block.BlockNumber, err
	}
}

// GetBlockHashByNumber returns a block's hash if it's a valid block
func (conn *Connection) GetBlockHashByNumber(bn base.Blknum) (base.Hash, error) {
	if block, err := conn.GetBlockHeaderByNumber(bn); err != nil {
		return base.Hash{}, err
	} else {
		return block.Hash, err
	}
}
