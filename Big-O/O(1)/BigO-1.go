package bigo

// FirstChar takes a slice of integers and returns the first element.
func FirstChar(arr []int16) int16 {
	if len(arr) != 0 {
		return arr[0]
	}
	return 0
}
