package shufflethearray

import (
	"math"
	"reflect"
	"testing"
)

func TestShuffle(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		n        int
		expected []int
	}{
		{
			name:     "Example 1",
			nums:     []int{2, 5, 1, 3, 4, 7},
			n:        3,
			expected: []int{2, 3, 5, 4, 1, 7},
		},
		{
			name:     "Example 2",
			nums:     []int{1, 2, 3, 4, 4, 3, 2, 1},
			n:        4,
			expected: []int{1, 4, 2, 3, 3, 2, 4, 1},
		},
		{
			name:     "Example 3",
			nums:     []int{1, 1, 2, 2},
			n:        2,
			expected: []int{1, 2, 1, 2},
		},
		{
			name:     "Invalid Param",
			nums:     []int{1, 1, 2, 2},
			n:        501,
			expected: nil,
		},
		{
			name:     "Invalid Length",
			nums:     []int{1, 1, 2, 2, 1},
			n:        2,
			expected: nil,
		},
		{
			name:     "Invalid Contents",
			nums:     []int{0, int(math.Pow10(3) + 1)},
			n:        1,
			expected: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := shuffle(test.nums, test.n)

			t.Logf(
				"Input - nums: %v, n: %v",
				test.nums,
				test.n,
			)
			t.Logf(
				"Result: %v | Expected: %v",
				result,
				test.expected,
			)

			if !reflect.DeepEqual(result, test.expected) {
				t.Error("ERROR - result != expected")
			}
		})
	}
}
