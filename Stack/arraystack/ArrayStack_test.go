package arraystack

import (
	"fmt"
	"testing"
)

func TestArrayStack(t *testing.T) {
	arrayStack := newArrayStack(10)

	for i := 0; i < 10; i++ {
		arrayStack.push(i)
	}
	arrayStack.printStack()
	fmt.Println(arrayStack.peek())

	arrayStack.pop()
	arrayStack.printStack()
	fmt.Println(arrayStack.peek())

	arrayStack.pop()
	arrayStack.printStack()
	fmt.Println(arrayStack.peek())

}
