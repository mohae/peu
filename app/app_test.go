package app

import (
	"testing"
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
