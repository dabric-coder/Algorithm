package MergeSort

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	arr := []int{11,1,9,2,8,3,7,4,6,5,10}
	MergeSort(arr)
	fmt.Println(arr)
}