// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package numeric

// #include "unicode/utypes.h"
// #include "bridge.h"
import "C"
import (
	"fmt"
	"unsafe"

	"github.com/uber-go/icu4go/constants"
)

// Format is function to format a number for a locale
func Format(l string, n float64) (string, error) {
	a := C.double(n)
	locale := C.CString(l)
	defer func() { C.free(unsafe.Pointer(locale)) }()

	resSize := C.size_t(constants.BufSize512 * C.sizeof_char)
	res := (*C.char)(C.malloc(resSize))
	defer func() { C.free(unsafe.Pointer(res)) }()

	err := C.formatNumber(a, locale, res, resSize)
	if err > 0 {
		return "", fmt.Errorf("error (%d) formatting a number %f", err, n)
	}
	return C.GoString(res), nil
}
