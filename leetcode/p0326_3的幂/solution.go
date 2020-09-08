package main

func main() {

}

func isPowerOfThree(n int) bool {
	// 3 ^ 19 = 1162261467
	// 是 int32 能装下的最大的 3 的指数
	return n > 0 && (1162261467%n == 0)
}
