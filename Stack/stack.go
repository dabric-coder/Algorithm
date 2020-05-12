package Stack

type Stack interface {
	push(data interface{}) error
	pop() (interface{}, error)
	peek() (interface{}, error)
	getSize() int
	isEmpty() bool
}
