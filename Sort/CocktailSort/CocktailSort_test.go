package CocktailSort

import (
	"fmt"
	"testing"
)

func TestCocktailSort(t *testing.T) {
	 arr := []int{1,9,2,8,3,7,4,6,5,10}
	 CocktailSort(arr)
	 fmt.Println(arr)
}
