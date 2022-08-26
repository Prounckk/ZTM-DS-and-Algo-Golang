# Big O(n^2) aka Quadratic

The number of operations it performs scales in proportion to the square of the input

Here is an example of my benchmark. It might have different numbers on your local. Note the difference in results between small and big slices. 
```
$ go test -bench=. -benchmem
goos: darwin
goarch: arm64
pkg: github.com/prounckk/ZTM-DS-and-Algo-Golang/Big-O/O-nn
Benchmark_SumOfAllIntegersInTwoSlicesByIndex/small_slice-8              62914952    18.46 ns/op  0 B/op 0 allocs/op
Benchmark_SumOfAllIntegersInTwoSlicesByIndex/big_slice_of_100_ints-8    2474331     489.0 ns/op  0 B/op  0 allocs/op
Benchmark_SumOfAllIntegersInTwoSlicesByIndex/a_little_bit_big_slice-8   240080      4797 ns/op   0 B/op  0 allocs/op
PASS
ok      github.com/prounckk/ZTM-DS-and-Algo-Golang/Big-O/O-nn   4.987s
```


If you need to create a huge slice, click here: https://onlineintegertools.com/create-integer-array