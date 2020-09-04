package p0066_加一

func plusOne(digits []int) []int {
	ret := make([]int, 1, len(digits)+1)
	ret = append(ret, digits...)
	for i := len(ret) - 1; i > 0; i++ {
		ret[i] += 1
		if ret[i] >= 10 {
			ret[i] -= 10
		} else {
			break
		}
	}
	if ret[0] != 0 {
		return ret
	}
	return ret[1:]
}
