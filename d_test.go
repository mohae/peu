package peu

import (
	"bytes"
	"testing"

	"github.com/mohae/magicnum/compress"
)

func TestDecompress(t *testing.T) {
	tests := []struct {
		format compress.Format
		data   []byte
		ddata  []byte
		n      int64
		err    error
	}{
		{compress.Unknown, []byte{}, []byte{}, 0, compress.ErrEmpty},
		{
			compress.LZ4,
			[]byte{
				0x04, 0x22, 0x4d, 0x18, 0x64, 0x70, 0xb9, 0x11, 0x00, 0x00,
				0x80, 0x4c, 0x6f, 0x72, 0x65, 0x6d, 0x20, 0x69, 0x70, 0x73,
				0x75, 0x6d, 0x20, 0x64, 0x6f, 0x6c, 0x6f, 0x72,
			},
			ipsum, 17, nil,
		},
		{
			compress.GZip,
			[]byte{
				0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff,
				0xf2, 0xc9, 0x2f, 0x4a, 0xcd, 0x55, 0xc8, 0x2c, 0x28, 0x2e,
				0xcd, 0x55, 0x48, 0xc9, 0xcf, 0xc9, 0x2f, 0x02, 0x04, 0x00,
				0x00, 0xff, 0xff, 0x32, 0xfb, 0x87, 0x4e, 0x11, 0x00, 0x00,
				0x00,
			},
			ipsum, 17, nil,
		},
		{
			compress.BZip2,
			[]byte{
				0x42, 0x5a, 0x68, 0x39, 0x31, 0x41, 0x59, 0x26, 0x53, 0x59,
				0x69, 0x54, 0xef, 0x59, 0x00, 0x00, 0x01, 0xd5, 0x80, 0x00,
				0x10, 0x40, 0x00, 0x00, 0x04, 0x06, 0x26, 0xda, 0x00, 0x20,
				0x00, 0x31, 0x00, 0xd0, 0x01, 0x09, 0x93, 0x18, 0xa6, 0x88,
				0x34, 0x4c, 0xf0, 0xe8, 0x08, 0xab, 0x6f, 0x6f, 0xc5, 0xdc,
				0x91, 0x4e, 0x14, 0x24, 0x1a, 0x55, 0x3b, 0xd6, 0x40,
			},
			append(ipsum, 0x0a), // append a newline; editor used adds one
			18, nil,
		},
	}

	for _, test := range tests {
		r := bytes.NewReader(test.data)
		var w bytes.Buffer
		n, err := Decompress(r, &w)
		if err != nil {
			if err != test.err {
				t.Errorf("%q: got %q; want %q", test.format, err, test.err)
			}
			continue
		}
		if test.err != nil {
			t.Errorf("%q: got no error; expected %q", test.format, test.err)
			continue
		}
		if n != test.n {
			t.Errorf("%q: read %d bytes; expected %d", test.format, n, test.n)
			continue
		}
		if bytes.Compare(w.Bytes(), test.ddata) != 0 {
			t.Errorf("%q: got %x; want %x", test.format, w.Bytes(), test.ddata)
		}
	}
}

func TestDecompressionIsSupported(t *testing.T) {
	tests := []struct {
		compress.Format
		t bool
	}{
		{compress.Unknown, false},
		{compress.BZip2, true},
		{compress.GZip, true},
		{compress.LZ4, true},
		{compress.Zip, false},
		{compress.ZipEmpty, false},
		{compress.ZipSpanned, false},
		{compress.Tar, false},
		{compress.Tar1, false},
		{compress.Tar2, false},
	}

	for _, test := range tests {
		b := DecompressionIsSupported(test.Format)
		if b != test.t {
			t.Errorf("%s: got %t want %t", test.Format, b, test.t)
		}
	}
}
