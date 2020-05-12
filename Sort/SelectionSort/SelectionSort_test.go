package SelectionSort

import (
	"algorithm/Sort/utils"
	"fmt"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	arr := []int{11, 1, 3, 9, 6, 8, 5, 0, 0, 2}

	a := utils.IntSlice(arr)

	SelectionSort(a)

	fmt.Println(a)
}

func TestSelectionSort2(t *testing.T) {
	arr := []string{"tom", "alex", "ajax", "jerry", "alice"}
	fmt.Println(arr)

	a := utils.StringSlice(arr)
	SelectionSort(a)
	fmt.Println(a)
}