package main

import (
	"errors"
	"fmt"
	"time"
)


type Task struct {
	Id  int
	Err error
	f   func() error
}


func (task *Task) Do() error {
	return task.f()
}


type WorkerPool struct {
	PoolSize    int
	tasksSize   int
	tasksChan   chan Task
	resultsChan chan Task
	Results     func() []Task
}


func NewWorkerPool(tasks []Task, size int) *WorkerPool {
	tasksChan := make(chan Task, len(tasks))
	resultsChan := make(chan Task, len(tasks))
	for _, task := range tasks {
		tasksChan <- task
	}
	close(tasksChan)
	pool := &WorkerPool{PoolSize: size, tasksSize: len(tasks), tasksChan: tasksChan, resultsChan: resultsChan}
	pool.Results = pool.results
	return pool
}


func (pool *WorkerPool) Start() {
	for i := 0; i < pool.PoolSize; i++ {
		go pool.worker()
	}
}


func (pool *WorkerPool) worker() {
	for task := range pool.tasksChan {
		task.Err = task.Do()
		pool.resultsChan <- task
	}
}


func (pool *WorkerPool) results() []Task {
	tasks := make([]Task, pool.tasksSize)
	for i := 0; i < pool.tasksSize; i++ {
		tasks[i] = <-pool.resultsChan
	}
	return tasks
}

/*
2 worker pool封装
1中所示的代码难以满足真实业务场景需求，我们需要对worker pool作一层抽象，封装的更通用一点。
如下代码封装了一个worker pool，WorkerPool的创建需要传入要处理的任务列表及指定pool的大小，Task为任务的封装，需提供该任务的实现。
创建完worker pool，调用pool.Start()即进入多goroutine处理，调用pool.Results()即会阻塞等待所有任务的执行结果。
*/
func main() {
	t := time.Now()


	tasks := []Task{
		{Id: 0, f: func() error { time.Sleep(2 * time.Second); fmt.Println(0); return nil }},
		{Id: 1, f: func() error { time.Sleep(time.Second); fmt.Println(1); return errors.New("error") }},
		{Id: 2, f: func() error { fmt.Println(2); return errors.New("error") }},
	}
	pool := NewWorkerPool(tasks, 2)
	pool.Start()

	tasks = pool.Results()
	fmt.Printf("all tasks finished, timeElapsed: %f s\n", time.Now().Sub(t).Seconds())
	for _, task := range tasks {
		fmt.Printf("result of task %d is %v\n", task.Id, task.Err)
	}
}