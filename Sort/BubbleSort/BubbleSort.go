package BubbleSort

import (
	"algorithm/Sort/utils"
)

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

// 如果数组已经完全有序，提前结束终止冒泡排序
func BubbleSort1(arr []int) {

	for i := len(arr)-1; i > 0; i-- {
		sorted := true
		for j := 1; j <= i; j++ {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
				sorted = false
			}
		}

		if sorted {
			break
		}
	}
}


// 如果数组尾部已经局部有序，可以记录最后1次交换的位置，提前终止冒泡排序
func BubbleSort2(arr []int) {

	for end := len(arr)-1; end > 0; end-- {
		sortIndex := 1
		for begin := 1; begin <= end; begin++ {
			if arr[begin] < arr[begin-1] {
				swap(arr, begin, begin-1)
				sortIndex = begin
			}
		}
		end = sortIndex
	}
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

/*
冒泡排序是一个稳定的排序过程
但是，编程的过程中稍有不慎就会把冒泡排序编程不稳定的排序
*/