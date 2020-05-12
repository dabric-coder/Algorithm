package BubbleSort

import "algorithm/Sort/utils"

func BubbleSort(data utils.Interface) {
	if data.Len() <= 1 {
		return
	}
	for i := data.Len() - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if data.Comparator(j+1, j) {
				data.Swap(j, j+1)
			}
		}
	}
}
