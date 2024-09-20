package binarysearch

import (
	"testing"
)

func TestRecursiveBinarySearch(t *testing.T) {
	tests := []struct {
		name     string
		list     []int
		target   int
		expected int
	}{
		{"Found in middle", []int{1, 3, 5, 7, 9}, 5, 2},
		{"Found at beginning", []int{1, 3, 5, 7, 9}, 1, 0},
		{"Found at end", []int{1, 3, 5, 7, 9}, 9, 4},
		{"Not found - too small", []int{1, 3, 5, 7, 9}, 0, -1},
		{"Not found - too large", []int{1, 3, 5, 7, 9}, 10, -1},
		{"Not found - in between", []int{1, 3, 5, 7, 9}, 4, -1},
		{"Empty list", []int{}, 5, -1},
		{"Single element - found", []int{5}, 5, 0},
		{"Single element - not found", []int{5}, 3, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RecursiveBinarySearch(tt.list, tt.target)
			if result != tt.expected {
				t.Errorf("RecursiveBinarySearch(%v, %v) = %v; want %v", tt.list, tt.target, result, tt.expected)
			}
		})
	}
}

func TestRecursiveBinarySearchFloat(t *testing.T) {
	list := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	result := RecursiveBinarySearch(list, 3.3)
	expected := 2
	if result != expected {
		t.Errorf("RecursiveBinarySearch(%v, 3.3) = %v; want %v", list, result, expected)
	}
}

func TestRecursiveBinarySearchString(t *testing.T) {
	list := []string{"apple", "banana", "cherry", "date", "elderberry"}
	result := RecursiveBinarySearch(list, "cherry")
	expected := 2
	if result != expected {
		t.Errorf("RecursiveBinarySearch(%v, \"cherry\") = %v; want %v", list, result, expected)
	}
}
