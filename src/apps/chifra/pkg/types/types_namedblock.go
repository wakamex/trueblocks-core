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

type NamedBlock struct {
	BlockNumber base.Blknum    `json:"blockNumber"`
	Component   string         `json:"component,omitempty"`
	Description string         `json:"description,omitempty"`
	Name        string         `json:"name,omitempty"`
	Timestamp   base.Timestamp `json:"timestamp"`
	// EXISTING_CODE
	// EXISTING_CODE
}

func (s NamedBlock) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

func (s *NamedBlock) Model(chain, format string, verbose bool, extraOpts map[string]any) Model {
	_ = chain
	_ = format
	_ = verbose
	_ = extraOpts
	var model = map[string]any{}
	var order = []string{}

	// EXISTING_CODE
	model = map[string]any{
		"blockNumber": s.BlockNumber,
		"timestamp":   s.Timestamp,
		"date":        s.Date(),
	}

	order = []string{
		"blockNumber",
		"timestamp",
		"date",
	}

	if format == "json" {
		if len(s.Name) > 0 {
			model["name"] = s.Name
		}
		if verbose {
			if len(s.Component) > 0 {
				model["component"] = s.Component
			}
			if len(s.Description) > 0 {
				model["description"] = s.Description
			}
		}

	} else {
		model["name"] = s.Name
		order = append(order, "name")

		if verbose {
			model["component"] = s.Component
			order = append(order, "component")

			model["description"] = s.Description
			order = append(order, "description")
		}
	}

	// EXISTING_CODE

	return Model{
		Data:  model,
		Order: order,
	}
}

func (s *NamedBlock) Date() string {
	return base.FormattedDate(s.Timestamp)
}

// FinishUnmarshal is used by the cache. It may be unused depending on auto-code-gen
func (s *NamedBlock) FinishUnmarshal(fileVersion uint64) {
	_ = fileVersion
	// EXISTING_CODE
	// EXISTING_CODE
}

// EXISTING_CODE
// EXISTING_CODE
