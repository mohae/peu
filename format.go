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
package peu

import (
	"errors"
	"io"

	"github.com/mohae/magicnum/mcompress"
)

var ErrUnsupported = errors.New("unsupported compression format")

// CompressionFormat checks to see if the data in the provided reader uses
// a supported compression format. If it does not, an UnsupportedErr is
// returned.
func CompressionFormat(r io.ReaderAt) (mcompress.Format, error) {
	f, err := mcompress.GetFormat(r)
	if err != nil {
		return mcompress.Unknown, err
	}
	// see if the format is a supported on
	switch f {
	case mcompress.GZip:
		return f, nil
	case mcompress.LZ4:
		return f, nil
	default:
		return f, ErrUnsupported
	}
}
