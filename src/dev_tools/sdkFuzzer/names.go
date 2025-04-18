// Copyright 2016, 2024 The TrueBlocks Authors. All rights reserved.
// Use of this source code is governed by a license that can
// be found in the LICENSE file.
/*
 * Parts of this file were auto generated. Edit only those parts of
 * the code inside of 'EXISTING_CODE' tags.
 */
package main

// EXISTING_CODE
import (
	"fmt"
	"strings"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/file"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
	sdk "github.com/TrueBlocks/trueblocks-sdk/v5"
)

// EXISTING_CODE

// DoNames tests the Names sdk function
func DoNames() {
	file.EstablishFolder("sdkFuzzer-output/names")
	opts := sdk.NamesOptions{}
	ShowHeader("DoNames", opts)

	globs := noCache(noEther(globals))
	expand := []bool{false, true}
	matchCase := []bool{false, true}
	all := []bool{false, true}
	custom := []bool{false, true}
	prefund := []bool{false, true}
	regular := []bool{false, true}
	dryRun := []bool{false, true}
	// Fuzz Loop
	// EXISTING_CODE
	_ = dryRun
	opts = sdk.NamesOptions{
		Terms: []string{"0xf"},
	}
	// names,command,default|
	for _, e := range expand {
		for _, m := range matchCase {
			for _, a := range all {
				for _, c := range custom {
					for _, p := range prefund {
						for _, r := range regular {
							opts := sdk.NamesOptions{
								Terms:     []string{"0xf"},
								Expand:    e,
								MatchCase: m,
								All:       a,
								Custom:    c,
								Prefund:   p,
								Regular:   r,
							}
							for _, g := range globs {
								opts.Globals = g
								fn := getFilenameNames("", &opts)
								TestNames("names", "", fn, &opts)
								fn = getFilenameNames("addr", &opts)
								TestNames("addr", "", fn, &opts)
								opts.Terms = []string{"0"}
								fn = getFilenameNames("tags", &opts)
								TestNames("tags", "", fn, &opts)
							}
						}
					}
				}
			}
		}
	}

	ShowHeader("DoNames-other", &opts)
	fn := getFilenameNames("autoname", &opts)
	TestNames("autoname", "0xde30da39c46104798bb5aa3fe8b9e0e1f348163f", fn, &opts)

	// DryRun    bool     `json:"dryRun,omitempty"`
	// func (opts *NamesOptions) NamesClean() ([]types.Message, *types.MetaData, error) {
	// func (opts *NamesOptions) NamesCreate() ([]types.Name, *types.MetaData, error) {
	// func (opts *NamesOptions) NamesUpdate() ([]types.Name, *types.MetaData, error) {
	// func (opts *NamesOptions) NamesDelete() ([]types.Name, *types.MetaData, error) {
	// func (opts *NamesOptions) NamesUndelete() ([]types.Name, *types.MetaData, error) {
	// func (opts *NamesOptions) NamesRemove() ([]types.Name, *types.MetaData, error) {
	// EXISTING_CODE
	Wait()
}

func TestNames(which, value, fn string, opts *sdk.NamesOptions) {
	fn = strings.Replace(fn, ".json", "-"+which+".json", 1)
	// EXISTING_CODE
	// EXISTING_CODE

	switch which {
	case "names":
		if names, _, err := opts.Names(); err != nil {
			ReportError(fn, opts, err)
		} else {
			if err := SaveToFile(fn, names); err != nil {
				ReportError2(fn, err)
			} else {
				ReportOkay(fn)
			}
		}
	case "addr":
		if addr, _, err := opts.NamesAddr(); err != nil {
			ReportError(fn, opts, err)
		} else {
			if err := SaveToFile(fn, addr); err != nil {
				ReportError2(fn, err)
			} else {
				ReportOkay(fn)
			}
		}
	case "tags":
		if tags, _, err := opts.NamesTags(); err != nil {
			ReportError(fn, opts, err)
		} else {
			if err := SaveToFile(fn, tags); err != nil {
				ReportError2(fn, err)
			} else {
				ReportOkay(fn)
			}
		}
	case "clean":
		if clean, _, err := opts.NamesClean(); err != nil {
			ReportError(fn, opts, err)
		} else {
			if err := SaveToFile(fn, clean); err != nil {
				ReportError2(fn, err)
			} else {
				ReportOkay(fn)
			}
		}
	case "autoname":
		if autoname, _, err := opts.NamesAutoname(base.HexToAddress(value)); err != nil {
			ReportError(fn, opts, err)
		} else {
			if err := SaveToFile(fn, autoname); err != nil {
				ReportError2(fn, err)
			} else {
				ReportOkay(fn)
			}
		}
	case "create":
		if create, _, err := opts.NamesCreate(); err != nil {
			ReportError(fn, opts, err)
		} else {
			if err := SaveToFile(fn, create); err != nil {
				ReportError2(fn, err)
			} else {
				ReportOkay(fn)
			}
		}
	case "update":
		if update, _, err := opts.NamesUpdate(); err != nil {
			ReportError(fn, opts, err)
		} else {
			if err := SaveToFile(fn, update); err != nil {
				ReportError2(fn, err)
			} else {
				ReportOkay(fn)
			}
		}
	case "delete":
		if delete, _, err := opts.NamesDelete(); err != nil {
			ReportError(fn, opts, err)
		} else {
			if err := SaveToFile(fn, delete); err != nil {
				ReportError2(fn, err)
			} else {
				ReportOkay(fn)
			}
		}
	case "undelete":
		if undelete, _, err := opts.NamesUndelete(); err != nil {
			ReportError(fn, opts, err)
		} else {
			if err := SaveToFile(fn, undelete); err != nil {
				ReportError2(fn, err)
			} else {
				ReportOkay(fn)
			}
		}
	case "remove":
		if remove, _, err := opts.NamesRemove(); err != nil {
			ReportError(fn, opts, err)
		} else {
			if err := SaveToFile(fn, remove); err != nil {
				ReportError2(fn, err)
			} else {
				ReportOkay(fn)
			}
		}
	default:
		ReportError(fn, opts, fmt.Errorf("unknown which: %s", which))
		logger.Fatal("Quitting...")
		return
	}
}

// EXISTING_CODE
func getFilenameNames(pre string, opts *sdk.NamesOptions) string {
	baseFn := "names/names"
	if len(pre) > 0 {
		baseFn += "-" + pre
	}
	if opts.Expand {
		baseFn += "-expand"
	}
	if opts.MatchCase {
		baseFn += "-match"
	}
	if opts.All {
		baseFn += "-all"
	}
	if opts.Custom {
		baseFn += "-custom"
	}
	if opts.Prefund {
		baseFn += "-prefund"
	}
	if opts.Regular {
		baseFn += "-regular"
	}
	return getFilename(baseFn, &opts.Globals)
}

// EXISTING_CODE
