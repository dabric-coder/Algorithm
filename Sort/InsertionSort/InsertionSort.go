package InsertionSort

import (
	"algorithm/Sort/utils"
)

func InsertionSort(data utils.Interface) {
	if data.Len() <= 1 {
		return
	}
	for i := 1; i < data.Len(); i++ {
		for j := i - 1; j >= 0 && data.Comparator(j, j + 1); j-- {
			data.Swap(j, j+1)
		}
	}
}




