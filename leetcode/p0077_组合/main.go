package main

func main() {

}

func combine(n int, k int) [][]int {
	a := make([]int, 0, k)
	var res [][]int
	var f func(u int)
	f = func(u int) {
		if len(a) == k {
			temp := make([]int, k)
			copy(temp, a)
			res = append(res, temp)
			return
		}
		for i := u + 1; i <= n; i++ {
			a = append(a, i)
			f(i)
			a = a[:len(a)-1]
		}
	}
	f(0)
	return res
}
