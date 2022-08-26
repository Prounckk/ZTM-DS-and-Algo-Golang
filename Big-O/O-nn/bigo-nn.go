package bigo

// SumOfAllIntegersInTwoSlicesByIndex takes a slice of integers and returns the sum of all elements if the element present in the indexes slice.
// don't even ask me why this function exists, I can't find any rational explanation.
func SumOfAllIntegersInTwoSlicesByIndex(slice []int, indexes []int) int {
	res := 0
	for int := range slice {
		for index := range indexes {
			if index == int {
				res = res + int
			}
		}
	}
	return res
}
