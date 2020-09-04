package p0190_颠倒二进制位

func reverseBits(num uint32) uint32 {
	var leftMask, rightMask uint32 = 1 << 31, 1
	for i := 0; i < 16; i++ {
		leftBit, rightBit := num&leftMask, num&rightMask
		if leftBit == 0 && rightBit != 0 {
			num |= leftMask
			num ^= rightMask
		}
		if leftBit != 0 && rightBit == 0 {
			num ^= leftMask
			num |= rightMask
		}
		leftMask >>= 1
		rightMask <<= 1
	}
	return num
}

// better
func reverseBits2(num uint32) uint32 {
	ret := uint32(0)
	i, j := uint32(0), uint32(31)
	for i < j {
		if (num & (1 << i)) != 0 {
			ret = ret | (1 << j)
		}
		if (num & (1 << j)) != 0 {
			ret = ret | (1 << i)
		}
		i++
		j--
	}
	return ret
}
