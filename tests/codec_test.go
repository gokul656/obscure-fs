package tests

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/gokul656/obscure-fs/internal/codec"
	"github.com/gokul656/obscure-fs/internal/hashing"
	"github.com/gokul656/obscure-fs/internal/storage"
	"github.com/gokul656/obscure-fs/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestCodec(t *testing.T) {
	filePath := "../README.md"
	fileName := filepath.Base(filePath)
	buf, _ := storage.ReadFile(filePath)

	hash, _ := hashing.HashFile(filePath)
	metadata := &storage.Metadata{
		Name:     fileName,
		Checksum: hash,
	}

	ec := codec.ErasureCodec{}
	err := ec.Encode(metadata, buf)
	if err != nil {
		panic(err)
	}

	outfile, err := ec.Decode(metadata)
	if err != nil {
		panic(err)
	}

	hash, err = hashing.HashFile(outfile)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, utils.Shards, metadata.Shards)
	assert.Equal(t, utils.Pairty, metadata.Pairty)
	assert.Equal(t, metadata.Checksum, hash)

	os.RemoveAll(filepath.Dir(outfile))
}
