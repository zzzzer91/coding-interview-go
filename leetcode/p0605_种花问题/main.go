// https://leetcode-cn.com/problems/can-place-flowers/

package main

func main() {

}

func canPlaceFlowers(flowerbed []int, n int) bool {
	if len(flowerbed) == 0 {
		return false
	}
	if n == 0 {
		return true
	}
	for i := 0; i < len(flowerbed); i++ {
		cond := flowerbed[i] == 0
		if i > 0 {
			cond = cond && flowerbed[i-1] == 0
		}
		if i < len(flowerbed)-1 {
			cond = cond && flowerbed[i+1] == 0
		}
		if cond {
			flowerbed[i] = 1
			n--
			if n == 0 {
				return true
			}
		}
	}
	return n == 0
}
