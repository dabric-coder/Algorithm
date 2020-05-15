package MergeSort




func MergeSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	MergeSortGo(arr, 0, len(arr)-1)
}

func MergeSortGo(arr []int, L, R int) {
	if L == R {
		return
	}
	mid := L + (R-L) >> 1
	MergeSortGo(arr, L, mid)
	MergeSortGo(arr, mid+1, R)
	merge(arr, L, R, mid)
}

func merge(arr []int, L, R, mid int) {
	help := make([]int, R-L+1)

	i := 0
	p1 := L
	p2 := mid+1

	for p1 <= mid && p2 <= R {
		if arr[p1] < arr[p2] {
			help[i] = arr[p1]
			p1++
			i++
		} else {
			help[i] = arr[p2]
			p2++
			i++
		}
	}

	for p1 <= mid {
		help[i] = arr[p1]
		p1++
		i++
	}

	for p2 <= R {
		help[i] = arr[p2]
		p2++
		i++
	}

	for j := 0; j < len(help); j++ {
		arr[L+j] = help[j]
	}
}
