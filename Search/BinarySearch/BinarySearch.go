package BinarySearch

import "fmt"

func BinarySearch(arr []int, num int) int {
	left := 0
	right := len(arr) - 1
	i := 0
	for left < right {
		i++
		fmt.Println(i)
		mid := (left + right) >> 1

		if arr[mid] < num {
			left = mid + 1
		} else if arr[mid] > num {
			right = mid - 1
		} else {
			return mid
		}
	}

	return -1
}

func BinarySearchInsert(arr []int, num int) int {
	low := 0
	high := len(arr) - 1
	i := 0  // 统计查找的次数
	for low < high {
		i++
		fmt.Println(i)
		mid := int(float64(low) + float64(num - arr[low])/float64(arr[high] - num)*float64(high - low))

		if mid < 0 || mid >= len(arr) {
			return -1
		}

		if arr[mid] > num {
			high = mid - 1
		} else if arr[mid] < num {
			low = mid + 1
		} else {
			return mid
		}
	}
	return -1
}