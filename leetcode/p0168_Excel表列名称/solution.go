package main

func main() {

}

func convertToTitle(n int) string {
	var result []byte
	for n != 0 {
		result = append(result, byte((n-1)%26+65))
		n = (n - 1) / 26
	}
	reverse(result)
	return string(result)
}

func reverse(b []byte) {
	i, j := 0, len(b)-1
	for i < j {
		b[i], b[j] = b[j], b[i]
		i++
		j--
	}
}
