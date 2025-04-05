package queue

import "errors"

// ArrayQueue 队列结构（基于数组）
type ArrayQueue[T any] struct {
	data     []T // 底层数组
	head     int // 队头指针
	tail     int // 队尾指针
	usedSize int // 存储的元素个数
	capacity int // 容量
}

func NewArrayQueue[T any](capacity int) Queue[T] {
	return &ArrayQueue[T]{
		data:     make([]T, capacity),
		head:     0,
		tail:     0,
		usedSize: 0,
		capacity: capacity,
	}
}

func (q *ArrayQueue[T]) Offer(elem T) error {
	// 判断队列是否已满
	if q.IsFull() {
		return errors.New("queue is full")
	}
	// 插入元素
	q.data[q.tail] = elem
	q.tail = (q.tail + 1) % q.capacity
	q.usedSize++
	return nil
}

func (q *ArrayQueue[T]) Poll() (T, error) {
	// 判断队列是否为空
	var zeroValue T
	if q.IsEmpty() {
		return zeroValue, errors.New("queue is empty")
	}
	var topValue = q.data[q.head]
	q.head = (q.head + 1) % q.capacity
	q.usedSize--
	return topValue, nil
}

func (q *ArrayQueue[T]) Peek() (T, error) {
	// 判断队列是否为空
	var zeroValue T
	if q.IsEmpty() {
		return zeroValue, errors.New("queue is empty")
	}
	return q.data[q.head], nil
}

// Size 获取队列存储元素个数
func (q *ArrayQueue[T]) Size() int {
	return q.usedSize
}

// IsFull 判断队列是否已满
func (q *ArrayQueue[T]) IsFull() bool {
	return q.usedSize == q.capacity
}

// IsEmpty 判断队列是否为空
func (q *ArrayQueue[T]) IsEmpty() bool {
	return q.usedSize == 0
}
