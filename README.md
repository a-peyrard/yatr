# Yatr

_Yet Another Test Runner_

----

## Abstract

Everytime I'm going back to typescript, I'm disappointed by jest poor parallelism performance. 

This project was just an attempt to write a basic test runner for javascript with a golang dispatcher. 
The dispatcher is responsible for finding the specs files, and managing workers. Then each worker is having a 
dedicated node process to run the tests and is communicating with the dispatcher through stdin/stdout.

## Build

```bash
go build
```

## Usage

```bash
./yatr 10 /opt/dev/js/jest-workspace/yatr/sample/100slowSpecs
./yatr 10 /opt/dev/js/jest-workspace/yatr/sample/100specs
```

First argument is the number of workers, second argument is the path of the project.

## Quick perf analysis

With yatr, we can run 100 specs in 1.1s, while jest is taking ~2s. So we are losing only 100ms in the process, 
as the fastest we could go would be 1s (1000 tests with 10ms sleep in each test). 
```bash
$ time ./yatr 10 /opt/dev/js/jest-workspace/yatr/sample/100slowSpecs
...
./yatr 10 /opt/dev/js/jest-workspace/yatr/sample/100slowSpecs  2.21s user 6.29s system 760% cpu 1.118 total

$ time jest --maxWorkers=10
...
jest --maxWorkers=10  13.39s user 1.91s system 770% cpu 1.987 total
```

And the difference is huge (10x), when running the 100specs sample with a single worker (1000 tests with no sleep):
```bash
./yatr 1 /opt/dev/js/jest-workspace/yatr/sample/100specs  0.01s user 0.03s system 27% cpu 0.158 total
jest --runInBand  1.62s user 0.50s system 108% cpu 1.960 total
```

## Conclusion

This is fast and that was fun to write. I also learned more about jest by reading the source code.
But this would require a lot of work to have something usable. Even without parity with jest we would need
at a bare minimum to support:
- typescript (a lot of time is spent in the transpilation, so it would requires some smart incremental build)
- async 
- mocks (we would probably need a custom resolver like jest)
- configuration 
- proper before/after hooks management

I'm not sure it's worth the effort, but it was a fun experiment.