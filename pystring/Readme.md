# PyString

An attempt to get as similar behavior as possible that exists in python.

Source reference: https://docs.python.org/3/library/string.html#string.Formatter


## Dialects
Occasionally, Python implementations may vary between versions, necessitating
specification of the Python version to achieve direct parity. The aim is to
outline required feature flags necessary to attain compatibility in "dialects".


## Out of scope
Python 2.X `(string % dict)` compatibility. 3.X is enough. Small PRs could be welcome if not too
obstructive for maintainability.

## TODO

Format() - Support locale aware formatting
- [x] The 'z' option coerces negative zero floating-point values to positive zero after rounding to the format precision. This option is only valid for floating-point presentation types.
- [ ] The ',' option signals the use of a comma for a thousands separator. For a locale aware separator, use the 'n' integer presentation type instead
- [ ] The '_' option signals the use of an underscore for a thousands separator for floating point presentation types and for integer presentation type 'd'. For integer presentation types 'b', 'o', 'x', and 'X', underscores will be inserted every 4 digits. For other presentation types, specifying this option is an error.

Str Functions
- [x] [capitalize](https://docs.python.org/3/library/stdtypes.html#str.capitalize)
- [x] [casefold](https://docs.python.org/3/library/stdtypes.html#str.casefold)
- [x] [center](https://docs.python.org/3/library/stdtypes.html#str.center)
- [x] [count](https://docs.python.org/3/library/stdtypes.html#str.count)
- [] [encode](https://docs.python.org/3/library/stdtypes.html#str.encode)
- [x] [endswith](https://docs.python.org/3/library/stdtypes.html#str.endswith)
- [x] [expandtabs](https://docs.python.org/3/library/stdtypes.html#str.expandtabs)
- [x] [find](https://docs.python.org/3/library/stdtypes.html#str.find)
- [x] [format](https://docs.python.org/3/library/stdtypes.html#str.format)
- [x] [format_map](https://docs.python.org/3/library/stdtypes.html#str.format_map) - Aliased to format since the difference in golang is not relevant
- [] [index](https://docs.python.org/3/library/stdtypes.html#str.index)
- [x] [isalnum](https://docs.python.org/3/library/stdtypes.html#str.isalnum)
- [x] [isalpha](https://docs.python.org/3/library/stdtypes.html#str.isalpha)
- [x] [isascii](https://docs.python.org/3/library/stdtypes.html#str.isascii)
- [x] [isdecimal](https://docs.python.org/3/library/stdtypes.html#str.isdecimal)
- [x] [isdigit](https://docs.python.org/3/library/stdtypes.html#str.isdigit)
- [] [isidentifier](https://docs.python.org/3/library/stdtypes.html#str.isidentifier)
- [] [islower](https://docs.python.org/3/library/stdtypes.html#str.islower)
- [] [isnumeric](https://docs.python.org/3/library/stdtypes.html#str.isnumeric)
- [] [isprintable](https://docs.python.org/3/library/stdtypes.html#str.isprintable)
- [] [isspace](https://docs.python.org/3/library/stdtypes.html#str.isspace)
- [] [istitle](https://docs.python.org/3/library/stdtypes.html#str.istitle)
- [] [isupper](https://docs.python.org/3/library/stdtypes.html#str.isupper)
- [] [join](https://docs.python.org/3/library/stdtypes.html#str.join)
- [] [ljust](https://docs.python.org/3/library/stdtypes.html#str.ljust)
- [] [lower](https://docs.python.org/3/library/stdtypes.html#str.lower)
- [] [lstrip](https://docs.python.org/3/library/stdtypes.html#str.lstrip)
- [] [maketrans](https://docs.python.org/3/library/stdtypes.html#str.maketrans)
- [] [partition](https://docs.python.org/3/library/stdtypes.html#str.partition)
- [] [removeprefix](https://docs.python.org/3/library/stdtypes.html#str.removeprefix)
- [] [removesuffix](https://docs.python.org/3/library/stdtypes.html#str.removesuffix)
- [] [replace](https://docs.python.org/3/library/stdtypes.html#str.replace)
- [] [rfind](https://docs.python.org/3/library/stdtypes.html#str.rfind)
- [] [rindex](https://docs.python.org/3/library/stdtypes.html#str.rindex)
- [] [rjust](https://docs.python.org/3/library/stdtypes.html#str.rjust)
- [] [rpartition](https://docs.python.org/3/library/stdtypes.html#str.rpartition)
- [] [rsplit](https://docs.python.org/3/library/stdtypes.html#str.rsplit)
- [] [rstrip](https://docs.python.org/3/library/stdtypes.html#str.rstrip)
- [] [split](https://docs.python.org/3/library/stdtypes.html#str.split)
- [] [splitlines](https://docs.python.org/3/library/stdtypes.html#str.splitlines)
- [] [startswith](https://docs.python.org/3/library/stdtypes.html#str.startswith)
- [] [strip](https://docs.python.org/3/library/stdtypes.html#str.strip)
- [] [swapcase](https://docs.python.org/3/library/stdtypes.html#str.swapcase)
- [] [title](https://docs.python.org/3/library/stdtypes.html#str.title)
- [] [translate](https://docs.python.org/3/library/stdtypes.html#str.translate)
- [] [upper](https://docs.python.org/3/library/stdtypes.html#str.upper)
- [] [zfill](https://docs.python.org/3/library/stdtypes.html#str.zfill)

Other features
- [ ] Support Template strings https://docs.python.org/3/library/string.html#template-strings
