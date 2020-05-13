package ShellSort

import (
	"fmt"
	"testing"
)

func TestShellSort(t *testing.T) {
	arr := []int{11,1,9,2,8,3,7,4,6,5,10}
	ShellSort(arr)
	fmt.Println(arr)
}