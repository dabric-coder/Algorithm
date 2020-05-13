package OddEvenSort

import (
	"fmt"
	"testing"
)

func TestOddEvenSort(t *testing.T) {
	arr := []int{1,9,2,8,3,7,4,6,5,10}
	OddEvenSort(arr)
	fmt.Println(arr)
}