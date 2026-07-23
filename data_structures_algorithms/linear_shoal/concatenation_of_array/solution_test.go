package concatenationofarray

import (
	"reflect"
	"testing"
)

func TestGetConcatenation(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 2, 1},
			expected: []int{1, 2, 1, 1, 2, 1},
		},
		{
			name:     "Example 2",
			nums:     []int{1, 3, 2, 1},
			expected: []int{1, 3, 2, 1, 1, 3, 2, 1},
		},
		{
			name:     "Example 3",
			nums:     []int{},
			expected: nil,
		},
		{
			name:     "Example 4",
			nums:     make([]int, 1001),
			expected: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := getConcatenation(test.nums)

			t.Logf(
				"Input: nums = %v | Result: %v | Expected: %v",
				test.nums, result, test.expected,
			)

			if !reflect.DeepEqual(result, test.expected) {
				t.Error("ERROR - result != expected")
			}
		})
	}
}
