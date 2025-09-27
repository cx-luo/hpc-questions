// Package src coding=utf-8
// @Project : concurrency
// @Time    : 2023/11/24 16:28
// @Author  : chengxiang.luo
// @Email   : chengxiang.luo@foxmail.com
// @File    : concurrency.go
// @Software: GoLand
package src

import "sync"

type Semaphore struct {
	c  chan struct{}
	wg sync.WaitGroup
}

// NewSemaphore returns a new Semaphore initialized to the given value.
func NewSemaphore(maxCount int) *Semaphore {
	return &Semaphore{c: make(chan struct{}, maxCount)}
}

// Acquire acquires a permit, blocking until it becomes available or ctx is done.
func (s *Semaphore) Acquire(delta int) {
	s.wg.Add(delta)
	for i := 0; i < delta; i++ {
		s.c <- struct{}{}
	}
}

// Release releases a permit.
func (s *Semaphore) Release() {
	<-s.c
	s.wg.Done()
}

// Wait blocks until all permits have been released.
func (s *Semaphore) Wait() {
	s.wg.Wait()
}
