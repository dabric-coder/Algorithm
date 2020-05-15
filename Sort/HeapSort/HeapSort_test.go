package HeapSort

import (
	"fmt"
	"testing"
)

func TestHeapInsert(t *testing.T) {
	arr := []int{11,1,9,2,8,3,7,4,6,5,10}

	for k, _ := range arr {
		heapInsert(arr, k)
	}
	arr[1] = 0
	fmt.Println(arr)

	heapify(arr, 1, len(arr))
	fmt.Println(arr)
	HeapSort(arr)
	fmt.Println(arr)
}