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
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

var (
	worker Worker
	name   = filepath.Base(os.Args[0])
	help   bool
)

func init() {
	worker = Worker{
		concurrency: runtime.NumCPU(), // default to number of cpus
		errs:        make(map[string]error),
		format:      "gzip",
	}
	flag.Usage = usage
	flag.StringVar(&worker.format, "f", worker.format, "compression format to use")
	flag.IntVar(&worker.concurrency, "p", worker.concurrency, "max concurrency")
	flag.BoolVar(&worker.decompress, "d", false, "decompress the source")
	flag.BoolVar(&help, "h", false, "help")
	flag.BoolVar(&help, "help", false, "help")
}

// this allows for defers to run and exiting with a return code.
func main() {
	os.Exit(realMain())
}

func realMain() int {
	flag.Parse()
	if help {
		flag.Usage()
		fmt.Fprint(os.Stderr, "Flags:\n")
		fmt.Fprint(os.Stderr, "\n")
		flag.PrintDefaults()
		return 0
	}

	worker.files = flag.Args() // this is the list of files
	if len(worker.files) == 0 {
		flag.Usage()
		return 1
	}
	err := worker.Work()
	if err != nil {
		if err == ErrProcess {
			return len(worker.errs) // the number of operations (files) that errored is the return code
		}
		return 1
	}

	return 0
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", name)
	fmt.Fprint(os.Stderr, "  compress:\n")
	fmt.Fprintf(os.Stderr, "    %s filename(s)...\n", name)
	fmt.Fprint(os.Stderr, "\n")
	fmt.Fprint(os.Stderr, "  decompress:\n")
	fmt.Fprintf(os.Stderr, "    %s -d filename(s)...\n", name)
	fmt.Fprint(os.Stderr, "\n")
	fmt.Fprint(os.Stderr, "  help:\n")
	fmt.Fprintf(os.Stderr, "    %s -h\n", name)
	fmt.Fprintf(os.Stderr, "    %s --help\n", name)
	fmt.Fprint(os.Stderr, "\n")
}
