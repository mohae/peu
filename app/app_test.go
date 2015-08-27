package app

import (
	"testing"

	"github.com/mohae/contour"
)

func TestFileParts(t *testing.T) {
	tests := []struct {
		path  string
		eDir  string
		eFile string
	}{
		{"", "", ""},
		{"test", "", "test"},
		{"test.txt", "", "test"},
		{"testDir/test", "testDir/", "test"},
		{"testDir/test.txt", "testDir/", "test"},
		{"path/to/testDir/test", "path/to/testDir/", "test"},
		{"path/to/testDir/test.txt", "path/to/testDir/", "test"},
	}
	for i, test := range tests {
		dir, file := fileParts(test.path)
		if dir != test.eDir {
			t.Errorf("%d: expected %q got %q", i, test.eDir, dir)
		}
		if file != test.eFile {
			t.Errorf("%d: expected %q got %q", i, test.eFile, file)
		}
	}
}

func TestCOutFile(t *testing.T) {
	tests := []struct {
		path  string
		ext   string
		odir  string
		ofile string
		err   string
	}{
		{"", "", "", "", "unable to create compression output filename: no filename received"},
		{"", "ext", "", "", "unable to create compression output filename: no filename received"},
		{"", "", "output", "", "unable to create compression output filename: no filename received"},
		{"file", "", "", "", "unable to create compression output filename: no extension received"},
		{"file", "zip", "", "file.zip", ""},
		{"file", "zip", "output", "output/file.zip", ""},
		{"file.txt", "", "", "file.txt", "unable to create compression output filename: no extension received"},
		{"file.txt", "zip", "", "file.txt.zip", ""},
		{"file.txt", "zip", "output", "output/file.txt.zip", ""},
		{"path/to/file", "", "", "path/to/file", "unable to create compression output filename: no extension received"},
		{"path/to/file", "zip", "", "path/to/file.zip", ""},
		{"path/to/file", "zip", "output", "output/file.zip", ""},
		{"path/to/file.txt", "", "", "path/to/file.txt", "unable to create compression output filename: no extension received"},
		{"path/to/file.txt", "zip", "", "path/to/file.txt.zip", ""},
		{"path/to/file.txt", "zip", "output", "output/file.txt.zip", ""},
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
