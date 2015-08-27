package app

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
)

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
func getFileFormat(r io.ReaderAt) (Format, error) {
	h := make([]byte, 8, 8) // 8 is minimum cap of a byte slice so...
	// Reat the first 8 bytes since that's where most magic numbers are
	r.ReadAt(h, 0)
	var h32 uint32
	// check for lz4
	hbuf := bytes.NewReader(headerLZ4)
	err := binary.Read(hbuf, binary.LittleEndian, &h32)
	if err != nil {
		return Unsupported, fmt.Errorf("error while checking if input matched LZ4's magic number: %s", err)
	}
	var m32 uint32
	mbuf := bytes.NewBuffer(headerLZ4)
	err = binary.Read(mbuf, binary.LittleEndian, &m32)
	if err != nil {
		return Unsupported, fmt.Errorf("error while converting LZ4 magic number for comparison: %s", err)
	}
	if h32 == m32 {
		return LZ4, nil
	}
	return Unsupported, errors.New("unsupported format: input format is not known")
}
