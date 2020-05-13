package ShellSort


/*
希尔排序是插入排序的按步长缩减的排序，排序过程如下示例：
1,9,2,8,3,7,4,6,5,10

步长为4： 1         7
步长为4：   9         4      			按插入排序的方式进行排序 4 9
步长为4：     2         6
步长为4：       8         5              按插入排序的方式进行排序 5 8
步长为4： 1 4 2 5 3 7 9 6 8 10
1 9 2 8 3 7 4 6 5 10
步长为3： 1       3       5
步长为3：   9       7       10			按插入排序的方式进行排序 7 9 10
步长为3：     2       4
步长为3： 1 9 2 8 3 7 4 6 5 10

*/

func ShellSortStep(slice []int, start, gap int) {
	// 插入排序
	for i := start+gap; i < len(slice); i += gap {
		for j := i - gap; j >=0 && slice[j] > slice[j+gap]; j -= gap {
			swap(slice, j, j+gap)
		}
	}
}

func ShellSort(slice []int) {
	if len(slice) <= 1 {
		return
	}

	gap := len(slice)/2

	for gap > 0 {
		for i := 0; i < gap; i++ {
			ShellSortStep(slice, i, gap)
		}
		//gap /= 2
		gap--
	}
}


func swap(slice []int, i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}