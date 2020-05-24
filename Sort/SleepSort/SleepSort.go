package main

import (
	"fmt"
	"time"
)

var flag bool
var container chan bool
var count int
func SleepSort1(arr []int) {
	flag = true  // 标识，区分
	container = make(chan bool, 5) // 5个管道
	for i := 0; i < len(arr); i++ {
		go toSleep(arr[i])
	}

	go listen1(len(arr))
	for flag {
		time.Sleep(1*time.Second)
	}
}


func toSleep1(data int) {
	time.Sleep(time.Duration(data)*time.Millisecond)
	fmt.Println(data)
	container <- true  // 管道输入OK，代表该数据已经排好序
}

func listen1(size int) {
	for flag {
		select {
		case <- container:
			count++
			if count >= size {
				flag = false
				break
			}
		}
	}
}


func SleepSort(arr []int) {

	flag = true
	container = make(chan bool, 10)

	for i := 0; i < len(arr); i++ {
		go toSleep(arr[i])
	}

	go listen(len(arr))

	for flag {
		time.Sleep(1 * time.Second)
	}
}

func toSleep(data int) {
	time.Sleep(time.Duration(data)*time.Millisecond)
	fmt.Println(data)
	container <- true
}

func listen(size int) {
	for flag {
		select {
		case <- container:
			count++
			if count >= size {
				flag = false
				break
			}
		}
	}
}


func main() {
	arr := []int{0,10,16,8,1,24,30,100}
	SleepSort(arr)
}