// Copyright 2021 The TrueBlocks Authors. All rights reserved.
// Use of this source code is governed by a license that can
// be found in the LICENSE file.

package rpc

import (
	"context"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/walk"
)

// TODO: use types.MetaData throughout

func (conn *Connection) GetMetaData(testmode bool) (*types.MetaData, error) {
	chainId, networkId, err := conn.GetClientIDs()
	if err != nil {
		return nil, err
	}

	if testmode {
		return &types.MetaData{
			Unripe:    0xdeadbeef,
			Ripe:      0xdeadbeef,
			Staging:   0xdeadbeef,
			Finalized: 0xdeadbeef,
			Latest:    0xdeadbeef,
			Chain:     conn.Chain,
			ChainId:   chainId,
			NetworkId: networkId,
		}, nil
	}

	var meta types.MetaData
	meta.Chain = conn.Chain
	meta.ChainId = chainId
	meta.NetworkId = networkId
	meta.Latest = conn.GetLatestBlockNumber()

	filenameChan := make(chan walk.CacheFileInfo)

	var nRoutines = 4
	go walk.WalkCacheFolder(context.Background(), conn.Chain, walk.Index_Bloom, nil, filenameChan)
	go walk.WalkCacheFolder(context.Background(), conn.Chain, walk.Index_Staging, nil, filenameChan)
	go walk.WalkCacheFolder(context.Background(), conn.Chain, walk.Index_Ripe, nil, filenameChan)
	go walk.WalkCacheFolder(context.Background(), conn.Chain, walk.Index_Unripe, nil, filenameChan)

	for result := range filenameChan {
		switch result.Type {
		case walk.Index_Bloom:
			fallthrough
		case walk.Index_Final:
			meta.Finalized = max(meta.Finalized, result.FileRange.Last)
		case walk.Index_Staging:
			meta.Staging = max(meta.Staging, result.FileRange.Last)
		case walk.Index_Ripe:
			meta.Ripe = max(meta.Ripe, result.FileRange.Last)
		case walk.Index_Unripe:
			meta.Unripe = max(meta.Unripe, result.FileRange.Last)
		case walk.Cache_NotACache:
			nRoutines--
			if nRoutines == 0 {
				close(filenameChan)
			}
		}
	}

	if meta.Staging == 0 {
		meta.Staging = meta.Finalized
	}
	meta.Ripe = max(meta.Staging, meta.Ripe)
	meta.Unripe = max(meta.Ripe, meta.Unripe)

	return &meta, nil
}
