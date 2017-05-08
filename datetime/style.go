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

// FormatType is one of [Full, Mid, Short, Long]
type FormatType struct {
	Name     string
	ICUStyle C.UDateFormatStyle
}

var (
	// Full format type
	Full = FormatType{"Full", C.UDAT_FULL}
	// Mid format type
	Mid = FormatType{"Mid", C.UDAT_MEDIUM}
	// Short format type
	Short = FormatType{"Short", C.UDAT_SHORT}
	// Long format type
	Long = FormatType{"Long", C.UDAT_LONG}

	lookup = map[string]FormatType{
		"Full":  Full,
		"Mid":   Mid,
		"Short": Short,
		"Long":  Long,
	}
)

// GetFormatType gets the format type for a string
func GetFormatType(style string) FormatType {
	return lookup[style]
}
