package peu

import (
	"bytes"
	"testing"

	"github.com/mohae/magicnum/compress"
)

var ipsum = []byte(`Lorem ipsum dolor`)

func TestCompress(t *testing.T) {
	tests := []struct {
		format string
		data   []byte
		cdata  []byte
		n      int64
		err    error
	}{
		{"", []byte{}, []byte{}, 0, ErrNoFormat},
		{"zz", []byte{}, []byte{}, 0, ErrUnsupported},
		{"tar", []byte{}, []byte{}, 0, ErrUnsupported},
		{
			"lz4", ipsum,
			[]byte{
				0x04, 0x22, 0x4d, 0x18, 0x64, 0x70, 0xb9, 0x11, 0x00, 0x00,
				0x80, 0x4c, 0x6f, 0x72, 0x65, 0x6d, 0x20, 0x69, 0x70, 0x73,
				0x75, 0x6d, 0x20, 0x64, 0x6f, 0x6c, 0x6f, 0x72,
			},
			17, nil,
		},
		{
			"gzip", ipsum,
			[]byte{
				0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff,
				0xf2, 0xc9, 0x2f, 0x4a, 0xcd, 0x55, 0xc8, 0x2c, 0x28, 0x2e,
				0xcd, 0x55, 0x48, 0xc9, 0xcf, 0xc9, 0x2f, 0x02, 0x04, 0x00,
				0x00, 0xff, 0xff, 0x32, 0xfb, 0x87, 0x4e, 0x11, 0x00, 0x00,
				0x00,
			},
			17, nil,
		},
		{
			"bzip2", ipsum,
			nil, 17, ErrUnsupported,
		},
	}
	for _, test := range tests {
		r := bytes.NewReader(test.data)
		var w bytes.Buffer
		n, err := Compress(test.format, r, &w)
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
		if bytes.Compare(w.Bytes(), test.cdata) != 0 {
			t.Errorf("%q: got %x; want %x", test.format, w.Bytes(), test.cdata)
		}
	}
}

func TestCompressionIsSupported(t *testing.T) {
	tests := []struct {
		compress.Format
		t bool
	}{
		{compress.Unknown, false},
		{compress.BZip2, false},
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
		b := CompressionIsSupported(test.Format)
		if b != test.t {
			t.Errorf("%s: got %t want %t", test.Format, b, test.t)
		}
	}
}
