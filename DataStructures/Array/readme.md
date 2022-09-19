
# Arrays


## Static VS Dynamic Arrays
In Golang, there are both versions of arrays: static and dynamic.
- Static arrays are fixed in size and cannot be resized -> arrays.
- Dynamic arrays are resizable and can be resized - > slices.

## Static Array Declaration
An array is a value-type.
```go
var nums = [4]int{1, 2, 3, 4} // int array with values
var arr = new([10]int32) // int32 array with 10 elements created with new()
var arr [10] int64 // int64 array with 10 elements
```


## Dynamic Array Declaration
A slice is a reference type. 
```go
var slice []string // a string slice
var slice []int64 // a int64 slice
var slice = []int{1, 2, 3, 4} // a int slice with values
var slice make([]string, 10) // a string slice with 10 elements
```

NOTE: never use a pointer to a slice because a slice itself is a pointer to an array.

Please, don't be shy to check [my blog post about arrays and slices in Golang](https://eremeev.ca/posts/golang-copy-vs-append/)