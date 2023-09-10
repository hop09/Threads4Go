// purely written by Muhammad Hamza (HOP)
// this is a low-level try of implementation of something like concurrent.futures ThreadPoolExecutor module of python
// there is much to add here. Hope i may add soon.
// i named it gothreads that's stupidity , go uses goroutines. im not expert hehehehehe

package gothreads

import (
	"sync"
)

type ThreadPoolExecutor struct {
	workers   int
	taskQueue chan func()
	wg        sync.WaitGroup
}

func NewThreadPoolExecutor(workers int) *ThreadPoolExecutor {
	return &ThreadPoolExecutor{
		workers:   workers,
		taskQueue: make(chan func()),
	}
}

func (e *ThreadPoolExecutor) Start() {
	for i := 0; i < e.workers; i++ {
		e.wg.Add(1)
		go func() {
			defer e.wg.Done()
			for task := range e.taskQueue {
				task()
			}
		}()
	}
}

func (e *ThreadPoolExecutor) Submit(task func()) {
	e.taskQueue <- task
}

func (e *ThreadPoolExecutor) Stop() {
	close(e.taskQueue)
	e.wg.Wait()
}
