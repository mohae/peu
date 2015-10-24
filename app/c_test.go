package app

import (
	"testing"

	"github.com/mohae/contour"
	magicnum "github.com/mohae/magicnum/mcompress"
)

func TestCOutFile(t *testing.T) {
	tests := []struct {
		path  string
		ext   magicnum.Format
		odir  string
		ofile string
		err   string
	}{
		{"", magicnum.Unknown, "", "", "unable to create compression output filename: no filename received"},
		{"", magicnum.Unknown, "output", "", "unable to create compression output filename: no filename received"},
		{"file", magicnum.Unknown, "", "", "unable to create compression output filename: unknown output format"},
		{"file", magicnum.Zip, "", "file.zip", ""},
		{"file", magicnum.Zip, "output", "output/file.zip", ""},
		{"file.txt", magicnum.Unknown, "", "file.txt", "unable to create compression output filename: unknown output format"},
		{"file.txt", magicnum.Zip, "", "file.txt.zip", ""},
		{"file.txt", magicnum.Zip, "output", "output/file.txt.zip", ""},
		{"path/to/file", magicnum.Unknown, "", "path/to/file", "unable to create compression output filename: unknown output format"},
		{"path/to/file", magicnum.Zip, "", "path/to/file.zip", ""},
		{"path/to/file", magicnum.Zip, "output", "output/file.zip", ""},
		{"path/to/file.txt", magicnum.Unknown, "", "path/to/file.txt", "unable to create compression output filename: unknown output format"},
		{"path/to/file.txt", magicnum.Zip, "", "path/to/file.txt.zip", ""},
		{"path/to/file.txt", magicnum.Zip, "output", "output/file.txt.zip", ""},
	}
	for i, test := range tests {
		contour.UpdateString(OutputDir, test.odir)
		fname, err := cOutFile(test.path, test.ext)
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
