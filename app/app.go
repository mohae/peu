package app

import (
	"io/ioutil"
	"log"
	"os"

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
