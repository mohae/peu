package app

import (
	"testing"
)

func TestDOutFile(t *testing.T) {
	tests := []struct {
		path  string
		odir  string
		ofile string
		err   string
	}{
		{"", "", "", "unable to create decompression output filename: no filename received"},
		{"file", "", "file", ""},
		{"file.txt", "", "file", ""},
		{"file.txt.lz4", "", "file.txt", ""},
		{"path/to/file.txt.lz4", "", "path/to/file.txt", ""},
		{"file", "out/dir", "out/dir/file", ""},
		{"file.txt", "out/dir", "out/dir/file", ""},
		{"file.txt.lz4", "out/dir", "out/dir/file.txt", ""},
		{"path/to/file.txt.lz4", "out/dir", "out/dir/file.txt", ""},
	}
	for i, test := range tests {
		contour.UpdateString(OutputDir, test.odir)
		fname, err := dOutFile(test.path)
		if err != nil {
			if err.Error() != test.err {
				t.Errorf("Expected %q, got %q", test.err, err)
			}
			continue
		}
		if fname != test.ofile {
			t.Errorf("%d: expected %q got %q", i, test.ofile, fname)
		}

	}
}
