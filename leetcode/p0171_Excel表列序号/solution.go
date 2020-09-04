package p0171_Excel表列序号

func titleToNumber(s string) int {
	num := 0
	for i := 0; i < len(s); i++ {
		num = num*26 + int(s[i]) - 'A' + 1
	}
	return num
}
