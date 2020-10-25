package binarysearch

func SearchInts(input []int, x int) int {
	var low = 0
	var high = len(input) - 1

	for low <= high {
		var mid = low + (high-low)/2
		if input[mid] == x {
			return mid
		} else if x > input[mid] {
			low = mid + 1
		} else if x < input[mid] {
			high = mid - 1
		}
	}
	return -1
}
