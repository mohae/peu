package app

import ()

const (
	Unsupported Format = iota // not supported
	LZ4                       // LZ4 compression
)

type Format int

func (f Format) String() string {
	switch f {
	case LZ4:
		return "lz4"
	}
	return "unsupported"
}

// Magic numbers for supported formats
var (
	headerLZ4 = []byte{0x18, 0x4d, 0x22, 0x04} // first 4 bytes
)

// getFileFormat tries to match up the data in the Reader to a supported
// magic number, if a match isn't found, UnsupportedFmt is returned
func getFileFormat(r io.ReaderAt) Format {
	h := make([]byte, 8, 8) // 8 is minimum cap of a byte slice so...
	// Reat the first 8 bytes since that's where most magic numbers are
	r.ReadAt(h, 0)
	// for the byte comparison, use only the parts we need
	if bytes.Equal(headerLZ4, h[0:4]) {
		return LZ4
	}
	return Unsupported
}
