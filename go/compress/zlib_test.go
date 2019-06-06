package compress_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/qinhan-shu/study/go/compress"
)

func TestCompressWithZlib(t *testing.T) {
	src := []byte("sfdsfdsfdsfdsfdsfdsfdsfdsfdsfdeCompressedBytes, err := compress.DecompressWithZlib(compressedBytes)sfdsfdsfdsfdsfdsfdsfdsfdsfdsfdsfd")
	compressedBytes, err := compress.CompressWithZlib(src)
	if err != nil {
		t.Error("compress with zlib error", err)
		return
	}
	deCompressedBytes, err := compress.DecompressWithZlib(compressedBytes)
	if err != nil {
		t.Error("decompress with zlib error", err)
		return
	}
	// length
	t.Log("origin length = ", len(src))
	t.Log("compressed length = ", len(compressedBytes))
	t.Log("decompressed length = ", len(deCompressedBytes))
	// detail bytes
	t.Log("origin bytes = ", string(src))
	t.Log("decompressed bytes = ", string(deCompressedBytes))
	// compare
	if !assert.Equal(t, deCompressedBytes, src) {
		t.Error("decompress bytes is not equal to origin bytes")
		return
	}
}
