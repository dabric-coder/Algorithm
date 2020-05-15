package HeapSort


// 形成大根堆

func heapInsert(arr []int, index int) {

	for arr[index] > arr[(index-1)/2] {
		swap(arr, index, (index-1)/2)
		index = (index-1)/2
	}
}


func heapify(arr []int, index int, heapSize int) {
	left := index * 2 + 1
	var largest int
	for left < heapSize {
		// 两个孩子节点比较出谁大谁小
		if left + 1 < heapSize && arr[left] < arr[left+1] {
			largest = left + 1
		} else {
			largest = left
		}

		if arr[index] > arr[largest] {
			largest = index
		}

		if largest == index {
			break
		}

		swap(arr, index, largest)
		index = largest
		left = index * 2 + 1

	}

}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}


func HeapSort(arr []int) {
	// 形成大根堆
	for k, _ := range arr {
		heapInsert(arr, k)
	}

	heapSize := len(arr)

	swap(arr, 0, heapSize-1)
	heapSize--

	for heapSize > 0 {
		heapify(arr, 0, heapSize)
		swap(arr, 0, heapSize-1)
		heapSize--
	}
}
