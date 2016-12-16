package main

import "testing"

func TestStripExt(t *testing.T) {
	tests := []struct {
		path     string
		expected string
	}{
		{"", ""},
		{"example", "example"},
		{"example.txt", "example"},
		{"example.txt.bz2", "example.txt"},
		{"path/example", "path/example"},
		{"path/example.txt", "path/example"},
		{"path/example.txt.gz", "path/example.txt"},
		{"path/to/example", "path/to/example"},
		{"path/to/example.txt", "path/to/example"},
		{"path/to/example.txt.gz", "path/to/example.txt"},
		{"/path/to/example", "/path/to/example"},
		{"/path/to/example.txt", "/path/to/example"},
		{"/path/to/example.txt.gz", "/path/to/example.txt"},
	}
	for i, test := range tests {
		v := stripExt(test.path)
		if v != test.expected {
			t.Errorf("%d: got %q; want %q", i, v, test.expected)
		}
	}
}
