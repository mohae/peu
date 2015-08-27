package command

import (
	"strings"

	"github.com/mohae/cli"
	"github.com/mohae/contour"
	"github.com/mohae/peu/app"
)

// DCommand is a Command implementation for decompression
type DCommand struct {
	UI cli.Ui
}

// Help prints the help text for the run sub-command.
func (d *DCommand) Help() string {
	helpText := `
Usage: peu d [flags] <filelist string...>

d will take a 1 or more files and decompress them. If the -o flag is not
passed, the files will be decompressed in place. The algorithm that is 
used for decompression is determined by matching magic numbers. If a match
is not found, the algorithm is not supported and an error will be returned.

    $ peu d example.txt

    $ peu d -o=out/path example.txt
    

peu supports 1 flag:

    -o=output/path    save the output of the command to the path specified.
`
	return strings.TrimSpace(helpText)
}

// Run runs the d command; the args are a variadic list of filenames to process.
func (c *DCommand) Run(args []string) int {
	// set up the command flags
	contour.SetUsage(func() { c.UI.Output(c.Help()) })
	// Filter the flags from the args and update the config with them.
	// The args remaining after being filtered are returned.
	filteredArgs, err := contour.FilterArgs(args)
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}
	// Run the command in the package.
	message, err := app.D(filteredArgs)
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}
	c.UI.Output(message)
	return 0
}

// Synopsis provides a precis of the hello command.
func (c *DCommand) Synopsis() string {
	ret := `Compresses files using the specified algorithm
`
	return ret
}
