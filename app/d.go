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
	// Process each file
	for _, file := range files {
		err = DToFile(file)
	}
	return msg, err
}

// DToFile decompresses each source to a file.
func DToFile(fname string) error {
	// open the file
	src, err := os.Open(fname)
	if err != nil {
		return err
	}
	// check its format
	format := getFileFormat(src)
	if format == Unsupported {
		return fmt.Errorf("%s's format is unsupported", fname)
	}
	dFname, err := dOutFile(fname)
	// open output file
	dst, err := os.OpenFile(dFname, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return err
	}
	// extract
	err = dlz4(dst, src)
	return nil
}

func dlz4(dst io.Writer, src io.Reader) error {
	lzr := lz4.NewReader(src)
	_, err := io.Copy(dst, lzr)
	if err != nil {
		return err
	}
	return nil
}
