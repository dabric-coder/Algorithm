package BucketSort

func SelectSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		minIndex := i
		for j := 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				arr[j], arr[minIndex] = arr[minIndex], arr[j]
			}
		}
	}
}

func BucketSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	// 求出数组中最大的数
	//maxIndex := 0
	//for i := 0; i < len(arr); i++ {
	//	if arr[i] > arr[maxIndex] {
	//		maxIndex = i
	//	}
	//}
	//
	//max := arr[maxIndex]

	num := 4  // 生成四个桶
	buckets := make([][]int, num)

	// 将每个数放入对应的桶中
	for i := 0; i < len(arr); i++ {
		buckets[arr[i]-1] = append(buckets[arr[i]-1], arr[i])
	}

	// 对每个桶进行排序
	for i := 0; i < len(buckets); i++ {
		SelectSort(buckets[i])
	}

	// 从左向右依次遍历非空桶中的元素，将桶中的元素赋值到原数组中对应的位置
	tmppose := 0  // 记录将每个桶中的数据复制到arr后的位置
	for i := 0; i < len(buckets); i++ {
		bucketLen := len(buckets[i])

		for j := 0; j < bucketLen && bucketLen > 0; j++ {
			arr[j+tmppose] = buckets[i][j]
		}
		tmppose += bucketLen

		//if bucketLen > 0 {
		//	copy(arr[tmppose:], buckets[i])
		//	tmppose += bucketLen
		//}
	}
}

func BucketSortPlus(arr []int) {
	if len(arr) <= 1 {
		return
	}

	num := len(arr)

	// 求出数组中的最大值
	maxIndex := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > arr[maxIndex] {
			maxIndex = i
		}
	}
	max := arr[maxIndex]

	// 创建桶
	buckets := make([][]int, num)

	// 将数组中对应的数据放入到相应的桶中
	index := 0
	for i := 0; i < len(arr); i++ {
		index = arr[i]*(num-1)/max  // 木桶自动分配公式
		buckets[index] = append(buckets[index], arr[i])
	}

	tmppose := 0
	for i := 0; i < len(buckets); i++ {
		bucketLen := len(buckets[i])

		if bucketLen > 0 {
			// 对每个非空桶进行排序
			SelectSort(buckets[i])

			copy(arr[tmppose:], buckets[i])
			tmppose += bucketLen
		}
	}
}
