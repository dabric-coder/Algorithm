package main

import "fmt"

func reverseAligment(slice []int, L, R int) (res [][2]int) {
	if L == R {
		return nil
	}
	mid := L + (R - L) >> 1
	res1 := reverseAligment(slice, L, mid)
	res2 := reverseAligment(slice, mid+1, R)
	res3 := merge(slice, L, mid, R)
	res = append(res, res1...)
	res = append(res, res2...)
	res = append(res, res3...)
	return
}

func merge(slice []int, L, mid, R int) (res [][2]int) {
	var help []int = make([]int, R-L+1)
	i := 0
	p1 := L
	p2 := mid + 1

	for p1 <= mid && p2 <= R {
		// merge的过程产生逆序对，并且，如果p1上的数大于p2指定的数，那么(p1, mid]上的数都大于p2指定的数
		if slice[p1] > slice[p2] {
			var r [2]int
			for index := p1; index <= mid; index++ {
				r[0] = slice[index]
				r[1] = slice[p2]
				res = append(res, r)
			}
		}
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
		p1++
		i++
	}

	for p2 <= R {
		help[i] = slice[p2]
		p2++
		i++
	}

	for j := 0; j < len(help); j++ {
		slice[L+j] = help[j]
	}
	return
}

func main() {
	s := []int{5,4,3,2,1}
	res := reverseAligment(s, 0, 4)
	fmt.Println(res)
}