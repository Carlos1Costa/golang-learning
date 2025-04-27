package main

import (
	"reflect"
	"strconv"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		// Test cases
		{input: []int{5, 2, 9, 1, 5, 6}, expected: []int{1, 2, 5, 5, 6, 9}},
		{input: []int{1, 2, 3, 4, 5}, expected: []int{1, 2, 3, 4, 5}},                                 // Already sorted
		{input: []int{5, 4, 3, 2, 1}, expected: []int{1, 2, 3, 4, 5}},                                 // Reverse order
		{input: []int{}, expected: []int{}},                                                           // Empty array
		{input: []int{42}, expected: []int{42}},                                                       // Single element
		{input: []int{3, 3, 3}, expected: []int{3, 3, 3}},                                             // All elements the same
		{input: []int{10, -1, 2, 5, 0, 6, 4, -5}, expected: []int{-5, -1, 0, 2, 4, 5, 6, 10}},         // Mixed positive and negative
		{input: []int{-3, -2, -1, -4}, expected: []int{-4, -3, -2, -1}},                               // All negative numbers
		{input: []int{100, 50, 20, 10, 5}, expected: []int{5, 10, 20, 50, 100}},                       // Descending order
		{input: []int{1, 2, 3, 2, 1}, expected: []int{1, 1, 2, 2, 3}},                                 // Duplicates
		{input: []int{0, 0, 0, 0}, expected: []int{0, 0, 0, 0}},                                       // All zeros
		{input: []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 0}, expected: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},   // Mixed order
		{input: []int{1, 3, 2, 3, 1}, expected: []int{1, 1, 2, 3, 3}},                                 // Repeated elements
		{input: []int{1000, 500, 100, 50, 10, 5}, expected: []int{5, 10, 50, 100, 500, 1000}},         // Large numbers
		{input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}}, // Already sorted
		{input: []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}}, // Reverse order
		{input: []int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}, expected: []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}},   // Duplicates
		{input: []int{100, 200, 300, 400, 500}, expected: []int{100, 200, 300, 400, 500}},             // Already sorted large numbers
		{input: []int{500, 400, 300, 200, 100}, expected: []int{100, 200, 300, 400, 500}},             // Reverse order large numbers
		{input: []int{1, 100, 50, 25, 75, 10}, expected: []int{1, 10, 25, 50, 75, 100}},               // Mixed order
	}

	for i, test := range tests {
		t.Run("TestCase"+strconv.Itoa(i), func(t *testing.T) {
			result := InsertionSort(append([]int{}, test.input...)) // Use a copy to avoid modifying the original
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("Test case %d failed: got %v, expected %v", i, result, test.expected)
			}
		})
	}
}
