package QuickSort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr := []int{4,6,5,3,5,8}

	QuickSort(arr)

	fmt.Println(arr)
}
