package QuickSort

import "math/rand"

func QuickSort(slice []int) {
	if len(slice) <= 1 {
		return
	}
	quickSort_recur(slice, 0, len(slice)-1)
}

func quickSort_recur(slice []int, L, R int) {
	if L < R {
		swap(slice, L + rand.Intn(R-L+1), R)
		p := partition(slice, L, R)
		quickSort_recur(slice, L, p[0] - 1)
		quickSort_recur(slice, p[1]+1, R)
	}
}

func partition(slice []int, L, R int) (p *[2]int) {
	less := L-1
	more := R+1

	cur := L
	for cur < more {
		if slice[cur] < slice[R] {
			swap(slice, less+1, cur)
			less++
			cur++
		} else if slice[cur] > slice[R] {
			swap(slice, cur, more-1)
			more--
		} else {
			cur++
		}
	}
	return &[2]int{less+1, more-1}  // 返回等于区域的两个边界
}

func swap(slice []int, i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}



