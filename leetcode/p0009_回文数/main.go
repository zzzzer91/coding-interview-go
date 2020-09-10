package main

func main() {

}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	temp := x
	p := 0
	for temp != 0 {
		p = p*10 + temp%10
		temp /= 10
	}
	return p == x
}
