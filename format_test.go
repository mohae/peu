package peu

import (
	"bytes"
	"testing"

	"github.com/mohae/magicnum/compress"
)

func TestCompressionFormat(t *testing.T) {
	tests := []struct {
		name   string
		bytes  []byte
		offset int
		compress.Format
		err string
	}{
		{"unknown", []byte{0x00, 0x00}, 0, compress.Unknown, "unknown compression format"},
		{"unsupported", []byte{0x50, 0x4b, 0x03, 0x04}, 0, compress.Zip, "unsupported compression format"}, //zip
		{"gzip", []byte{0x1f, 0x8b}, 0, compress.GZip, ""},
		{"lz4", []byte{0x04, 0x22, 0x4d, 0x18}, 0, compress.LZ4, ""},
	}
	for i, test := range tests {
		var b []byte
		if i == 0 {
			b = make([]byte, 64)
		}
		b = make([]byte, 512)
		for i := test.offset; i < test.offset+len(test.bytes); i++ {
			b[i] = test.bytes[i]
		}
		r := bytes.NewReader(b)
		f, err := CompressionFormat(r)
		if err != nil {
			if err.Error() != test.err {
				t.Errorf("%s: got %s; want %s", test.name, err, test.err)
				continue
			}
		}
		if f != test.Format {
			t.Errorf("%s: got %s; want %s", test.name, f, test.Format)
		}
	}
}
