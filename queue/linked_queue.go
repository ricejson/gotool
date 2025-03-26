package queue

import (
	"errors"
	"fmt"
	"github.com/ricejson/gotool/internal/list"
)

// LinkedQueue 队列结构（基于链表）
type LinkedQueue[T any] struct {
	list *list.LinkedList[T]
}

func NewLinkedQueue[T any]() *LinkedQueue[T] {
	return &LinkedQueue[T]{
		list: list.NewLinkedList[T](),
	}
}

// Offer 往队尾插入元素
func (q *LinkedQueue[T]) Offer(elem T) error {
	q.list.AddLast(elem)
	return nil
}

// Poll 从队头取出元素
func (q *LinkedQueue[T]) Poll() (T, error) {
	// 判断队列是否为空
	var zeroValue T
	if q.IsEmpty() {
		return zeroValue, errors.New("queue is empty")
	}
	val, err := q.list.RemoveFirst()
	if err != nil {
		return zeroValue, fmt.Errorf("queue poll error: %w", err)
	}
	return val, nil
}

// Peek 查看队头元素
func (q *LinkedQueue[T]) Peek() (T, error) {
	// 判断队列是否为空
	var zeroValue T
	if q.IsEmpty() {
		return zeroValue, errors.New("queue is empty")
	}
	val, err := q.list.GetFirst()
	if err != nil {
		return zeroValue, fmt.Errorf("queue poll error: %w", err)
	}
	return val, nil
}

// Size 判断队列存储的元素个数
func (q *LinkedQueue[T]) Size() int {
	return q.list.Size()
}

// IsEmpty 判断队列是否为空
func (q *LinkedQueue[T]) IsEmpty() bool {
	return q.Size() == 0
}
