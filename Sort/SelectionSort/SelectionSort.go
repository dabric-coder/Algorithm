package SelectionSort

import (
	"algorithm/Sort/utils"
)

func SelectionSort(data utils.Interface) {
	if data.Len() <= 1 {
		return
	}
	for i := 0; i < data.Len(); i++ {
		minIndex := i
		for j := i + 1; j < data.Len(); j++ {
			if data.Comparator(minIndex, j) {
				data.Swap(minIndex, j)
			}
		}
	}
}


func SelectSortX(arr []int) {
	for i := 0; i < len(arr); i++ {
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				swap(arr, j, minIndex)
			}
		}
	}
}

func swap(arr []int, x, y int) {
	arr[x] = arr[x] ^ arr[y]
	arr[y] = arr[x] ^ arr[y]
	arr[x] = arr[x] ^ arr[y]
}