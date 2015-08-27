package app

import (
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

// outFile creates the output filename. The output filename's separator is
// normalized to the os.Separator. Empty filenames return "", empty string.
//
// If the filename has an extension, it will be replaced with the passed
// extension. If it does not have an extension, it will be appended with the
// passed extension.
//
// If an output directory was specified, this will replace the file's
// directory, if it has one.
func outFile(s, ext string) string {
	if s == "" {
		return ""
	}
	dir, fname := fileParts(s)
	// see if there is an output dir specified; if so override the current dir info
	odir := contour.GetString(OutputDir)
	if odir != "" {
		dir = odir
	}
	// normalize ext to .ext
	if !strings.HasPrefix(ext, ".") && ext != "" {
		ext = fmt.Sprintf(".%s", ext)
	}
	// add the extension to the filename
	return filepath.Join(filepath.FromSlash(dir), fmt.Sprintf("%s%s", fname, ext))
}
