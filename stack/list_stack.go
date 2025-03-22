package stack

import (
	"container/list"
	"errors"
)

type ListStack[T any] struct {
	list *list.List // 双向链表
}

func NewListStack[T any]() Stack[T] {
	return &ListStack[T]{
		list: list.New(),
	}
}

// Push 入栈方法
func (s *ListStack[T]) Push(elem T) error {
	// 链表尾部插入元素
	s.list.PushBack(elem)
	return nil
}

// Pop 出栈方法
func (s *ListStack[T]) Pop() (T, error) {
	// 特殊处理栈为空的情况
	var zeroValue T
	if s.Empty() {
		return zeroValue, errors.New("stack is empty")
	}
	// 链表尾部删除元素
	remElem := s.list.Remove(s.list.Back())
	return remElem.(T), nil
}

// Peek 获取栈顶元素
func (s *ListStack[T]) Peek() (T, error) {
	// 特殊处理栈为空的情况
	var zeroValue T
	if s.Empty() {
		return zeroValue, errors.New("stack is empty")
	}
	return s.list.Back().Value.(T), nil
}

// Empty 判断栈是否为空
func (s *ListStack[T]) Empty() bool {
	return s.list.Len() == 0
}
