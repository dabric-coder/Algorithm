package main

import "fmt"

func partition(slice []int, L, R, num int) [2]int {
	less := -1
	more := R + 1

	cur := 0
	for cur < more {
		if slice[cur] < num {
			swap(slice, cur, less+1)
			less++
			cur++
		} else if slice[cur] > num {
			swap(slice, cur, more-1)
			more--
		} else {
			cur++
		}
	}
	return [2]int{less+1, more-1}
}

func swap(slice []int, x, y int) {
	slice[x] = slice[x] ^ slice[y]
	slice[y] = slice[x] ^ slice[y]
	slice[x] = slice[x] ^ slice[y]
}

func main() {
	s := []int{4,6,5,3,5,8}
	p := partition(s, 0, 5, 5)
	fmt.Println(s)
	fmt.Println(p)
}