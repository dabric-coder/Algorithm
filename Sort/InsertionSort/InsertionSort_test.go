package InsertionSort

import (
	"fmt"
	"testing"
	"algorithm/Sort/utils"
)

type Person struct {
	name string
	age int
}

type Persons []*Person

func (p Persons) Len() int {
	return len(p)
}

func (p Persons) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

type ByName struct {
	Persons
}

func (s ByName) Comparator(i, j int) bool {
	return s.Persons[i].name > s.Persons[j].name
}


func TestInsertionSortIntSlice(t *testing.T) {
	arr := []int{10, 1, 3, 9, 6, 8, 5, 0}

	a := utils.IntSlice(arr)

	InsertionSort(a)

	fmt.Println(a)

}

func TestInsertionSortPerson(t *testing.T) {
	persons := []*Person{
		{"tom", 18},
		{"alice", 20},
		{"jerry", 22},
		{"jarry", 21},
	}
	fmt.Println(persons[0].name)
	nameSort := ByName{persons}
	InsertionSort(nameSort)
	for i := 0; i < len(persons); i++ {
		fmt.Println(nameSort.Persons[i].name)
	}
}