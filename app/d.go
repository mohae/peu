package app

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	contour "github.com/mohae/contour"
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
	format, err := getFileFormat(src)
	if err != nil {
		return fmt.Errorf("%s: %s", fname, err)
	}
	dFname, err := dOutFile(fname)
	// open output file
	dst, err := os.OpenFile(dFname, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return err
	}
	switch format {
	case LZ4:
		// extract
		err = dlz4(dst, src)
		return err
	}
	return nil
}

// dOutFile creates the output filename for decompression. The output filename
// is created by stripping the extension from it. If the received filename is
// an empty string, an error will be returned.
//
// If an output directory was specified, this will replace the file's directory,
// if it has one; otherwise it will become the file's directory.
func dOutFile(fname string) (string, error) {
	if fname == "" {
		return "", errors.New("unable to create decompression output filename: no filename received")
	}
	dir, name := fileParts(fname)
	// see if there is an output dir specified; if so overrid e the current dir info
	odir := contour.GetString(OutputDir)
	if odir != "" {
		dir = odir
	}
	// return the finalized name
	return filepath.Join(filepath.FromSlash(dir), name), nil
}

func dlz4(dst io.Writer, src io.Reader) error {
	lzr := lz4.NewReader(src)
	_, err := io.Copy(dst, lzr)
	if err != nil {
		return err
	}
	return nil
}
