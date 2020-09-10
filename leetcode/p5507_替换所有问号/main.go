package main

func main() {

}

func modifyString(s string) string {
	res := []byte(s)
	for i := range res {
		if s[i] != '?' {
			continue
		}
		var c byte = 'a'
		for (i > 0 && res[i-1] == c) || (i+1 < len(s) && res[i+1] == c) {
			c++
		}
		res[i] = c
	}
	return string(res)
}
