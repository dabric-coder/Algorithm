package main

import (
	"fmt"
	"math"
)

func maxGap(slice []int) int {
	if slice == nil || len(slice) < 2 {
		return 0
	}

	len := len(slice)
	min := 0
	max := 0

	for _, value := range slice {
		min = int(math.Min(float64(min), float64(value)))
		max = int(math.Max(float64(min), float64(value)))
	}

	if max == min  {
		return 0
	}

	var hasNum = make([]bool, len+1)  // 桶是否为空桶
	var maxs = make([]int, len+1)   // 存每个桶中的最大值
	var mins = make([]int, len+1)	  // 存每个桶中的最小值

	bid := 0

	// 设置每个桶中的最大最小值
	for i := 0; i < len; i++ {
		// 当前数去哪个桶，哪个桶的对应的信息要改写，或者更新或者改写
		bid = bucket(slice[i], len, min, max)  // 得出数组中的每个数对应放的桶id
		if hasNum[bid] {  // false 表示不为空桶
			mins[bid] = int(math.Min(float64(mins[bid]), float64(slice[i])))
		} else {
			mins[bid] = slice[i]
		}

		if hasNum[bid] {
			maxs[bid] = int(math.Max(float64(maxs[bid]), float64(slice[i])))
		} else {
			maxs[bid] = slice[i]
		}
		hasNum[bid] = true
	}

	res := 0
	lastMax := maxs[0]
	for i := 1; i <= len; i++ {
		if hasNum[i] {
			res = int(math.Max(float64(res), float64(mins[i] - lastMax)))
			lastMax = maxs[i]
		}
	}

	return res
}

// 根据数组中的最大值最小值以及数组的长度，怎么判断一个数属于哪个桶
func bucket(num, len, min, max int) int {
	return  (num - min) * len / (max - min)
}

func main() {
	s := []int{1,3,4,5,7,10,20}
	fmt.Println(maxGap(s))

}