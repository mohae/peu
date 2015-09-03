package app

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	contour "github.com/mohae/contour"
)

const (
	// Name is the name of the application
	Name = "peu"
)

// Variables for configuration entries, or just hard code them.
var (
	OutputDir = "output_dir" // output directory, if it's not cwd
)

// Application config.
var Cfg = contour.AppCfg()

// set-up the application defaults and let contour know about the app.
func init() {
	contour.RegisterStringFlag(OutputDir, "", "", "", "the output directory; if it's not CWD")
}

// fileParts returns the directory and filename. If the filename has an
// extension, it is dropped. The path separators are normalized with ToSlash.
// It is the callers responsibility to ensure that path separator is compatible
// with the host OS.
func fileParts(s string) (dir string, base string) {
	s = filepath.ToSlash(s)
	dir, base = filepath.Split(s)
	ext := filepath.Ext(base)
	if ext != "" {
		base = strings.TrimSuffix(base, ext)
	}
	return dir, base
}
