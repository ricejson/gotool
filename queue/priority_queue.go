package queue

import "errors"

// PriorityQueue 优先级队列
type PriorityQueue[T any] struct {
	arr       []T
	usedSize  int
	capacity  int
	compareTo func(e1, e2 T) int // 返回小于等于0表示升序，大于0表示降序
}

func NewPriorityQueue[T any](capacity int, compareTo func(e1, e2 T) int) *PriorityQueue[T] {
	return &PriorityQueue[T]{
		arr:       make([]T, capacity),
		capacity:  capacity,
		usedSize:  0,
		compareTo: compareTo,
	}
}

func (pq *PriorityQueue[T]) ensureCapacity() error {
	if pq.usedSize >= pq.capacity {
		// 进行扩容
		pq.capacity <<= 1
		var newArr = make([]T, pq.capacity)
		copy(newArr, pq.arr)
		pq.arr = newArr
	}
	return nil
}

func (pq *PriorityQueue[T]) adjust() {
	if pq.usedSize == 1 {
		// 不需要调整
		return
	}
	// childIdx 为最后一个左孩子
	var childIdx = pq.usedSize - 1
	var root = (childIdx - 1) / 2
	for i := root; i >= 0; i-- {
		// 开始调整
		pq.adjustDown(i)
	}
}

func (pq *PriorityQueue[T]) adjustDown(root int) {
	// childIdx 为根节点左孩子
	var childIdx = 2*root + 1
	// 开始调整
	for childIdx <= pq.usedSize-1 {
		var targetIdx int
		if childIdx+1 >= pq.usedSize {
			// 已经是最后一个元素了
			targetIdx = childIdx
		} else {
			if pq.compareTo(pq.arr[childIdx], pq.arr[childIdx+1]) <= 0 {
				targetIdx = childIdx
			} else {
				targetIdx = childIdx + 1
			}
		}
		if pq.compareTo(pq.arr[targetIdx], pq.arr[root]) <= 0 {
			// 交换
			pq.arr[root], pq.arr[targetIdx] = pq.arr[targetIdx], pq.arr[root]
			root = childIdx
			childIdx = 2*root + 1
		} else {
			break
		}
	}
}

func (pq *PriorityQueue[T]) Offer(elem T) error {
	// 扩容处理
	err := pq.ensureCapacity()
	if err != nil {
		return errors.New("ensure capacity error")
	}
	pq.arr[pq.usedSize] = elem
	pq.usedSize++
	// 调整
	pq.adjust()
	return nil
}

func (pq *PriorityQueue[T]) Peek() (T, error) {
	// 判断队列是否为空
	var zeroValue T
	if pq.IsEmpty() {
		return zeroValue, errors.New("priority queue is empty")
	}
	return pq.arr[0], nil
}

func (pq *PriorityQueue[T]) Poll() (T, error) {
	// 判断队列是否为空
	var zeroValue T
	if pq.IsEmpty() {
		return zeroValue, errors.New("priority queue is empty")
	}
	// 交换头尾元素
	var topValue = pq.arr[0]
	pq.arr[0], pq.arr[pq.usedSize-1] = pq.arr[pq.usedSize-1], pq.arr[0]
	pq.usedSize--
	// 调整
	pq.adjustDown(0)
	return topValue, nil
}

func (pq *PriorityQueue[T]) IsEmpty() bool {
	return pq.Size() == 0
}

func (pq *PriorityQueue[T]) Size() int {
	return pq.usedSize
}
