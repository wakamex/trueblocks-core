// Copyright 2016, 2024 The TrueBlocks Authors. All rights reserved.
// Use of this source code is governed by a license that can
// be found in the LICENSE file.
/*
 * Parts of this file were auto generated. Edit only those parts of
 * the code inside of 'EXISTING_CODE' tags.
 */

// blocksPkg implements the chifra blocks command.
//
// The chifra blocks tool retrieves block data from your Ethereum node or, if previously cached, from the
// TrueBlocks cache. You may specify multiple blocks per invocation.
//
// By default, chifra blocks queries the full transactional details of the block (including receipts).
// You may optionally retrieve only the transaction hashes in the block (which is significantly faster).
// Additionally, you may also use this tool to retrieve uncle blocks at a give height.
//
// Another useful feature of chifra blocks is the ability to extract address appearances from a block.
// TrueBlocks uses a similar feature internally to build its index of appearances. This type of data
// is very insightful when studying end user behavior and chain-wide adoption analysis.
package blocksPkg
