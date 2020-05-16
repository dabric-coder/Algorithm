package RadixSort


func RadixSort(arr []int) {

	/*
	实现思路：
	一个数组中最大的数决定要使用几次计数排序对每位上的数进行排序
	一个数如何计算其个位、百位、千位上的数。比如586
	个位数： 586 / 1 % 10
	百位数： 586 / 10 % 10
	千位数： 586 / 100 % 10

	*/
	maxIndex := 0

	for i := 0; i < len(arr); i++ {
		if arr[i] > arr[maxIndex] {
			maxIndex = i
		}
	}
	max := arr[maxIndex]

	for i := 1; i <= max; i *= 10 {
		CountingSort(arr, i)
	}


}

func CountingSort(arr []int, divider int) {

	counts := make([]int, 10)

	for i := 0; i < len(arr); i++ {
		counts[arr[i] / divider % 10]++
	}

	for i := 1; i < len(counts); i++ {
		counts[i] = counts[i] + counts[i-1]
	}

	help := make([]int, len(arr))

	for i := len(arr)-1; i >= 0; i-- {
		help[counts[arr[i] / divider % 10]-1] = arr[i]
		counts[arr[i] /divider % 10]--
	}

	for i :=0; i < len(arr); i++ {
		arr[i] = help[i]
	}
}

