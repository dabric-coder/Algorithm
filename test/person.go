package main

import (
	"fmt"
)

type Person struct {
	name string
	age int
}

/*
现在要实现的功能：
对于任意一个类型，可以根据类型的大小输出不同的值
比如传入
*/

type Comparator func(i, j interface{}) int

type cmpPerson struct {
	persons []*Person
	comparator Comparator
}

func NewPersons(persons []*Person, comparator Comparator) *cmpPerson {
	return &cmpPerson{persons:persons, comparator:comparator}
}


func main() {
	ageCmp := func(i, j interface{}) int {
		p1 := i.(Person)
		p2 := j.(Person)
		return p1.age - p2.age
	}
	p1 := Person{"tom", 18}
	p2 := Person{name:"jerry", age:20}
	persons := []*Person{&p1, &p2}
	cmp := NewPersons(persons, ageCmp)
	if cmp.comparator(p1, p2) > 0 {
		fmt.Printf("%s的年龄大于%s的", p1.name, p2.name)
	} else if cmp.comparator(p1, p2) < 0 {
		fmt.Printf("%s的年龄小于%s的", p1.name, p2.name)
	} else {
		fmt.Printf("%s的年龄等于%s的", p1.name, p2.name)
	}

}



