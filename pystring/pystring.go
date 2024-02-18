package pystring

type PyString string

func New(s string) PyString {
	return PyString(s)
}

func (s PyString) String() string {
	return string(s)
}

func (s PyString) Format(vargs []any, kwarg map[string]any) (PyString, error) {
	res, err := DefaultDialect.Format(string(s), vargs, kwarg)
	if err != nil {
		return "", err
	}
	return PyString(res), nil
}

func (s PyString) FormatWithDialect(vargs []any, kwarg map[string]any,
	d Dialect) (PyString, error) {
	res, err := d.Format(string(s), vargs, kwarg)
	if err != nil {
		return "", err
	}
	return PyString(res), nil
}
