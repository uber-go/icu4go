# icu4go [![Build Status][ci-img]][ci] [![Coverage Status][cov-img]][cov]

## A go binding for icu4c library

This library provides functions for locale, date, time, currency and number formatting using icu4c.
We will add more functions as and when needed. In order to use this library you must install icu4c library
for your platform and define CGO_CFLAGS and CGO_LDFLAGS in your Makefile. You can take a look at this library's
[Makefile](Makefile) for example


### Install ICU4C
1. On Mac:
http://macappstore.org/icu4c/

2. On Linux:
http://www.linuxfromscratch.org/blfs/view/8.0/general/icu.html



## Usage
#### Locale
All locale specific functions reside in the `locale` package.

```
import "github.com/uber-go/icu4go/locale"

```

To normalize a locale use Normalized function as below:

```
locale.Normalized("en_US")
> en-US
```

To get a country code from a locale use GetCountryCode as below:

```
locale.GetCountrycode("en-US")
> US
```

To get a language code from locale use GetLanguageCode as below:

```
locale.GetLanguageCode("en-US")
> en
```

#### Number formats
All number formatting functions reside in `number` package.

```
import "github.com/uber-go/icu4go/number"
```

To format a number into a locale use Format function as below:
```
number.Format("en-US", 123456.78)
> 123,456.78
```

#### Currency formats
All currency formatting functions reside in `currency` package.

```
import "github.com/uber-go/icu4go/currency"
```

To format a currency into a locale use Format function as below:
```
currency.Format("en-US", 123456.78, "USD")
> $123,456.78

```

#### DateTime Formats
All date and time formatting functions reside in `datetime` package. The supported Styles are Short, Mid, Long, and Full.

```
import "github.com/uber-go/icu4go/datetime"

```

To format a date time into locale use Format function as below:

```
datetime.Format("de-DE", time.Now(), datetime.Short, "UTC")
> 02.01.17, 12:35

```

To format a date time into locale with your pattern Override the GetPattern function as below:
```
// Override the func to return Short pattern always
GetPattern = func(string, FormatType, bool) (string, error) {
    return "dd.MM.yy, HH:mm", nil
}

datetime.Format("de-DE", time.Now(), datetime.Full, "UTC")
> 02.01.17, 12:35
```

## Development Status
Stable.

<hr>
Released under the [MIT License](LICENSE.txt).

[ci-img]: https://travis-ci.org/uber-go/icu4go.svg?branch=master
[ci]: https://travis-ci.org/uber-go/icu4go
[cov-img]: https://coveralls.io/repos/github/uber-go/icu4go/badge.svg?branch=master
[cov]: https://coveralls.io/github/uber-go/icu4go?branch=master
