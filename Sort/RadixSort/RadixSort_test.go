package RadixSort

import (
	"fmt"
	"testing"
)

func TestRadixSort(t *testing.T) {
	arr := []int{126,69,539,23,6,89,54,8}
	RadixSort(arr)

	fmt.Println(arr)
}