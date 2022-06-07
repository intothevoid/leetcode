package neetcode

import (
	"fmt"
	"regexp"
)

func isPalindrome(s string) bool {
	// strip all non numeric keywords
	r, _ := regexp.Compile(`[\w]+`)
	s2 := r.FindString(s)

	fmt.Println("String is : " + s2)

	// reverse string
	// var s2_rev string = ""

	return true
}
