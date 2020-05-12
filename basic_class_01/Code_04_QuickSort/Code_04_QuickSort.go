package main

import (
	"fmt"
	"math/rand"
)

func quickSort(slice []int) {
	if len(slice) < 2 {
		return
	}
	quickSort_recur(slice, 0, len(slice)-1)
}

func quickSort_recur(slice []int, L, R int) {
	if L < R {
		swap(slice, L+ rand.Intn(R-L+1), R)
		p := partition(slice, L, R)
		quickSort_recur(slice, L, p[0]-1)
		quickSort_recur(slice, p[1]+1, R)
	}
}

func partition(slice []int, L, R int) (p *[2]int) {
	less := L - 1
	more := R
	for L < more {
		if slice[L] < slice[R] {
			swap(slice, L, less+1)
			L += 1
			less += 1
		} else if slice[L] > slice[R] {
			swap(slice, L, more-1)
			more -= 1
		} else {
			L += 1
		}
	}
	swap(slice, more, R)
	return &[2]int{less+1, more-1}
}

func swap(slice []int, x, y int) {
	slice[x], slice[y] = slice[y], slice[x]
}

func main() {
	s := []int{3,4,0,1,2,8,7,9}
	quickSort(s)
	fmt.Println(s)
}