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

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "unicode/unum.h"
#include "unicode/ucurr.h"
#include "unicode/ustring.h"

#include "bridge.h"

UErrorCode getCurrencySymbol(const char *currencyCode, const char* locale, char* result, const size_t resultLength)
{
	UErrorCode status = U_ZERO_ERROR;
	const int32_t bufSize = resultLength / sizeof(UChar); // considering it's UNICODE.
	int32_t needed;

	UChar cc[strlen(currencyCode) * sizeof(UChar)];
	u_uastrncpy(cc, currencyCode, strlen(currencyCode));

	UBool isChoiceFormat = FALSE;
	int32_t resSize;
	const UChar *res = ucurr_getName(cc, locale, UCURR_SYMBOL_NAME, &isChoiceFormat, &resSize, &status);
	if (!U_FAILURE(status)) {
		u_austrcpy(result, res);
	}

	return status;
}

UErrorCode formatCurrency(const double a, const char *currencyCode, const char* locale, char* result, const size_t resultLength)
{
	UErrorCode status = U_ZERO_ERROR;
	const int32_t bufSize = resultLength / sizeof(UChar); // considering it's UNICODE.
	int32_t needed;

	// TODO (ds): understand how expensive this is.
	/* Create a formatter for the US locale */
	UNumberFormat *fmt = unum_open(
			UNUM_CURRENCY,     /* style         */
			0,                 /* pattern       */
			0,                 /* patternLength */
			locale,            /* locale        */
			0,                 /* parseErr      */
			&status);

	if (!U_FAILURE(status)) {
		UChar buf[strlen(currencyCode) * sizeof(UChar)];
		u_uastrncpy(buf, currencyCode, strlen(currencyCode));
		unum_setTextAttribute(fmt, UNUM_CURRENCY_CODE, buf, u_strlen(buf), &status);
	}
	if (!U_FAILURE(status)) {
		/* Use the formatter to format the number as a string
			 in the given locale. The return value is the buffer size needed.
			 We pass in NULL for the UFieldPosition pointer because we don't
			 care to receive that data. */
		UChar buf[bufSize];
		needed = unum_formatDouble(fmt, a, buf, bufSize, NULL, &status);
		/**
		 * u_austrcpy docs from the header:
		 *
		 * Copy ustring to a byte string encoded in the default codepage.
		 * Adds a null terminator.
		 * Performs a UChar to host byte conversion
		 */
		u_austrcpy(result, buf);

		/* Release the storage used by the formatter */
		unum_close(fmt);
	}

	return status;
}

