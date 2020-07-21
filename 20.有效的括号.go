/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start
func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	// write code here
	stack := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		b := s[i]
		switch b {
		case '{', '[', '(':
			stack = append(stack, b)
			continue
		case '}', ']', ')':
			if len(stack) == 0 {
				return false
			}
		default:
			return false
		}
		if b == '}' && stack[len(stack)-1] == '{' {
			stack = stack[:len(stack)-1]
		} else if b == ']' && stack[len(stack)-1] == '[' {
			stack = stack[:len(stack)-1]
		} else if b == ')' && stack[len(stack)-1] == '(' {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}

// @lc code=end
