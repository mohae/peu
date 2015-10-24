package app

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	contour "github.com/mohae/contour"
	magicnum "github.com/mohae/magicnum/mcompress"
	"github.com/pierrec/lz4"
)

// C is the handler for compression
func C(parms []string) (msg string, err error) {
	if len(parms) == 0 {
		return "", errors.New("compress: no compression algorithm specified")
	}
	if len(parms) == 1 {
		return "", errors.New("compress: no file specified")
	}
	switch parms[0] {
	case "lz4":
		msg, err = clz4(parms[1:])
	default:
		return "", fmt.Errorf("compress: %s is not a supported compression algorithm", parms[0])
	}
	return msg, err
}

// cOutFile creates the output filename for compression. The output filename's
// separator is normalized to the os.Separator. An error will occur if either
// the filename or the extension is an empty string.
//
// The filename will be appended with the received extension.
//
// If an output directory was specified, this will replace the file's
// directory, if it has one; otherwise it will become the file's directory.
func cOutFile(fname string, format magicnum.Format) (string, error) {
	if fname == "" {
		return "", errors.New("unable to create compression output filename: no filename received")
	}
	dir, fname := filepath.Split(fname)
	// see if there is an output dir specified; if so override the current dir info
	odir := contour.GetString(OutputDir)
	if odir != "" {
		dir = odir
	}
	// add the extension to the filename
	return filepath.Join(filepath.FromSlash(dir), fmt.Sprintf("%s%s", fname, format.Ext())), nil
}

// clz4 compresses using lz4 compression
func clz4(files []string) (string, error) {
	var errMsg string
	var errCnt int
	for _, file := range files {
		// open the file
		srcF, err := os.Open(file)
		if err != nil {
			return "", err
		}
		fname, err := cOutFile(file, magicnum.LZ4)
		if err != nil {
			srcF.Close()
			return "", err
		}
		// create the output file
		dstF, err := os.OpenFile(fname, os.O_CREATE|os.O_WRONLY, 0755)
		if err != nil {
			srcF.Close()
			return "", err
		}
		// create the lz4 writer
		lzw := lz4.NewWriter(dstF)
		_, err = io.Copy(lzw, srcF)
		if err != nil {
			// errors get counted and aggregated
			errMsg += fmt.Sprintf("\n%s", err)
			errCnt++
		}
		lzw.Close()
		srcF.Close()
		dstF.Close()
	}
	if errCnt > 0 {
		return fmt.Sprintf("%d files were processed\n%d were successfully compressed using lz4\n%d had errors\n", len(files), len(files)-errCnt, errCnt), fmt.Errorf("lz4 compression error(s): %s", errMsg)
	}
	if len(files) == 1 {
		return fmt.Sprintf("%s was successfully compressed using lz4", files[0]), nil
	}
	return fmt.Sprintf("%d files were proceseed\n^d were successfully compressed using lz4\n", len(files), len(files)-errCnt), nil

}
