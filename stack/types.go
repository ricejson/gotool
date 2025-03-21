package stack

// Stack 栈的接口定义
type Stack[T any] interface {
	Push(elem T) error
	Pop() (T, error)
	Peek() (T, error)
	Empty() bool
}
