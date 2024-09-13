package binarysearch

import (
	"math/rand"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	t.Run("Integer Tests", func(t *testing.T) {
		testIntegerSearch(t)
	})

	t.Run("Float Tests", func(t *testing.T) {
		testFloatSearch(t)
	})

	t.Run("String Tests", func(t *testing.T) {
		testStringSearch(t)
	})

	t.Run("Large Array Test", func(t *testing.T) {
		testLargeArray(t)
	})
}

func testIntegerSearch(t *testing.T) {
	ints := []int{-50, -20, 0, 3, 5, 9, 12, 15, 50, 100}

	testCases := []struct {
		name     string
		target   int
		expected int
	}{
		{"Found at beginning", -50, 0},
		{"Found at end", 100, 9},
		{"Found in middle", 9, 5},
		{"Not found - too small", -100, -1},
		{"Not found - too large", 200, -1},
		{"Not found - in between", 10, -1},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := BinarySearch(ints, tc.target)
			if result != tc.expected {
				t.Errorf("Expected %d, got %d", tc.expected, result)
			}
		})
	}
}

func testFloatSearch(t *testing.T) {
	floats := []float64{-5.5, -1.1, 0.0, 2.2, 3.3, 4.4, 5.5}

	testCases := []struct {
		name     string
		target   float64
		expected int
	}{
		{"Found at beginning", -5.5, 0},
		{"Found at end", 5.5, 6},
		{"Found in middle", 2.2, 3},
		{"Not found", 1.1, -1},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := BinarySearch(floats, tc.target)
			if result != tc.expected {
				t.Errorf("Expected %d, got %d", tc.expected, result)
			}
		})
	}
}

func testStringSearch(t *testing.T) {
	strings := []string{"apple", "banana", "cherry", "date", "elderberry"}

	testCases := []struct {
		name     string
		target   string
		expected int
	}{
		{"Found at beginning", "apple", 0},
		{"Found at end", "elderberry", 4},
		{"Found in middle", "cherry", 2},
		{"Not found", "fig", -1},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := BinarySearch(strings, tc.target)
			if result != tc.expected {
				t.Errorf("Expected %d, got %d", tc.expected, result)
			}
		})
	}
}

func testLargeArray(t *testing.T) {
	size := 1000000
	nums := make([]int, size)
	for i := 0; i < size; i++ {
		nums[i] = i
	}

	testCases := []struct {
		name     string
		target   int
		expected int
	}{
		{"Found at beginning", 0, 0},
		{"Found at end", 999999, 999999},
		{"Found in middle", 500000, 500000},
		{"Not found - too small", -1, -1},
		{"Not found - too large", 1000000, -1},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := BinarySearch(nums, tc.target)
			if result != tc.expected {
				t.Errorf("Expected %d, got %d", tc.expected, result)
			}
		})
	}

	t.Run("Random Searches", func(t *testing.T) {
		for i := 0; i < 1000; i++ {
			target := rand.Intn(size)
			result := BinarySearch(nums, target)
			if result != target {
				t.Errorf("Expected %d, got %d", target, result)
			}
		}
	})
}
