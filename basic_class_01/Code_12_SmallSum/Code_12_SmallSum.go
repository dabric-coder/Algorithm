package main

import "fmt"

func smallSum(slice []int, L, R int) (res int) {
	if L == R {
		return 0
	}
	mid := L + (R - L) >> 1
	return smallSum(slice, L, mid) + smallSum(slice, mid + 1, R) + merge(slice, L, mid, R)
}

func merge(slice []int, L, mid, R int) (res int) {
	var help []int = make([]int, R-L+1)
	i := 0
	p1 := L
	p2 := mid + 1
	for p1 <= mid && p2 <= R {

		if slice[p1] < slice[p2] {
			// 产生小和，并且，p1小于(p2, R]上的数，于是产生的小和为p1 *(R-p2+1)
			res = res + slice[p1] *(R-p2+1)
		} else {
			res += 0
		}

		if slice[p1] < slice[p2] {
			help[i] = slice[p1]
			p1++
		} else {
			help[i] = slice[p2]
			p2++
		}
		i++
	}

	for p1 <= mid {
		help[i] = slice[p1]
		i++
		p1++
	}

	for p2 <= R {
		help[i] = slice[p2]
		i++
		p2++
	}

	for j := 0; j < len(help); j++ {
		slice[L+j] = help[j]
	}
	return
}


func main() {
	s := []int{1,3,4,2,5}
	res := smallSum(s, 0, 4)
	fmt.Println(res)
}