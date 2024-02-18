# GoPyLib

Libs used to mimic python functionality in go. Note, this is not a python interpreter in
golang such as [Gython](https://github.com/go-python/gpython) but rather these libs might
aid such libraries get feature completeness faster.

## Why!?
1) The idea is to use it in [Gonja](https://github.com/NikolaLohinski/gonja).

2) Wanted to see how far ChatGTP could do x-language translations (only
partial success i'm afraid but at point i was hooked)

3) fun project


## Howto use

Check each directory for specific usage;
- [pystring](https://github.com/Zatte/gopylibs/tree/main/pystring)


## Priorities

1) Correctness - As close alignment with python functionality as possible; even when python is stupid.
2) Performance - Since it is fun
3) More features can be added through feature flags.

## State
Experimental

## Testing
There are integrations tests that run validation against python which passes.

`go test ./...`

`go test ./... --tags=integration` - Requires `python3.11` to be in path (this is the current version of python i'm
testing against).

## Benchmarks  / Performance

*pystring* - Each automatic replacement blocks, {}, add one more allocation;
formatting directives vary in allocations needed based on settings.

```shell
$ go test --benchmem --bench=. ./... -v

cpu: AMD Ryzen 5 3600 6-Core Processor
BenchmarkFomatSingleReplacement-12                       8757632               138.7 ns/op             8 B/op          1 allocs/op
BenchmarkFomatSingleAutoReplacement-12                   6694593               180.0 ns/op            16 B/op          2 allocs/op
BenchmarkFomatDoubleReplacement-12                       4421235               269.2 ns/op             8 B/op          1 allocs/op
BenchmarkFomatDoubleAutoReplacement-12                   3469920               341.1 ns/op            16 B/op          3 allocs/op
BenchmarkFomatSingleReplacementWithPadding-12            3114746               387.8 ns/op            56 B/op          5 allocs/op
BenchmarkFomatDoubleReplacementWithPadding-12            1527813               783.7 ns/op           128 B/op         10 allocs/op
BenchmarkComplexSpec-12                                   966310               1153 ns/op            208 B/op         12 allocs/op

```
