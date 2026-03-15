package main

import "fmt"

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func main() {
	fmt.Printf("abba is Palindrome: %v\n", isPalindrome("abba"))
	fmt.Printf("arora is Palindrome: %v\n", isPalindrome("arora"))
	fmt.Printf("abc is Palindrome: %v\n", isPalindrome("abc"))
}
