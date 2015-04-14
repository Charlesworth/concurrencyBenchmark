# concurrencyBenchmark
Benchmarking go-routines: thread matching vs the go scheduler using an atomic counter

###Results (i5-4690K):

| Benchmark | Ops per second |
| ---- | ---- |
| goroutines match CPU count | 52765551 |
| many goroutines using runtime.Gosched | 10129459 |
