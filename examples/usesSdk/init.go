package main

import (
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
)

// DoInit tests the When sdk function
func DoInit() {
	logger.Info("DoInit")

	// opts := sdk.InitOptions{
	// }

	// buf := bytes.Buffer{}
	// if err := opts.InitBytes(&buf); err != nil {
	// 	logger.Error(err)
	// }

	// file.StringToAsciiFile("usesSDK/init.json", buf.String())
	// fmt.Println(buf.String())
}

// func (opts *InitOptions) InitAll() ([]bool, *types.MetaData, error) {