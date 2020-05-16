package CountingSort

import "fmt"

func GetMax(arr []int) int {
	maxIndex := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[maxIndex] {
			maxIndex = i
		}
	}
	return arr[maxIndex]
}

func CountingSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	arrMax := GetMax(arr)
	counts := make([]int, arrMax+1)

	// 计数
	for i := 0; i < len(arr); i++ {
		counts[arr[i]]++
	}
	//fmt.Println(counts)
	// 复原数组
	index := 0
	for i := 0; i < len(counts); i++ {
		for ; counts[i] > 0; counts[i]-- {
			arr[index] = i
			index++
		}
	}
}

func CountingSortPlus(arr []int) {
	// 求出数组中的最大最小值
	minIndex := 0
	maxIndex := 0

	for i := 0; i < len(arr); i++ {
		if arr[i] > arr[maxIndex] {
			maxIndex = i
		}

		if arr[i] < arr[minIndex] {
			minIndex = i
		}
	}
	max := arr[maxIndex]
	min := arr[minIndex]

	// 构建counts数组，存储arr数中出现的次数
	counts := make([]int, max-min+1)

	// 统计arr中每个数出现的次数，arr中的元素和counts索引对应关系是index = k - min
	for i := 0; i < len(arr); i++ {
		counts[arr[i]-min]++
	}

	// counts数组中存储的每个次数，累加前面的次数
	for i := 1; i < len(counts); i++ {
		counts[i] = counts[i] + counts[i-1]
	}
	fmt.Println(counts)

	// 遍历arr数组，根据counts数组存储的信息，复原数组。
	//其中，arr中元素在有序序列对应的索引为counts[k-min]-p
	// 创建一个临时数组
	help := make([]int, len(arr))
	for i := len(arr)-1; i >= 0; i-- {
		// 从后往前遍历arr数组，进行复原，保证排序的稳定性
		help[counts[arr[i]-min]-1] = arr[i]
		counts[arr[i]-min]--
	}

	// 将help数组中的数组复制到原数组中
	for i := 0; i < len(arr); i++ {
		arr[i] = help[i]
	}

}