package pystring

// there has been multiples changes in python in regards to how format specifiers are handled
// to enable all possible formats we captures these changes feature flags which can be opted
// in our out into.
var DefaultDialect = NewDialect(3.11)
var DialectPython3_11 = NewDialect(3.11)
var DialectPython3_0 = NewDialect(3.0)

type Dialect struct {
	zeroPaddingAlignment  rune
	tryTypeJugglingString bool
}

type DialectOption func(*Dialect)

func NewDialect(version float64, options ...DialectOption) Dialect {
	res := Dialect{
		zeroPaddingAlignment:  '=',
		tryTypeJugglingString: false,
	}

	// Changed in version 3.10: Preceding the width field by '0' no longer affects the default alignment for strings.
	// https://docs.python.org/3/library/string.html#string.Formatter
	if version >= 3.10 {
		res.zeroPaddingAlignment = 0
	}

	for _, option := range options {
		option(&res)
	}

	return res
}

func (d Dialect) CloneWithOptions(options ...DialectOption) Dialect {
	res := d
	for _, option := range options {
		option(&res)
	}
	return res
}

func WithTypeJugglingString(d *Dialect) {
	d.tryTypeJugglingString = true
}
