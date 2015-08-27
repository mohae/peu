package app

import (
	"io/ioutil"
	"log"
	"os"

	contour "github.com/mohae/contour"
)

var (
	// Name is the name of the application
	Name = "peu"

	// CfgFile is the name of the configuration file for the application.
	CfgFilename = "app.json"
)

// Variables for configuration entries, or just hard code them.
var (
	CfgFile     = "cfgfile"     // configuration filename; format type is inferred from ext.
	Log         = "log"         // to log or not to log
	LogFile     = "logfile"     // output filename for log output, stderr if empty
	Verbose     = "Verbose"     // Verbose output bool
	VerboseFile = "VerboseFile" // output filename for Verbose output; stdout if empty.
)

// Application config.
var Cfg = contour.AppCfg()
var logF *os.File

// set-up the application defaults and let contour know about the app.
func init() {
	// Register the cfg file. Contour determines the format that the cfg file is
	// in by looking at its extension.
	contour.RegisterCfgFile(CfgFile, CfgFilename)
	contour.RegisterStringCore("name", Name)
	// Logging and output related
	contour.RegisterBoolFlag(Log, "l", false, "false", "enable/disable logging")
	contour.RegisterStringFlag(LogFile, "", "", "", "logfile, will use stderr if not set")
	contour.RegisterBoolFlag(Verbose, "v", false, "false", "enable/disable verbose output")
	initApp()
	// Now that the configuration is set, set app logging. May be overridden later.
	SetAppLogging()
}

// InitApp is the best place to add custom defaults for your application,
func initApp() {
	contour.RegisterBoolFlag("lower", "", false, "false", "lowercase output")
}

// SetCfg initialized the application's configuration. When the config is
// has been initialized, the preset-enivronment variables, application
// defaults, and your application's configuration file, should it have one,
// will all be merged according to the setting properties.
//
// After this, only overrides can occur via command flags.
func SetCfg() error {
	// Set config:
	err := contour.SetCfg()
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}

// SetAppLog sets the logger for package loggers and allow for custom-
// ization of the applications log. This is where app specific code for
// setting up the application's log should be.
//
// SetAppLog assumes that log is enabled if it has been called as its
// caller should be SetLog(). If you are going to call this from elsewhere,
// first make sure that log is enabled.
func SetAppLogging() {
	if !contour.GetBool(Log) {
		log.SetOutput(ioutil.Discard)
		return
	}
	// get the logfilename, if it's not set, use stderr
	if contour.GetString(LogFile) != "" {
		var err error
		logF, err = os.OpenFile(contour.GetString(LogFile), os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
		if err != nil {
			panic(err)
		}
		log.SetOutput(logF)
	} else {
		log.SetOutput(os.Stderr)
	}
	return
}

func CloseLog() error {
	return logF.Close()
}
