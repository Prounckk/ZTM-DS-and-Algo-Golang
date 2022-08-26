package bigo

// SumOfAllIntegersInSlice takes a slice of integers and returns the sum of all elements.
func SumOfAllIntegersInSlice(arr []int) int {
	res := 0
	for int := range arr {
		res = res + int
	}
	return res
}
