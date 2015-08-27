// Copyright © 2014, All rights reserved
// Joel Scoble, https://github.com/mohae/peu
//
// This is licensed under The MIT License. Please refer to the included
// LICENSE file for more information. If the LICENSE file has not been
// included, please refer to the url above.
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License
//
package main

import (
	"fmt"
	"os"

	"github.com/mohae/cli"
	"github.com/mohae/peu/app"
)

// This is modeled on mitchellh's realmain wrapper
func main() {
	os.Exit(realMain())
}

// realMain, is the actual main for the application. This keeps all changes
// needed for a new application to one file in the main application directory.
// In addition to this, only commands/ needs to be modified, adding the app's
// commands and any handler codes for those commands, like the 'cmd' package.
//
// realMain allows for defers to be executed before the os.Exit(), should there
// be any.
//
// No logging is done until the flags are processed, since the flags could
// enable/disable output, alter it, or alter its output locations. Everything
// must go to stdout until then.
func realMain() int {
	// Get the command line args.
	args := os.Args[1:]
	// Setup the args, Commands, and Help info.
	cli := &cli.CLI{
		Name:     app.Name,
		Version:  Version,
		Commands: Commands,
		Args:     args,
		HelpFunc: cli.BasicHelpFunc(app.Name),
	}
	// Run the passed command, recieve back a message and error object.
	exitCode, err := cli.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error executing CLI: %s\n", err.Error())
		return 1
	}
	// Return the exitcode.
	return exitCode
}
