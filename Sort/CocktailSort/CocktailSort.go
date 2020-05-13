package CocktailSort


/*
鸡尾酒排序又称为双冒泡排序，示例如下：
8 1 4 2 9 5 3

正向冒泡一次，大的数沉底；
反向冒泡一次，小的数飘起来

正向冒泡： 1 4 2 8 5 3 9
反向冒泡： 1 2 4 3 8 5 9

正向冒泡：1 2 3 4 5 8 9


*/

func CocktailSort(slice []int) {
	if len(slice) <= 1 {
		return
	}

	for i := 0; i < len(slice)/2; i++ {  // 每次循环，正向冒泡一次，反向冒泡一次
		left := 0
		right := len(slice) - 1
		for left <= right {
			if slice[left] > slice[left+1] {
				swap(slice, left, left+1)
			}
			left++
			if slice[right] < slice[right-1] {
				swap(slice, right, right-1)
			}
			right--
		}

	}

}


func swap(slice []int, i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}