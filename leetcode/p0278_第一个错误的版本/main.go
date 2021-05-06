// https://leetcode-cn.com/problems/first-bad-version/

package main

func main() {

}

func firstBadVersion(n int) int {
	l, r := 0, n-1
	for l < r {
		m := l + (r-l)>>1
		if isBadVersion(m + 1) {
			r = m
		} else {
			l = m + 1
		}
	}
	return l + 1
}

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 */
// unknown
func isBadVersion(n int) bool {
	return false
}
