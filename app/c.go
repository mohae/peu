package app

import (
	"fmt"
	"log"
	"strings"

	"github.com/mohae/contour"
)

// C is the handler for compression
func C(parms ...string) (msg string, err error) {
	if len(parms) == 0 {
		return "", errors.Error("compress: no compression algorithm specified")
	}
	if len(parms) == 1 {
		return "", errors.Error("compress: no file specified")
	}
	switch case parms[0] {
	case "lz4":
		msg, err = clz4(parms[1:])
	default:
		return "", fmt.Errorf("compress: %s is not a supported compression algorithm", parms[0])
	}
	return msg, err
}

func clz4(files []string) (msg string, err error) {
	return msg, err
}