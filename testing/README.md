# What is it about?
Here is an examples related to testing in GO.

# How to run?

## Unit-tests
Run tests:

>go test pkg/*

OR

>go test pkg/reverse.go pkg/reverse_test.go

*Note: specify '_test.go' file, and a file with testable code.*

To run specific test:
>go test pkg/* **-run 'TestReverse'**

Verbose mode:
>go test pkg/* **-v** -run 'TestReverse'


## Test coverage
Coverage:
>go test pkg/* **-cover**

Run tests and capture coverage details:
>go test pkg/* **-coverprofile**=coverage.out

Generate HTML report showing lines run and not run:
>go tool cover **-html**=coverage.out


## Benchmarks
Just run benchmark:
>go test ./... **-run=xxx -bench=.**

Add memory stat:
>go test ./... -run=xxx -bench=. **-benchmem**

Profile for particular benchmark:
>go test ./... -bench=**BenchmarkReverse** -cpuprofile=**cpu.prof**

Observe profile:
``` bash
go tool pprof cpu.prof
(pprof) top 7 # see top N consumption operations
(pprof) list Reverse # see detailed consumption for each line in Reverse
```
