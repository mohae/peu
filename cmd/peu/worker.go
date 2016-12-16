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
	"errors"
	"fmt"
	"os"
	"sync"

	"github.com/mohae/magicnum/compress"
	"github.com/mohae/peu"
)

var ErrProcess = errors.New("processing error")

// Worker holds information about what needs to be done and the results.
type Worker struct {
	decompress      bool             // false == compress
	format          string           // the compression format
	compress.Format                  // Format per format
	files           []string         // files to process
	concurrency     int              // the maximum number of workers in the pool
	errs            map[string]error // map of errors that occured with the filename as the key
	n               int64            // number of bytes processed
	mu              sync.Mutex
}

// Work distrubutes and manages the work. If any processing results in an
// error, the errCnt is incremented and a ErrProcess is returned. For more
// details the errs map can be checked. The key for the errs map is the
// filename that had the error.
//
// For compression operations, if the format specified is unknown or not
// supported, either compress.Unknown or a peu.Unsupported error will be
// returned.
func (w *Worker) Work() error {
	// make suer concurrency is at least 1
	if w.concurrency < 1 {
		w.concurrency = 1
	}
	if w.decompress {
		return ErrProcess
		//return w.Decompress()
	}
	w.Format = compress.ParseFormat(w.format)
	return w.Compress()
}

// Compress manages the compression process.
func (w *Worker) Compress() error {
	var wg sync.WaitGroup
	ch := make(chan string)
	wg.Add(w.concurrency)
	for i := 0; i < w.concurrency; i++ {
		go w.c(ch, &wg)
	}
	for _, v := range w.files {
		ch <- v
	}
	close(ch)
	wg.Wait()
	if len(w.errs) == 0 {
		fmt.Printf("%d files totaling %d bytes were successfully compressed\n", len(w.files), w.n)
		return nil
	}
	fmt.Printf("%d files totaling %d bytes processed with %d errors:\n", len(w.files), w.n, len(w.errs))
	for k, v := range w.errs {
		fmt.Printf("\t%s: %s", k, v)
	}
	return ErrProcess
}

// c is the compression worker. It works until the channel has been drained.
func (w *Worker) c(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		v, ok := <-ch
		if !ok {
			return
		}
		src, err := os.Open(v)
		if err != nil {
			w.mu.Lock()
			w.errs[v] = err
			w.mu.Unlock()
			continue
		}
		dstName := v + w.Format.Ext()
		// TODO use original files permissions? This would probably require
		// platform specific code.
		dst, err := os.OpenFile(dstName, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0777)
		if err != nil {
			w.mu.Lock()
			w.errs[v] = err
			w.mu.Unlock()
			src.Close()
			continue
		}
		n, err := peu.Compress(w.format, src, dst)
		src.Close()
		dst.Close()
		w.mu.Lock()
		w.n += n
		fmt.Println(n)
		if err != nil {
			w.errs[v] = err
		}
		w.mu.Unlock()
	}
}
