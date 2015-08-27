package app

import (
	"errors"
	"fmt"
	"io"
	"os"

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
		defer srcF.Close()
		fname := outFile(file, ".lz4")
		// create the output file
		dstF, err := os.OpenFile(fname, os.O_CREATE|os.O_WRONLY, 0755)
		if err != nil {
			return "", err
		}
		defer dstF.Close()
		// create the lz4 writer
		lzw := lz4.NewWriter(dstF)
		defer lzw.Close()
		_, err = io.Copy(lzw, srcF)
		if err != nil {
			// errors get counted and aggregated
			errMsg += fmt.Sprintf("\n%s", err)
			errCnt++
		}
	}
	if errCnt > 0 {
		return fmt.Sprintf("%d files were processed\n%d were successfully compressed using lz4\n%d had errors\n", len(files), len(files)-errCnt, errCnt), fmt.Errorf("lz4 compression error(s): %s", errMsg)
	}
	if len(files) == 1 {
		return fmt.Sprintf("%s was successfully compressed using lz4", files[0]), nil
	}
	return fmt.Sprintf("%d files were proceseed\n^d were successfully compressed using lz4\n", len(files), len(files)-errCnt), nil

}
