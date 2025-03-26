package queue

import (
	"sync"
)

// ArrayBlockingQueue 阻塞队列（基于数组）
type ArrayBlockingQueue[T any] struct {
	q         *ArrayQueue[T]
	mutex     sync.Mutex
	emptyCond *sync.Cond
	fullCond  *sync.Cond
}

func NewArrayBlockingQueue[T any](capacity int) *ArrayBlockingQueue[T] {
	var arrayQueue = ArrayQueue[T]{
		data:     make([]T, capacity),
		head:     0,
		tail:     0,
		usedSize: 0,
		capacity: capacity,
	}
	var arrayBlockingQueue = &ArrayBlockingQueue[T]{
		q: &arrayQueue,
	}
	arrayBlockingQueue.emptyCond = sync.NewCond(&arrayBlockingQueue.mutex)
	arrayBlockingQueue.fullCond = sync.NewCond(&arrayBlockingQueue.mutex)
	return arrayBlockingQueue
}

// Put 插入元素到阻塞队列中（队列已满就阻塞）
func (blockingQueue *ArrayBlockingQueue[T]) Put(elem T) error {
	blockingQueue.mutex.Lock()
	defer blockingQueue.mutex.Unlock()
	for blockingQueue.q.IsFull() {
		// 进行阻塞
		blockingQueue.fullCond.Wait()
	}
	err := blockingQueue.q.Offer(elem)
	blockingQueue.emptyCond.Signal()
	return err
}

// Take 从阻塞队列中取出元素（队列为空就阻塞）
func (blockingQueue *ArrayBlockingQueue[T]) Take() (T, error) {
	blockingQueue.mutex.Lock()
	defer blockingQueue.mutex.Unlock()
	for blockingQueue.q.IsEmpty() {
		// 进行阻塞
		blockingQueue.emptyCond.Wait()
	}
	topValue, err := blockingQueue.q.Poll()
	blockingQueue.fullCond.Signal()
	return topValue, err
}
