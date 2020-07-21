
/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 */

// @lc code=start
import (
	"strconv"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	m := strconv.Itoa(x)
	n := Reverse(m)
	if m == n {
		return true
	} else {
		return false
	}
}

func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

// @lc code=end

