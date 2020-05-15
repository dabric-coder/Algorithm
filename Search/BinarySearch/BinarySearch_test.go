package BinarySearch

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	arr := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		arr[i] = i
	}

	//fmt.Println(arr)

	res := BinarySearchInsert(arr, 1000)

	if res == -1 {
		fmt.Println("未找到")
	} else {
		fmt.Println("找到了", res)
	}
}