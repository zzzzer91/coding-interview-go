package p0032_最长有效括号

func longestValidParentheses(s string) int {
	res := 0
	lP, rP := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			lP++
		} else {
			rP++
		}
		if lP == rP {
			res = max(res, lP+rP)
		} else if lP < rP {
			lP = 0
			rP = 0
		}
	}
	// 从左到右扫描完毕后，同样的方法从右到左再来一次，
	//因为类似这样的情况 ( ( ( ) ) ，如果从左到右扫描到最后，
	//left = 3，right = 2，期间不会出现 left == right。
	//但是如果从右向左扫描，扫描到倒数第二个位置的时候，
	//就会出现 left = 2，right = 2 ，就会得到一种合法序列。
	lP, rP = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ')' {
			rP++
		} else {
			lP++
		}
		if lP == rP {
			res = max(res, lP+rP)
		} else if rP < lP {
			lP = 0
			rP = 0
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
