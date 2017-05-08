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

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// The Test Suite
type LocaleTestSuite struct {
	suite.Suite
}

func TestLocales(t *testing.T) {
	suite.Run(t, new(LocaleTestSuite))
}

func (s *LocaleTestSuite) TestNormalizeLocale() {
	assert := assert.New(s.T())
	l := Normalized("en")
	assert.Equal("en", l)

	l = Normalized("en_US")
	assert.Equal("en-US", l)

	l = Normalized("en_us")
	assert.Equal("en-US", l)

	l = Normalized("en_Latn_US_REVISED")
	assert.Equal("en-US", l)

	l = Normalized("en@collation=phonebook;calendar=islamic-civil")
	assert.Equal("en", l)

	l = Normalized("en_Latn_US_REVISED@currency=USD")
	assert.Equal("en-US", l)

	l = Normalized("en_Latn")
	assert.Equal("en", l)
}

func (s *LocaleTestSuite) TestGetCountrycode() {
	assert := assert.New(s.T())
	assert.Equal(GetCountryCode("en-US"), "US")
}

func (s *LocaleTestSuite) TestGetLanguageCode() {
	assert := assert.New(s.T())
	assert.Equal(GetLanguageCode("en-US"), "en")
}
