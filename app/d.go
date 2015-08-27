package app

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/pierrec/lz4"
)

// D is the handler for decompression
func D(files []string) (msg string, err error) {
	if len(files) == 0 {
		return "", errors.New("decompress: no file specified")
	}
	// go through each file, check for compression type, and call appropriately
	for _, file := range files {

	}
	return msg, err
}
