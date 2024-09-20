package binarysearch

func RecursiveBinarySearch[T int | float64 | string](orderedList []T, target T) int {
	return recursiveSearch(orderedList, target, 0, len(orderedList)-1)
}

func recursiveSearch[T int | float64 | string](orderedList []T, target T, low, high int) int {
	if low > high {
		return -1
	}

	mid := (low + high) / 2
	guess := orderedList[mid]

	switch {
	case guess == target:
		return mid
	case guess < target:
		return recursiveSearch(orderedList, target, mid+1, high)
	default:
		return recursiveSearch(orderedList, target, low, mid-1)
	}
}
