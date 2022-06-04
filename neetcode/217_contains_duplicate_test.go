package neetcode_test

import (
	"testing"

	"github.com/intothevoid/leetcode/neetcode"
)

func Test_217_contains_duplicate(t *testing.T) {
	nums := []int{1, 2, 3, 45, 1, 6}
	result := neetcode.ContainsDuplicate(nums)
	if result != true {
		t.Error("Expected true, got", result)
	}
}
