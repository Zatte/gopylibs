# PyString

An attempt to get as similar behavior as possible that exists in python.

Source reference: https://docs.python.org/3/library/string.html#string.Formatter


## Dialects
Python from time to time has differences in implementation meaning getting
1:1 with python needs to be specified towards which python version. The idea
here is to specify the feature flags required to reach compatibility's in different
"Dialects" that are responsible for formatting strings.


## Out of scope
Python 2.X `(string % dict)` compatibility. 3.X is enough. Small PRs could be welcome if not too
obstructive for maintainability.

## TODO

Format() - Support locale aware formatting
- [ ] The 'z' option coerces negative zero floating-point values to positive zero after rounding to the format precision. This option is only valid for floating-point presentation types.
- [ ] The ',' option signals the use of a comma for a thousands separator. For a locale aware separator, use the 'n' integer presentation type instead
- [ ] The '_' option signals the use of an underscore for a thousands separator for floating point presentation types and for integer presentation type 'd'. For integer presentation types 'b', 'o', 'x', and 'X', underscores will be inserted every 4 digits. For other presentation types, specifying this option is an error.


Other features
- [ ] Support Template strings https://docs.python.org/3/library/string.html#template-strings