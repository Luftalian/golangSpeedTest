## mapとstructの配列の実行時間比較

```bash
$ go test -bench=.
goos: linux
goarch: amd64
pkg: example.com
cpu: 11th Gen Intel(R) Core(TM) i5-1135G7 @ 2.40GHz
BenchmarkLargeMapLength-8     	1000000000	         0.2412 ns/op
BenchmarkLargeArrayLength-8   	1000000000	         0.2400 ns/op
PASS
ok  	example.com	0.537s
```
