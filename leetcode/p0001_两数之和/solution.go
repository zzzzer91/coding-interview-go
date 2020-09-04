package p0001_两数之和

func twoSum(nums []int, target int) []int {
	dict := make(map[int]int)
	for i, n := range nums {
		if v, ok := dict[n]; ok {
			return []int{v, i}
		}
		dict[target-n] = i
	}
	return []int{-1, -1}
}
