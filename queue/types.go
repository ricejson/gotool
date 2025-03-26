package queue

// Queue 队列接口定义
type Queue[T any] interface {
	Offer(elem T) error
	Poll() (T, error)
	Peek() (T, error)
	Size() int
	IsEmpty() bool
}
