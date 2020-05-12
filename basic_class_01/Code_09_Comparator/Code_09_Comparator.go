package main

import (
	"fmt"
	"sort"
)

type student struct {
	id   int
	name string
	age  int
}

type comparator func (i, j *student) bool

type studentWrapper struct {
	students []student
	comparator comparator
}


func (s studentWrapper) Len() int {
	return len(s.students)
}

func (s studentWrapper) Less(i, j int) bool {
	return s.comparator(&s.students[i], &s.students[j])   // 从小到大排序
}

func (s studentWrapper) Swap(i, j int) {
	s.students[i], s.students[j] = s.students[j], s.students[i]
}

func sortStudent(students []student, comparator comparator) {
	sort.Sort(studentWrapper{students, comparator})
}

func main() {
	stu1 := student{
		id:   0,
		name: "tom",
		age:  18,
	}
	stu2 := student{2,"jerry",16}
	stu3 := student{1, "mary", 20}
	students  := []student{stu1, stu2, stu3}

	idComparator := func(i, j *student) bool {
		return i.id < j.id
	}
	s := studentWrapper{students, idComparator}
	// 使用封装的sortStudent函数对三个学生实例进行排序
	sortStudent(students, idComparator)
	fmt.Println(s.students)   // [{0 tom 18} {1 mary 20} {2 jerry 16}]
}
