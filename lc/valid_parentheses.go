/*
https://leetcode.com/problems/valid-parentheses/

20. Valid Parentheses
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.


Example 1:

Input: s = "()"
Output: true
Example 2:

Input: s = "()[]{}"
Output: true
Example 3:

Input: s = "(]"
Output: false


Constraints:

1 <= s.length <= 104
s consists of parentheses only '()[]{}'.
*/

package main

import "fmt"

func isValid(s string) bool {
	var leftover []string
	pairs := map[string]string{
		")": "(",
		"}": "{",
		"]": "[",
	}

	for _, v := range s {
		c := string(v)
		switch c {
		case "(", "{", "[":
			leftover = append(leftover, c)
		case ")", "}", "]":
			idx := len(leftover) - 1
			if idx < 0 || leftover[idx] != pairs[c] {
				return false
			} else { //matched, remove the last one
				leftover = leftover[:idx]
			}
		}
	}
	if len(leftover) > 0 {
		return false
	}
	return true
}

func main() {
	s := "()[]{}"
	fmt.Printf("%s is %t\n", s, isValid(s))
}
