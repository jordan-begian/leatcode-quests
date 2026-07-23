/* Package shufflethearray
* Given the array nums consisting of 2n elements in the form [x1,x2,...,xn,y1,y2,...,yn].
* Return the array in the form [x1,y1,x2,y2,...,xn,yn].
*
* Example 1:
* - Input: nums = [2,5,1,3,4,7], n = 3
* - Output: [2,3,5,4,1,7]
* - Explanation: Since x1=2, x2=5, x3=1, y1=3, y2=4, y3=7 then the answer is [2,3,5,4,1,7].
*
* Example 2:
* - Input: nums = [1,2,3,4,4,3,2,1], n = 4
* - Output: [1,4,2,3,3,2,4,1]
*
* Example 3:
* - Input: nums = [1,1,2,2], n = 2
* - Output: [1,2,1,2]
*
* Constraints:
* - 1 <= n <= 500
* - nums.length == 2n
* - 1 <= nums[i] <= 10^3
 */
package shufflethearray

import (
	"math"
	"slices"
)

func shuffle(nums []int, n int) []int {
	if !isValid(nums, n) {
		return nil
	}

	a := nums[:n]
	b := nums[n:]
	answer := make([]int, 0, 2*n)

	for index := range n {
		answer = append(answer, a[index], b[index])
	}

	return answer
}

func isValid(nums []int, n int) bool {
	invalidParam := n < 1 || n > 500
	invalidLength := len(nums) != (2 * n)
	invalidContents := slices.ContainsFunc(nums, func(num int) bool {
		return num < 1 || num > int(math.Pow10(3))
	})

	if invalidParam || invalidLength || invalidContents {
		return false
	}

	return true
}
