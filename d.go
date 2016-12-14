package peu

import (
	"io"

	"github.com/mohae/magicnum/compress"
	"github.com/pierrec/lz4"
)

// Decompress is the handler for decompression. The decompression format to use
// is determined by the magic bytes within the data. If no match is found, an
// ErrUnsupported is returned. Bytes read is returned along with any non io.EOF
// error that may have occurred. If the reader doesn't implement io.ReaderAt
// a panic will occur.
func Decompress(r io.Reader, w io.Writer) (int64, error) {
	a, ok := r.(io.ReaderAt)
	if !ok {
		panic("io.Reader does not implement io.ReaderAt")
	}
	f, err := compress.GetFormat(a)
	if err != nil {
		return 0, err
	}
	switch f {
	case compress.LZ4:
		return DecompressLZ4(r, w)
	default:
		return 0, ErrUnsupported
	}
}

// DecompressLZ4 decompresses data compressed with lz4 compression. Bytes read
// is returned along with any non io.EOF error that may have occurred.
func DecompressLZ4(r io.Reader, w io.Writer) (int64, error) {
	// create the lz4 writer
	d := lz4.NewReader(r)
	n, err := io.Copy(w, d)
	if err != nil {
		return n, err
	}
	return n, nil
}
