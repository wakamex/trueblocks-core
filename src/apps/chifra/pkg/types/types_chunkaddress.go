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

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
)

// EXISTING_CODE

type ChunkAddress struct {
	Address    base.Address `json:"address"`
	Count      uint64       `json:"count"`
	Offset     uint64       `json:"offset"`
	Range      string       `json:"range"`
	RangeDates *RangeDates  `json:"rangeDates,omitempty"`
	// EXISTING_CODE
	// EXISTING_CODE
}

func (s ChunkAddress) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

func (s *ChunkAddress) Model(chain, format string, verbose bool, extraOpts map[string]any) Model {
	_ = chain
	_ = format
	_ = verbose
	_ = extraOpts
	var model = map[string]any{}
	var order = []string{}

	// EXISTING_CODE
	model = map[string]any{
		"address": s.Address,
		"range":   s.Range,
		"offset":  s.Offset,
		"count":   s.Count,
	}
	order = []string{
		"address",
		"range",
		"offset",
		"count",
	}

	if verbose && format == "json" {
		if s.RangeDates != nil {
			model["rangeDates"] = s.RangeDates.Model(chain, format, verbose, extraOpts).Data
		}
	} else if verbose {
		model["firstTs"], model["firstDate"], model["lastTs"], model["lastDate"] = 0, "", 0, ""
		order = append(order, []string{"firstTs", "firstDate", "lastTs", "lastDate"}...)
		if s.RangeDates != nil {
			model["firstTs"] = s.RangeDates.FirstTs
			model["firstDate"] = s.RangeDates.FirstDate
			model["lastTs"] = s.RangeDates.LastTs
			model["lastDate"] = s.RangeDates.LastDate
		}
	}
	// EXISTING_CODE

	return Model{
		Data:  model,
		Order: order,
	}
}

// FinishUnmarshal is used by the cache. It may be unused depending on auto-code-gen
func (s *ChunkAddress) FinishUnmarshal(fileVersion uint64) {
	_ = fileVersion
	// EXISTING_CODE
	// EXISTING_CODE
}

// EXISTING_CODE
// EXISTING_CODE
