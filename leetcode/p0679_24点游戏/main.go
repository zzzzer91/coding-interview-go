package main

import "math"

func judgePoint24(nums []int) bool {
	a := make([]float64, 0, len(nums))
	for _, n := range nums {
		a = append(a, float64(n))
	}
	return f(a)
}

func main() {

}

func f(a []float64) bool {
	if len(a) == 1 {
		return equal(a[0], 24)
	}
	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j < len(a); j++ {
			var temp []float64
			for k := 0; k < len(a); k++ {
				if k != i && k != j {
					temp = append(temp, a[k])
				}
			}
			if f(append(temp, a[i]+a[j])) {
				return true
			}
			if f(append(temp, a[i]-a[j])) {
				return true
			}
			if f(append(temp, a[j]-a[i])) {
				return true
			}
			if f(append(temp, a[i]*a[j])) {
				return true
			}
			if !equal(a[j], 0) && f(append(temp, a[i]/a[j])) {
				return true
			}
			if !equal(a[i], 0) && f(append(temp, a[j]/a[i])) {
				return true
			}
		}
	}
	return false
}

func equal(a, b float64) bool {
	return math.Abs(a-b) < 1e-6
}
