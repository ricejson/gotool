package list

import (
	"container/list"
	"errors"
)

// LinkedList 双向链表泛型实现
type LinkedList[T any] struct {
	list *list.List
}

func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{
		list: list.New(),
	}
}

// AddLast 尾插元素
func (l *LinkedList[T]) AddLast(elem T) T {
	return l.list.PushBack(elem).Value.(T)
}

// RemoveLast 尾删元素
func (l *LinkedList[T]) RemoveLast() (T, error) {
	var zeroValue T
	// 特殊判断为空的情况
	if l.list.Len() == 0 {
		return zeroValue, errors.New("list is empty")
	}
	return l.list.Remove(l.list.Back()).(T), nil
}

// GetLast 获取尾部元素
func (l *LinkedList[T]) GetLast() (T, error) {
	var zeroValue T
	// 特殊判断为空的情况
	if l.list.Len() == 0 {
		return zeroValue, errors.New("list is empty")
	}
	return l.list.Back().Value.(T), nil
}

// Size 链表元素个数
func (l *LinkedList[T]) Size() int {
	return l.list.Len()
}
