package main

import (
	"fmt"
)

func mergeSort(slice []int, L, R int) {
	if L == R {
		return
	}
	mid := L + (R - L) >> 1
	mergeSort(slice, L, mid)
	mergeSort(slice, mid+1, R)
	merge(slice, L, R, mid)
}

func merge(slice []int, L, R, mid int) {
	var help []int = make([]int, R-L+1)
	i := 0
	p1 := L
	p2 := mid+1
	for p1 <= mid && p2 <= R {
		if slice[p1] <= slice[p2] {
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
}

func main() {
	s := []int{5,3,4,0,6,1}
	mergeSort(s, 0, 5)
	fmt.Println(s)
}