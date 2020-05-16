package BucketSort

import (
	"fmt"
	"testing"
)

func TestBucketSort(t *testing.T) {
	arr := []int{1,2,3,4,4,3,2,2,3,1}
	BucketSortPlus(arr)
	fmt.Println(arr)
}


