/*
Valid palindrome
https://leetcode.com/problems/valid-palindrome/
A palindrome is a word that reads the same backwards and forwards. Given a string "s", determine if it is a palindrome. Note: You can begin by just considering lowercase alphanumeric characters.
Example 1:

Input: s = "A man, a plan, a canal: Panama"
Output: true
Explanation: "amanaplanacanalpanama" is a palindrome.
Example 2:

Input: s = "race a car"
Output: false
Explanation: "raceacar" is not a palindrome.
Example 3:

Input: s = " "
Output: true
Explanation: s is an empty string "" after removing non-alphanumeric characters.
Since an empty string reads the same forward and backward, it is a palindrome.

*/

package main

import (
	"fmt"
	"unicode"
)

func isPalindrome(s string) bool {
	l := 0
	r := len(s) - 1
	var lc, rc rune
	for l < r {
		lc = unicode.ToLower(rune(s[l]))
		rc = unicode.ToLower(rune(s[r]))
		if !unicode.IsLetter(lc) {
			l++
			continue
		}
		if !unicode.IsLetter(rc) {
			r--
			continue
		}
		// fmt.Printf("%c %c\n", lc, rc)
		if lc == rc {
			l++
			r--
		} else {
			return false
		}
	}
	return true
}

func main() {
	s1 := "A man, a plan, a canal: Panama"
	s2 := "race a car"
	s3 := " "
	fmt.Printf("'%s' is Palindrome? %t\n", s1, isPalindrome(s1))
	fmt.Printf("'%s' is Palindrome? %t\n", s2, isPalindrome(s2))
	fmt.Printf("'%s' is Palindrome? %t\n", s3, isPalindrome(s3))
}
