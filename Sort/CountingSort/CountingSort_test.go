package CountingSort

import (
	"fmt"
	"testing"
)

func TestGetMax(t *testing.T) {
	arr := []int{7,3,5,8,6,7,4,5}
	CountingSort(arr)
	fmt.Println(arr)
}