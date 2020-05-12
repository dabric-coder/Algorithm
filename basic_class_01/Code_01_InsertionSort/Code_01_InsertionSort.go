package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func insertionSort(slice []int) {
	for i := 1; i < len(slice); i++ {
		for j := i - 1; j >= 0 && slice[j] > slice[j+1]; j-- {
			swap(slice, j, j+1)
		}
	}
}

func swap(slice []int, x, y int) {
	slice[x] = slice[x] ^ slice[y]
	slice[y] = slice[x] ^ slice[y]
	slice[x] = slice[x] ^ slice[y]
}

func comparator(slice []int) {
	sort.Ints(slice)
}

func generateRandomArray(maxSize, maxValue int) (slice []int) {
	for i := 0; i < rand.Intn(maxSize+1); i++ {
		slice = append(slice, rand.Intn(maxValue+1)-rand.Intn(maxValue))
	}
	return
}

func copyArray(slice []int) []int {
	if slice == nil {
		return nil
	}

	var res []int
	for i := 0; i < len(slice); i++ {
		res = append(res, slice[i])
	}
	return res
}

func isEqual(slice1, slice2 []int) bool {
	if (slice1 == nil && slice2 != nil) || (slice1 != nil && slice2 == nil) {
		return false
	}

	if slice1 == nil && slice2 == nil {
		return true
	}

	if len(slice1) != len(slice2) {
		return false
	}

	for i := 0; i < len(slice1); i++ {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}




func main() {
	var testTime = 100
	var maxSize = 100
	var maxValue = 100
	var succeed = true
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < testTime; i++ {
		slice1 := generateRandomArray(maxSize, maxValue)
		slice2 := copyArray(slice1)
		insertionSort(slice1)
		comparator(slice2)
		fmt.Println(slice1)
		fmt.Println(slice2)
		if !isEqual(slice1, slice2) {
			succeed = false
			break
		}
	}
	if succeed {
		fmt.Println("Nice!")
	} else {
		fmt.Println("Fucking fucked!")
	}
}
