package binarysearch

func BinarySearch[T int | float64 | string](ordenedList []T, target T) int {
	low, high := 0, len(ordenedList)-1

	for low <= high {
		mid := (high + low) / 2
		guess := ordenedList[mid]
		switch {
		case guess == target:
			return mid

		case guess < target:
			low = mid + 1
		default:
			high = mid - 1
		}
	}
	return -1
}
