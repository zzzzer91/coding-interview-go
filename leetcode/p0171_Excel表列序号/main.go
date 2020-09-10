package main

func main() {

}

func titleToNumber(s string) int {
	num := 0
	for i := 0; i < len(s); i++ {
		num = num*26 + int(s[i]) - 'A' + 1
	}
	return num
}
