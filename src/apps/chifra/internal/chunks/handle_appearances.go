// Copyright 2021 The TrueBlocks Authors. All rights reserved.
// Use of this source code is governed by a license that can
// be found in the LICENSE file.

package chunksPkg

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/file"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/index"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/output"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/walk"
)

func (opts *ChunksOptions) HandleAppearances(rCtx *output.RenderCtx, blockNums []base.Blknum) error {
	chain := opts.Globals.Chain

	fetchData := func(modelChan chan types.Modeler, errorChan chan error) {
		showAppearances := func(walker *walk.CacheWalker, path string, first bool) (bool, error) {
			_ = first
			if path != index.ToBloomPath(path) {
				return false, fmt.Errorf("should not happen in showAppearances")
			}

			path = index.ToIndexPath(path)
			if !file.FileExists(path) {
				// Bloom files exist, but index files don't. It's okay.
				return true, nil
			}

			indexChunk, err := index.OpenIndex(path, true /* check */)
			if err != nil {
				return false, err
			}
			defer indexChunk.Close()

			_, err = indexChunk.File.Seek(indexChunk.AppTableStart, io.SeekStart)
			if err != nil {
				return false, err
			}

			for i := 0; i < int(indexChunk.Header.AppearanceCount); i++ {
				if opts.Globals.TestMode && i > walker.MaxTests() {
					continue
				}
				rec := types.AppRecord{}
				if err := binary.Read(indexChunk.File, binary.LittleEndian, &rec); err != nil {
					return false, err
				}

				s := types.Appearance{
					BlockNumber:      rec.BlockNumber,
					TransactionIndex: rec.TransactionIndex,
				}

				modelChan <- &s
			}

			return true, nil
		}

		walker := walk.NewCacheWalker(
			chain,
			opts.Globals.TestMode,
			10, /* maxTests */
			showAppearances,
		)

		if err := walker.WalkBloomFilters(blockNums); err != nil {
			errorChan <- err
			rCtx.Cancel()
		}
	}

	extraOpts := map[string]any{
		"appearances": true,
	}
	return output.StreamMany(rCtx, fetchData, opts.Globals.OutputOptsWithExtra(extraOpts))
}
