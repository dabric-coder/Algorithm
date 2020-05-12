package main

import "fmt"

func heapInsert(slice []int, index int) {

	for slice[index] > slice[(index-1)/2] {
		swap(slice, index, (index-1)/2)
		index = (index-1)/2     // 调整节点，继续往上比较
	}
}


func heapify(slice []int, index, heapSize int)  {
	left := index * 2 + 1
	var largest int
	for left < heapSize {
		// 两个孩子之间比较出谁大谁小
		if left + 1 < heapSize && slice[left] < slice[left+1] {
			largest = left + 1
		} else {
			largest = left
		}

		if slice[index] > slice[largest] {
			largest = index
		} else {
			largest = largest
		}

		if largest == index {  // 当前位置的值变小后和左右孩子比较后，还是自身的值大时，此时该值不用下沉；你和你孩子的最大值是自己，直接break
			break
		}
		swap(slice, largest, index)
		index = largest        // 下沉
		left = index * 2 + 1
	}
}

func swap(slice []int, x, y int) {
	slice[x], slice[y] = slice[y], slice[x]
}

func heapSort(slice []int) {
	// 形成大根堆

	// 交换数组中的第一个数和最后一个数，heapSize--

	// 重新调整大根堆，重复上面的过程，直到heapSize==0

	for k, _ := range slice {
		heapInsert(slice, k)
	}

	fmt.Println(slice)

	heapSize := len(slice)

	swap(slice,0, heapSize-1)
	heapSize--

	for heapSize > 0 {
		heapify(slice, 0, heapSize)
		swap(slice,0, heapSize-1)
		heapSize--
	}
}

func main() {
	s := []int{3,4,5,0,1,6}
	heapSort(s)
	fmt.Println(s)
}