package p0067_二进制求和

func addBinary(a string, b string) string {
	// a 放置长字符串
	if len(a) < len(b) {
		a, b = b, a
	}
	res := make([]byte, len(a)+1)
	val := byte(0)
	for i, j, k := len(a)-1, len(b)-1, len(res)-1; i >= 0 || j >= 0; i, j, k = i-1, j-1, k-1 {
		if i >= 0 {
			val += a[i] - '0'
		}
		if j >= 0 {
			val += b[j] - '0'
		}
		res[k] = val%2 + '0'
		val /= 2
	}
	if val != 0 {
		res[0] = val + '0'
		return string(res)
	}
	return string(res[1:])
}
