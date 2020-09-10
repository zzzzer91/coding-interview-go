package main

import "math"

func main() {

}

// int 在 32 位平台是 4 字节，会溢出
func reverse(x int) int {
	ret := int64(0)
	temp := int64(x)
	for temp != 0 {
		ret = ret*10 + temp%10
		if ret < math.MinInt32 || ret > math.MaxInt32 {
			return 0
		}
		temp /= 10
	}
	return int(ret)
}
