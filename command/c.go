package command

import (
	"strings"

	"github.com/mohae/cli"
	"github.com/mohae/contour"
	"github.com/mohae/peu/app"
)

// CCommand is a Command implementation for compression
type CCommand struct {
	UI cli.Ui
}

// Help prints the help text for the run sub-command.
func (c *CCommand) Help() string {
	helpText := `
Usage: peu c [flags] <algorithm> <filelist string...>

c will take a 1 or more files and compress them using the specified
algorithm. It the algorithm is not supported, an error message will
be printed. 

    $ peu c lz4 example.txt

    $ peu c -o=out/path lz4 example.txt
    

peu supports 1 flag:

    -o=output/path    save the output of the command to the path specified.
`
	return strings.TrimSpace(helpText)
}

// Run runs the c command; the args are a variadic list of filenames to process.
func (c *CCommand) Run(args []string) int {
	// set up the command flags
	contour.SetUsage(func() { c.UI.Output(c.Help()) })
	// Filter the flags from the args and update the config with them.
	// The args remaining after being filtered are returned.
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}
	// Run the command in the package.
	message, err := app.C(filteredArgs)
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}
	c.UI.Output(message)
	return 0
}

// Synopsis provides a precis of the hello command.
func (c *HelloCommand) Synopsis() string {
	ret := `Concatonates the list of words it recieved
to 'Hello', applies any formatting required,
and returns the result.
`
	return ret
}
