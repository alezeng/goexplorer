/*
First unique character in a string
https://leetcode.com/problems/first-unique-character-in-a-string/
Given a string "s", find the first character in the string that does not repeat itself and return its index.
Example 1:
Input: s="tiktok"
Output: 1
*/

package main

import "fmt"

func first_unique(input string) int {
	charCounter := make(map[rune]int)

	for _, s := range input {
		charCounter[s] += 1
	}

	for i, s := range input {
		if charCounter[s] == 1 {
			return i
		}
	}
	return -1
}

func main() {
	input := "tiktok"
	fmt.Printf("First unique string index of %s is %d\n", input, first_unique(input))
}
