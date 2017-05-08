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

package datetime

// #include "bridge.h"
import "C"

import (
	"fmt"
	"time"
	"unsafe"

	"github.com/uber-go/icu4go/constants"
)

// formatDateTime to format a given date/time into string based on locale and format type
func formatDateTime(t time.Time, format string, timeZone string, localeCode string) (string, error) {
	localeID := C.CString(localeCode)
	epochMillis := C.double(t.UnixNano() / int64(time.Millisecond))
	resultSize := C.size_t(constants.BufSize512 * C.sizeof_char)
	result := (*C.char)(C.malloc(resultSize))
	pattern := C.CString(format)
	tz := C.CString(timeZone)

	defer func() {
		C.free(unsafe.Pointer(localeID))
		C.free(unsafe.Pointer(result))
		C.free(unsafe.Pointer(pattern))
		C.free(unsafe.Pointer(tz))
	}()

	if err := C.formatDateTime(
		epochMillis,
		localeID,
		pattern,
		tz,
		result,
		resultSize,
	); err > 0 {
		return "", fmt.Errorf("errno %d occurred formatting the date for locale %s", err, localeCode)
	}

	return C.GoString(result), nil
}

// getDateTimeFormat to retrieve a date time format for a locale and style
func getDateTimePattern(localeCode string, style C.UDateFormatStyle, localized bool) (string, error) {
	localeID := C.CString(localeCode)
	resultSize := C.size_t(constants.BufSize512 * C.sizeof_char)
	result := (*C.char)(C.malloc(resultSize))

	defer func() {
		C.free(unsafe.Pointer(localeID))
		C.free(unsafe.Pointer(result))
	}()

	if err := C.getDateTimePattern(
		localeID,
		style,
		C.bool(localized),
		result,
		resultSize,
	); err > 0 {
		return "", fmt.Errorf("errno %d occurred getting date time pattern for locale %s", err, localeCode)
	}

	return C.GoString(result), nil
}
