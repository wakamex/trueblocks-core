// Package index provides tools needed to acquire, read, write and test for set inclusion in an index chunk.
//
// An index chunk is a data structure with three parts. A FileRange which indicates the first block
// and last block of the chunk (inclusive), the Index which carries the list of address appearances
// in the given block range, and a Bloom which allows for rapid queries to determine if a given address
// appears in the Index without having to read the data from disc.
//
// The bloom filter returns true or false indicating either that the address MAY appear in the index or
// that it definitely does not. (In other words, there are false positives but no false negatives.)
//
// We do not read the actual data into memory, choosing instead to Seek the data directly from disc. Experimentation
// teaches us that this is faster given due to the nature of the data.

package index

import (
	"io"
	"os"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/config"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/ranges"
	shell "github.com/ipfs/go-ipfs-api"
)

// The Chunk data structure consists of three parts. A FileRange, a Index structure, and a Bloom that
// carries set membership information for the Index.
type Chunk struct {
	Range ranges.FileRange
	Index Index
	Bloom Bloom
}

// versions returns the version strings found in both the bloom and the index file.
func versions(fileName string) (string, string, error) {
	chunk, err := OpenChunk(fileName, false)
	if err != nil {
		return "", "", err
	}
	defer chunk.Close()
	indexVersion := config.VersionTags[chunk.Index.Header.Hash.Hex()]
	bloomVersion := config.VersionTags[chunk.Bloom.Header.Hash.Hex()]
	return bloomVersion, indexVersion, nil
}

// OpenChunk returns a fully initialized index chunk. The path argument may point to either a bloom filter file or the
// index data file. Either will work. The bloom filter file must exist and will be opened for reading and its header
// will be read into memory, but the filter itself is not. The index data file need not exist (it will be downloaded
// later if the bloom indicates that its needed). If the index file does exist, however, it will be opened for reading
// and its header will be read into memory, but the index data itself will not be.
func OpenChunk(path string, check bool) (chunk Chunk, err error) {
	chunk.Range, err = ranges.RangeFromFilenameE(path)
	if err != nil {
		return
	}

	chunk.Bloom, err = OpenBloom(ToBloomPath(path), check /* check */)
	if err != nil {
		return
	}

	chunk.Index, err = OpenIndex(ToIndexPath(path), check /* check */)
	return
}

// Close closes both the bloom filter file pointer and the index data file pointer (if they are open)
func (chunk *Chunk) Close() {
	if chunk.Bloom.File != nil {
		chunk.Bloom.File.Close()
		chunk.Bloom.File = nil
	}

	if chunk.Index.File != nil {
		chunk.Index.File.Close()
		chunk.Index.File = nil
	}
}

// ChunkCid returns IPFS CID for the chunk without uploading it
func ChunkCid(path string) (chunkCid string, err error) {
	file, err := os.OpenFile(path, os.O_RDONLY, 0)
	if err != nil {
		return
	}
	defer file.Close()
	return calculateCid(file)
}

func calculateCid(r io.Reader) (chunkCid string, err error) {
	sh := shell.NewShell(config.GetPinning().LocalPinUrl)
	return sh.AddNoPin(r)
}
