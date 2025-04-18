// Copyright 2016, 2024 The TrueBlocks Authors. All rights reserved.
// Use of this source code is governed by a license that can
// be found in the LICENSE file.
/*
 * Parts of this file were auto generated. Edit only those parts of
 * the code inside of 'EXISTING_CODE' tags.
 */

package monitorsPkg

import (
	// EXISTING_CODE
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/internal/globals"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/caps"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/rpc"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/validate"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/walk"
	// EXISTING_CODE
)

// MonitorsOptions provides all command options for the chifra monitors command.
type MonitorsOptions struct {
	Addrs    []string              `json:"addrs,omitempty"`    // One or more addresses (0x...) to process
	Delete   bool                  `json:"delete,omitempty"`   // Delete a monitor, but do not remove it
	Undelete bool                  `json:"undelete,omitempty"` // Undelete a previously deleted monitor
	Remove   bool                  `json:"remove,omitempty"`   // Remove a previously deleted monitor
	Clean    bool                  `json:"clean,omitempty"`    // Clean (i.e. remove duplicate appearances) from monitors, optionally clear stage
	List     bool                  `json:"list,omitempty"`     // List monitors in the cache (--verbose for more detail)
	Count    bool                  `json:"count,omitempty"`    // Show the number of active monitors (included deleted but not removed monitors)
	Staged   bool                  `json:"staged,omitempty"`   // For --clean, --list, and --count options only, include staged monitors
	Globals  globals.GlobalOptions `json:"globals,omitempty"`  // The global options
	Conn     *rpc.Connection       `json:"conn,omitempty"`     // The connection to the RPC server
	BadFlag  error                 `json:"badFlag,omitempty"`  // An error flag if needed
	// EXISTING_CODE
	// EXISTING_CODE
}

var defaultMonitorsOptions = MonitorsOptions{}

// testLog is used only during testing to export the options for this test case.
func (opts *MonitorsOptions) testLog() {
	logger.TestLog(len(opts.Addrs) > 0, "Addrs: ", opts.Addrs)
	logger.TestLog(opts.Delete, "Delete: ", opts.Delete)
	logger.TestLog(opts.Undelete, "Undelete: ", opts.Undelete)
	logger.TestLog(opts.Remove, "Remove: ", opts.Remove)
	logger.TestLog(opts.Clean, "Clean: ", opts.Clean)
	logger.TestLog(opts.List, "List: ", opts.List)
	logger.TestLog(opts.Count, "Count: ", opts.Count)
	logger.TestLog(opts.Staged, "Staged: ", opts.Staged)
	opts.Conn.TestLog()
	opts.Globals.TestLog()
}

// String implements the Stringer interface
func (opts *MonitorsOptions) String() string {
	b, _ := json.MarshalIndent(opts, "", "  ")
	return string(b)
}

// monitorsFinishParseApi finishes the parsing for server invocations. Returns a new MonitorsOptions.
func monitorsFinishParseApi(w http.ResponseWriter, r *http.Request) *MonitorsOptions {
	values := r.URL.Query()
	if r.Header.Get("User-Agent") == "testRunner" {
		values.Set("testRunner", "true")
	}
	return MonitorsFinishParseInternal(w, values)
}

func MonitorsFinishParseInternal(w io.Writer, values url.Values) *MonitorsOptions {
	copy := defaultMonitorsOptions
	copy.Globals.Caps = getCaps()
	opts := &copy
	for key, value := range values {
		switch key {
		case "addrs":
			for _, val := range value {
				s := strings.Split(val, " ") // may contain space separated items
				opts.Addrs = append(opts.Addrs, s...)
			}
		case "delete":
			opts.Delete = true
		case "undelete":
			opts.Undelete = true
		case "remove":
			opts.Remove = true
		case "clean":
			opts.Clean = true
		case "list":
			opts.List = true
		case "count":
			opts.Count = true
		case "staged":
			opts.Staged = true
		default:
			if !copy.Globals.Caps.HasKey(key) {
				err := validate.Usage("Invalid key ({0}) in {1} route.", key, "monitors")
				if opts.BadFlag == nil || opts.BadFlag.Error() > err.Error() {
					opts.BadFlag = err
				}
			}
		}
	}
	opts.Conn = opts.Globals.FinishParseApi(w, values, opts.getCaches())

	// EXISTING_CODE
	// EXISTING_CODE
	opts.Addrs, _ = opts.Conn.GetEnsAddresses(opts.Addrs)

	return opts
}

// monitorsFinishParse finishes the parsing for command line invocations. Returns a new MonitorsOptions.
func monitorsFinishParse(args []string) *MonitorsOptions {
	// remove duplicates from args if any (not needed in api mode because the server does it).
	dedup := map[string]int{}
	if len(args) > 0 {
		tmp := []string{}
		for _, arg := range args {
			if cnt := dedup[arg]; cnt == 0 {
				tmp = append(tmp, arg)
			}
			dedup[arg]++
		}
		args = tmp
	}

	defFmt := "txt"
	opts := GetOptions()
	opts.Conn = opts.Globals.FinishParse(args, opts.getCaches())

	// EXISTING_CODE
	for _, arg := range args {
		if base.IsValidAddress(arg) {
			opts.Addrs = append(opts.Addrs, arg)
		}
	}
	// EXISTING_CODE
	opts.Addrs, _ = opts.Conn.GetEnsAddresses(opts.Addrs)
	if len(opts.Globals.Format) == 0 || opts.Globals.Format == "none" {
		opts.Globals.Format = defFmt
	}

	return opts
}

func GetOptions() *MonitorsOptions {
	// EXISTING_CODE
	// EXISTING_CODE
	return &defaultMonitorsOptions
}

func getCaps() caps.Capability {
	var capabilities caps.Capability // capabilities for chifra monitors
	capabilities = capabilities.Add(caps.Default)
	capabilities = capabilities.Add(caps.Caching)
	capabilities = capabilities.Add(caps.Names)
	// EXISTING_CODE
	// EXISTING_CODE
	return capabilities
}

func ResetOptions(testMode bool) {
	// We want to keep writer between command file calls
	w := GetOptions().Globals.Writer
	opts := MonitorsOptions{}
	globals.SetDefaults(&opts.Globals)
	opts.Globals.TestMode = testMode
	opts.Globals.Writer = w
	opts.Globals.Caps = getCaps()
	defaultMonitorsOptions = opts
}

func (opts *MonitorsOptions) getCaches() (caches map[walk.CacheType]bool) {
	// EXISTING_CODE
	// EXISTING_CODE
	return
}

// EXISTING_CODE
// EXISTING_CODE
