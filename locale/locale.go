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

package locale

// #include "bridge.h"
import "C"
import (
	"strings"
	"unsafe"

	"github.com/uber-go/icu4go/constants"
)

// Normalized returns the normalized version of the input locale code. It will first try to get the
// language code from the input locale. If it cant get a valid language then it will return "en" as a locale.
// It will then try to get country and if it cant get country then it will just return that language as a locale.
// At last, it will return language and country joined by hyphen.
func Normalized(code string) string {
	cc := getCountryCode(code)
	if cc == constants.EmptyString {
		return getLanguageCode(code)
	}

	return strings.Join([]string{getLanguageCode(code), cc}, constants.Hyphen)
}

// GetCountryCode find the country code for a locale
func GetCountryCode(code string) string {
	return getCountryCode(code)
}

// GetLanguageCode returns the language code for this locale
func GetLanguageCode(code string) string {
	return getLanguageCode(code)
}

func getLanguageCode(code string) string {
	l := C.CString(code)
	defer func() { C.free(unsafe.Pointer(l)) }()

	resSize := C.size_t(constants.BufSize512 * C.sizeof_char)
	result := (*C.char)(C.malloc(resSize))
	defer func() { C.free(unsafe.Pointer(result)) }()

	err := C.getLanguageCode(l, result, constants.BufSize512*C.sizeof_char)
	if err > C.U_ZERO_ERROR {
		return constants.DefaultLanguageCode
	}

	return C.GoString(result)
}

func getCountryCode(code string) string {
	l := C.CString(code)
	defer func() {
		C.free(unsafe.Pointer(l))
	}()

	resSize := C.size_t(constants.BufSize512 * C.sizeof_char)
	result := (*C.char)(C.malloc(resSize))
	defer func() {
		C.free(unsafe.Pointer(result))
	}()

	err := C.getCountryCode(l, result, constants.BufSize512*C.sizeof_char)
	if err > C.U_ZERO_ERROR {
		return constants.DefaultCountryCode
	}

	return C.GoString(result)
}
