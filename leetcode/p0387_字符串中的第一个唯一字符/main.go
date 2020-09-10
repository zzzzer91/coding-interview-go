package main

func main() {

}

func firstUniqChar(s string) int {
	m := make([]int, 26)
	for _, c := range s {
		m[c-'a']++
	}
	for i, c := range s {
		if m[c-'a'] == 1 {
			return i
		}
	}
	return -1
}
