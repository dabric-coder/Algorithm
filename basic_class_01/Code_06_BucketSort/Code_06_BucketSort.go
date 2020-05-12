package main

import (
	"fmt"
	"math"
)

// only for 0-200 value
func bucketSort(slice []int) {

	// 找到数组中的最大值
	max := 0
	for _, value := range slice {
		max = int(math.Max(float64(max), float64(value)))
	}

	// 创建一个桶
	var bucket = make([]int, max+1)

	// 遍历slice
	for i := 0; i < len(slice); i++ {
		bucket[slice[i]]++
	}

	i := 0
	for j := 0; j < len(bucket); j++ {
		for bucket[j] > 0 {
			slice[i] = j
			i++
			bucket[j]--
		}
	}
}


func main() {
	s := []int{7,7,8,1,2,3,0,0,10}
	bucketSort(s)
	fmt.Println(s)
}