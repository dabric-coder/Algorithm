package CountingSort

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