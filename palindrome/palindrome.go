package palindrome

import (
	"strconv"
)

// IsPalindrome - checks if an integer is a palindrome
// This will only handle integers up to length 19 due to limitations on the int64 type
func IsPalindrome(i int64) (r bool) {
	// is negative
	if i < 0 {
		return false
	}

	s := strconv.FormatInt(i, 10)
	c := len(s)
	isEven := c%2 == 0

	half := c / 2
	var last string
	first := s[:half]
	if isEven {
		last = s[half:]
	} else {
		last = s[half+1:]
	}

	return first == reverseString(last)
}

func reverseString(s string) (result string) {
	for _,v := range s {
	  result = string(v) + result
	}
	return 
  }
