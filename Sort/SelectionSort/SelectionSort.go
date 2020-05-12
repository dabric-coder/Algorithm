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
