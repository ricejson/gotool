package queue

import (
	"fmt"
	"sync"
	"testing"
	"time"
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
			fmt.Println("生产了: ", i)
		}
	}()
	// 模拟消费者
	go func() {
		defer wg.Done()
		for i := 1; i <= 100; i++ {
			time.Sleep(time.Second * 2)
			val, _ := blockingQueue.Take()
			fmt.Println("消费了: ", val)
		}
	}()
	wg.Wait()
}
