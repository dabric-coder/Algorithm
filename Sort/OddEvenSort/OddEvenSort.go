package OddEvenSort


/*
9 2 1 6 0 7

奇数位： 2 9 1 6 0 7
偶数位： 2 1 9 0 6 7

奇数位： 1 2 0 9 6 7
偶数位： 1 0 2 6 9 7

奇数位： 0 1 2 6 7 9
偶数位： 0 1 2 6 7 9

*/


func OddEvenSort(slice []int) {
	isSorted := false // 用来标识奇数位和偶数位是否进行了交换

	for ;isSorted == false; {
		isSorted = true  // 进入循环后认为此时的奇数位和偶数位没有交换

		for i := 1; i < len(slice)-1; i += 2{
			// 奇数位
			if slice[i] > slice[i+1] {
				// 此时表明奇数位需要排序
				swap(slice, i, i+1)
				isSorted = false
			}
		}

		for i := 0; i < len(slice)-1; i += 2 {
			// 偶数位
			if slice[i] > slice[i+1] {
				// 此时表明奇数位需要排序
				swap(slice, i, i+1)
				isSorted = false
			}
		}
	}

}


func swap(slice []int, i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}