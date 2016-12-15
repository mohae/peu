// Copyright 2015-2016 Joel Scoble.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (
	"flag"
	"os"
)

var (
	worker = Worker{errs: make(map[string]error), format: "gzip"}
)

func init() {
	flag.StringVar(&worker.format, "f", worker.format, "compression format to use")
	flag.BoolVar(&worker.decompress, "d", false, "decompress the source")
}

// this allows for defers to run and exiting with a return code.
func main() {
	os.Exit(realMain())
}

func realMain() int {
	flag.Parse()
	worker.files = flag.Args() // this is the list of files
	err := worker.Work()
	if err != nil {
		if err == ErrProcess {
			return len(worker.errs) // the number of operations (files) that errored is the return code
		}
		return 1
	}

	return 0
}
