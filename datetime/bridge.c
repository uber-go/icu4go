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

#include "bridge.h"

#include <unicode/udat.h>
#include <unicode/ustring.h>
#include <unicode/uvernum.h>
#include <string.h>

const UErrorCode formatDateTime(double epochMillis,
                                const char* localeId,
                                const char* pattern,
                                const char* tz,
                                char* const result,
                                const size_t resultCapacity) {

    UErrorCode status = U_ZERO_ERROR;
    UChar outBuffer[resultCapacity];

    UChar tzUchar[strlen(tz) * sizeof(UChar)];
    u_uastrcpy(tzUchar, tz);

    UChar patternUchar[strlen(pattern) * sizeof(UChar)];
    u_uastrcpy(patternUchar, pattern);

    UDateFormat* dfmt = udat_open(
        #if U_ICU_VERSION_MAJOR_NUM >= 5
            UDAT_PATTERN,   // The dateStyle, UDAT_PATTERN means take it from the pattern supplied
            UDAT_PATTERN,   // The timeStyle, UDAT_PATTERN means take it from the pattern supplied
        #else
            UDAT_IGNORE,
            UDAT_IGNORE,
        #endif
        localeId,       // The localeId
        tzUchar,        // The timezone Id
        -1,             // The timezone Id len ( -1 because tzUchar is null terminated )
        patternUchar,   // The custom datetime pattern
        -1,             // The pattern len ( -1 because patternUchar is null terminated )
        &status         // The output error status
    );
    if (U_FAILURE(status)) {
        return status;
    }

    udat_format(
        dfmt,                   // The datetime formatter
        (UDate)epochMillis,     // Millseconds since epoch
        outBuffer,              // The output buffer to store the result
        resultCapacity,         // Capacity of the output buffer
        NULL,                   // Position ( NULL because we want it from start )
        &status                 // The output error status
    );
    if (U_FAILURE(status)) {
        return status;
    }

    u_austrcpy(result, outBuffer);
    return status;
 }

const UErrorCode getDateTimePattern(const char* localeId,
                                   UDateFormatStyle style,
                                   bool localized,
                                   char* const result,
                                   const size_t resultCapacity) {

    UErrorCode status = U_ZERO_ERROR;
    UChar outBuffer[resultCapacity];

    UDateFormat* dfmt = udat_open(
        style,          // The dateStyle
        style,          // The timeStyle
        localeId,       // The localeId
        0,              // The timezone Id 0 meaning use the default timezone
        -1,             // The timezone Id len
        NULL,           // The datetime pattern
        -1,             // The pattern len
        &status         // The output error status
    );
    if (U_FAILURE(status)) {
        return status;
    }

    udat_toPattern(
        dfmt,                   // The datetime formatter
        localized,              // localized
        outBuffer,              // The output buffer to store the result
        resultCapacity,         // Capacity of the output buffer
        &status                 // The output error status
    );
    if (U_FAILURE(status)) {
        return status;
    }

    u_austrcpy(result, outBuffer);
    return status;
 }

