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

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/uber-go/icu4go/constants"
)

var (
	testEpochSecs = int64(1483360520)
	testLocale    = "de-DE"
)

func TestDateTimeFormatter_FormatTime(t *testing.T) {
	assert := assert.New(t)
	formattedDateTime, err := Format(testLocale, time.Unix(testEpochSecs, 0), Full, constants.Utc)
	assert.NoError(err)
	assert.Equal("Montag, 2. Januar 2017 um 12:35:20 GMT", formattedDateTime)

	formattedDateTime, err = Format(testLocale, time.Unix(testEpochSecs, 0), Mid, constants.Utc)
	assert.NoError(err)
	assert.Equal("02.01.2017, 12:35:20", formattedDateTime)

	formattedDateTime, err = Format(testLocale, time.Unix(testEpochSecs, 0), Long, constants.Utc)
	assert.NoError(err)
	assert.Equal("2. Januar 2017 um 12:35:20 GMT", formattedDateTime)

	formattedDateTime, err = Format(testLocale, time.Unix(testEpochSecs, 0), Short, constants.Utc)
	assert.NoError(err)
	assert.Equal("02.01.17, 12:35", formattedDateTime)
}

func TestDateTimeFormatter_GetPattern(t *testing.T) {
	assert := assert.New(t)
	pattern, err := GetPatternFromICU(testLocale, Full, false)
	assert.NoError(err)
	assert.Equal("EEEE, d. MMMM y 'um' HH:mm:ss zzzz", pattern)

	pattern, err = GetPatternFromICU(testLocale, Mid, false)
	assert.NoError(err)
	assert.Equal("dd.MM.y, HH:mm:ss", pattern)

	pattern, err = GetPatternFromICU(testLocale, Long, false)
	assert.NoError(err)
	assert.Equal("d. MMMM y 'um' HH:mm:ss z", pattern)

	pattern, err = GetPatternFromICU(testLocale, Short, false)
	assert.NoError(err)
	assert.Equal("dd.MM.yy, HH:mm", pattern)
}

func TestDateTimeFormatter_GetPatternFuncOverride(t *testing.T) {

	oldValue := GetPattern
	defer func() {
		GetPattern = oldValue
	}()

	// Override the func to return Short pattern always
	GetPattern = func(string, FormatType, bool) (string, error) {
		return "dd.MM.yy, HH:mm", nil
	}

	assert := assert.New(t)
	formattedDateTime, err := Format(testLocale, time.Unix(testEpochSecs, 0), Full, constants.Utc)
	assert.NoError(err)
	// We should get Short format even if we ask for Full as override is in place
	assert.Equal("02.01.17, 12:35", formattedDateTime)

}
