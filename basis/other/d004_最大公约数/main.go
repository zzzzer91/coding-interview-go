package main

func gcd(a, b int) int {
	if a%b != 0 {
		return gcd(b, a%b)
	}
	return b
}

func main() {

}
