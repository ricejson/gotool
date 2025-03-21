package stack

import "errors"

type ArrayStack[T any] struct {
	elems    []T // 底层切片
	top      int // 栈顶指针
	capacity int // 栈容量
}

func NewArrayStack[T any](capacity int) Stack[T] {
	var elems = make([]T, capacity)
	return &ArrayStack[T]{
		elems:    elems,
		top:      0,
		capacity: capacity,
	}
}

// ensureCapacity 自动扩容处理
func (s *ArrayStack[T]) ensureCapacity() {
	if s.top == s.capacity {
		// 二倍扩容
		s.capacity <<= 1
		var newElems = make([]T, s.capacity)
		copy(newElems, s.elems)
		s.elems = newElems
	}
}

// Push 入栈方法
func (s *ArrayStack[T]) Push(elem T) error {
	// 自动扩容处理
	s.ensureCapacity()
	s.elems[s.top] = elem
	s.top++
	return nil
}

// Pop 出栈方法
func (s *ArrayStack[T]) Pop() (T, error) {
	// 特殊处理栈为空的情况
	var zeroValue T
	if s.Empty() {
		return zeroValue, errors.New("stack is empty")
	}
	var target = s.elems[s.top-1]
	s.top--
	return target, nil
}

// Peek 获取栈顶元素
func (s *ArrayStack[T]) Peek() (T, error) {
	// 特殊处理栈为空的情况
	var zeroValue T
	if s.Empty() {
		return zeroValue, errors.New("stack is empty")
	}
	return s.elems[s.top-1], nil
}

// Empty 判断栈是否为空
func (s *ArrayStack[T]) Empty() bool {
	return s.top == 0
}
