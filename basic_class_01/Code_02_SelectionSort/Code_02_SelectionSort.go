package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func selectionSort(slice []int) {
	for i := 0; i < len(slice); i++ {
		minIndex := i
		for j := i + 1; j < len(slice); j++ {
			if slice[minIndex] > slice[j] {
				swap(slice, minIndex, j)
			}
		}
	}
}

func swap(slice []int, x, y int) {
	slice[x] = slice[x] ^ slice[y]
	slice[y] = slice[x] ^ slice[y]
	slice[x] = slice[x] ^ slice[y]
}

// for test
func comparator(slice []int) {
	sort.Ints(slice)
}

func generateRandomArray(maxSize, maxValue int) (slice []int) {
	for i := 0; i < rand.Intn(maxSize+1); i++ {
		slice = append(slice, rand.Intn(maxValue+1) - rand.Intn(maxValue))
	}
	return
}

func copyArry(slice []int) []int {
	if slice == nil {
		return nil
	}
	var res []int
	for i := 0; i < len(slice); i++ {
		res = append(res, slice[i])
	}
	return res
}

func isEqule(slice1, slice2 []int) bool {
	if (slice1 == nil && slice2 != nil) || (slice1 != nil && slice2 == nil) {
		return false
	}

	if len(slice1) != len(slice2) {
		return false
	}

	if slice1 == nil && slice2 == nil {
		return true
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
	for i := 0; i < testTime; i++ {
		slice1 := generateRandomArray(maxSize, maxValue)
		slice2 := copyArry(slice1)
		selectionSort(slice1)
		comparator(slice2)
		fmt.Println(slice1)
		fmt.Println(slice2)
		if !isEqule(slice1, slice2) {
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
