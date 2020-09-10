package main

func main() {

}

func countPrimes(n int) int {
	a := make([]bool, n)
	ret := 0
	// 从 2 开始，1 不是质数
	for i := 2; i < n; i++ {
		// 是质数
		if a[i] == false {
			ret++
			// 将某质数的倍数全部筛掉
			for j := 2 * i; j < n; j += i {
				// 标记为筛掉
				a[j] = true
			}
		}
	}
	return ret
}
