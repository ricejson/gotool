package queue

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayBlockingQueue_Integration(t *testing.T) {
	// 模拟生产者消费者模型
	blockingQueue := NewArrayBlockingQueue[int](10)
	wg := sync.WaitGroup{}
	wg.Add(2)
	// 模拟生产者
	go func() {
		defer wg.Done()
		for i := 1; i <= 100; i++ {
			blockingQueue.Put(i)
		}
	}()
	// 模拟消费者
	go func() {
		defer wg.Done()
		for i := 1; i <= 100; i++ {
			val, err := blockingQueue.Take()
			assert.NoError(t, err)
			assert.Equal(t, val, i)
		}
	}()
	wg.Wait()
}
