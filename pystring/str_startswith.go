package pystring

import "strings"

// Return True if string starts with the prefix, otherwise return False. prefix can also be a tuple of prefixes to look for. With optional start, test string beginning at that position. With optional end, stop comparing string at that position.
func StartsWith(s string, prefix string, start, end *int) bool {
	return strings.HasPrefix(Idx(s, start, end), prefix)
}

// Return True if string starts with the prefix, otherwise return False. prefix can also be a tuple of prefixes to look for. With optional start, test string beginning at that position. With optional end, stop comparing string at that position.
func (pys PyString) StartsWith(prefix string, start, end *int) bool {
	return StartsWith(string(pys), prefix, start, end)
}
