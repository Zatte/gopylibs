package pystring

import "testing"

func TestStartsWith(t *testing.T) {
	tests := []struct {
		s        string
		prefix   string
		start    *int
		end      *int
		expected bool
	}{
		{"abcdef", "abc", nil, nil, true},
		{"abcdef", "def", nil, nil, false},
		{"abcdef", "bcd", nil, nil, false},
		{"abcdef", "a", nil, nil, true},
		{"abcdef", "cde", nil, nil, false},
		{"abcdef", "abc", nil, intP(2), false},
		{"abcdef", "cd", intP(2), nil, true},
		{"abcdef", "abc", intP(1), intP(3), false},
		{"abcdef", "abc", intP(0), intP(2), false},
		{"abcdef", "abc", intP(0), intP(4), true},
	}

	for _, test := range tests {
		t.Run(test.s, func(t *testing.T) {
			result := PyString(test.s).StartsWith(test.prefix, test.start, test.end)
			if result != test.expected {
				t.Errorf("Expected %q.StartsWith(%q, %v, %v) %v but got %v", test.s, test.prefix, test.start, test.end, test.expected, result)
			}
		})
	}
}
