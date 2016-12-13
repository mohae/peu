package peu

import (
	"io"

	"github.com/mohae/magicnum/compress"
	"github.com/pierrec/lz4"
)

// Compress is the handler for compression. Bytes read is returned along with
// any non io.EOF error that may have occurred.
func Compress(format string, r io.Reader, w io.Writer) (int64, error) {
	if format == "" {
		return 0, ErrNoFormat
	}
	f := compress.ParseFormat(format)
	switch f {
	case compress.LZ4:
		return CompressLZ4(r, w)
	default:
		return 0, ErrUnsupported
	}
}

// CompressLZ4 compresses using lz4 compression. Bytes read is returned along
// with any non io.EOF error that may have occurred.
func CompressLZ4(r io.Reader, w io.Writer) (int64, error) {
	// create the lz4 writer
	lzw := lz4.NewWriter(w)
	n, err := io.Copy(lzw, r)
	if err != nil {
		// errors get counted and aggregated
		return n, err
	}
	return n, nil
}
