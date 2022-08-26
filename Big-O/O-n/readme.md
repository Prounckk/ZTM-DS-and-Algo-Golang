# Big O(n) aka Linear Order

An algorithm that time consuming of which will grow linearly and in direct proportion to the size of the input data


Here is an example of my benchmark. It might have different numbers on your local. Note the difference in results between small and big slice. 
```
$ go test -bench=. -benchtime=10s -benchmem
goos: darwin
goarch: arm64
pkg: github.com/prounckk/ZTM-DS-and-Algo-Golang/Big-O/O-n
Benchmark_SumOfAllIntegersInSlice/small_slice-8     1000000000  4.066 ns/op 0 B/op  0 allocs/op
Benchmark_SumOfAllIntegersInSlice/big_slice-8       371872671   32.26 ns/op 0 B/op  0 allocs/op
PASS
ok  	github.com/prounckk/ZTM-DS-and-Algo-Golang/Big-O/O-n	19.850s
```


If you need to create a huge slice, click here: https://onlineintegertools.com/create-integer-array