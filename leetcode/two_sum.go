/*
# My Two Sum
# Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
# You may assume that each input would have more than one solution, and you may not use the same element twice.
# You can return the answer in any order.

# Example 1:
# Input: nums = [2,7,11,15], target = 9
# Output: [[0,1]]
# Explanation: Because nums[0] + nums[1] == 9, we return [[0, 1]].

# Example 2:
# Input: nums = [3,3,3], target = 6
# Output: [[0,1],[1,2],[0,2]]

# Example 3:
# Input: nums = [3,3,3,3], target = 6
# Output: [[[0,1],[0,2],[0,3],[1,2],[1,3],[2,3]]

# Example 4:
# Input: nums = [2,9,2,9], target = 11
# Output: [[[0,1],[0,3],[1,2],[1,3]]
*/

package main

import "fmt"

func twoSum(nums []int, target int) [][]int {
	visited := make(map[int][]int)
	var result = [][]int{}

	for i, num := range nums {
		complement := target - num

		if indices, found := visited[complement]; found {
			for _, j := range indices {
				result = append(result, []int{j, i})
			}
		}
		if _, found := visited[num]; !found {
			visited[num] = []int{}
		}
		visited[num] = append(visited[num], i)
	}
	return result
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println("Example 1 Output:", twoSum(nums, target)) // Output: [0 1]

	nums = []int{3, 2, 4}
	target = 6
	fmt.Println("Example 2 Output:", twoSum(nums, target)) // Output: [1 2]

	nums = []int{3, 3, 3, 3}
	target = 6
	fmt.Println("Example 3 Output:", twoSum(nums, target)) // Output: [0 1]
}
