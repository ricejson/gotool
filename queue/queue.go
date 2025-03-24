package queue

import (
	"errors"
	"fmt"
	"github.com/ricejson/gotool/internal/list"
)

// Queue 队列结构（基于链表）
type Queue[T any] struct {
	list *list.LinkedList[T]
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		list: list.NewLinkedList[T](),
	}
}

// Offer 往队尾插入元素
func (q *Queue[T]) Offer(elem T) error {
	q.list.AddLast(elem)
	return nil
}

// Poll 从队头取出元素
func (q *Queue[T]) Poll() (T, error) {
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
func (q *Queue[T]) Peek() (T, error) {
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
func (q *Queue[T]) Size() int {
	return q.list.Size()
}

// IsEmpty 判断队列是否为空
func (q *Queue[T]) IsEmpty() bool {
	return q.Size() == 0
}
