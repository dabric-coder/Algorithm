package BubbleSort

import (
	"algorithm/Sort/utils"
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{10, 1, 3, 9, 6, 8, 5, 0}
	a := utils.IntSlice(arr)
	BubbleSort2(a)
	fmt.Println(a)
}