package neetcode

import (
	"fmt"
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	// strip all non numeric keywords
	r, _ := regexp.Compile(`[a-z0-9]+`)
	sl := r.FindAllString(strings.ToLower(s), -1)

	s2 := ""
	for _, v := range sl {
		s2 += v
	}
	fmt.Println(s2)

	// reverse string
	var s2_rev []rune
	for i := len(s2); i > 0; i-- {
		s2_rev = append(s2_rev, rune(s2[i-1]))

	}

	// check
	return s2 == string(s2_rev)
}
