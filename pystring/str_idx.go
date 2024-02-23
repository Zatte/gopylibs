package pystring

// Idx replicates indexing behavior in python. As such it supports negative
// indexing and don't crash on out of bound indexes.
func Idx(s string, start, end *int) string {
	sLen := len(s)
	actualStart := 0
	actualEnd := len(s)
	if end != nil {
		if *end < 0 {
			actualEnd = sLen + *end
		}
		if *end < sLen {
			actualEnd = *end
		}
		if actualEnd < 0 {
			actualEnd = 0
		}
	}

	if start != nil {
		if *start < 0 {
			actualStart = sLen + *start
		} else {
			actualStart = *start
		}
		if actualStart < 0 {
			actualStart = 0
		}
	}

	return s[actualStart:actualEnd]
}

// Idx replicates indexing behavior in python. As such it supports negative
// indexing and don't crash on out of bound indexes.
func (pys PyString) Idx(start, end *int) PyString {
	return PyString(Idx(string(pys), start, end))
}
