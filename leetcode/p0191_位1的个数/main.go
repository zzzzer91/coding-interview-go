package main

func main() {

}

func hammingWeight(num uint32) int {
	count := 0
	for num > 0 {
		count += int(num & 1)
		num >>= 1
	}
	return count
}

func hammingWeight2(num uint32) int {
	bitCount := 0
	var mask uint32 = 1
	for i := 0; i < 32; i++ {
		if num&mask != 0 {
			bitCount++
		}
		mask <<= 1
	}
	return bitCount
}

func hammingWeight3(num uint32) int {
	count := 0
	for num != 0 {
		// 这个技巧能清0右边最后一个bit
		num &= num - 1
		count++
	}
	return count
}
