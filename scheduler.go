package main

import (
	"fmt"
	"sync"
	"time"
)

type Task func()

type Scheduler struct {
	tasks []Task
}

func (s *Scheduler) addTask(task Task) {
	s.tasks = append(s.tasks, task)
}

func (s *Scheduler) run() {
	var wg sync.WaitGroup
	for i, task := range s.tasks {
		wg.Add(1)
		go func(i int, task Task) {
			defer wg.Done()
			start := time.Now()

			task()
			end := time.Now()
			fmt.Printf("Task %v took %v\n", i+1, end.Sub(start))
		}(i, task)
	}
}

func testScheduler() {
	s := &Scheduler{}

	s.addTask(func() {
		time.Sleep(1 * time.Second)
		fmt.Println("任务 1 完成")
	})

	s.addTask(func() {
		time.Sleep(1 * time.Second)
		fmt.Println("任务 1 完成")
	})

	s.addTask(func() {
		time.Sleep(1 * time.Second)
		fmt.Println("任务 2 完成")
	})

	s.addTask(func() {
		time.Sleep(1 * time.Second)
		fmt.Println("任务 3 完成")
	})
	s.addTask(func() {
		time.Sleep(1 * time.Second)
		fmt.Println("任务 4 完成")
	})

	s.run()

	time.Sleep(10 * time.Second)
}

// 执行所有任务，并记录耗时
func (s *Scheduler) Run() {
	var wg sync.WaitGroup
	for i, task := range s.tasks {
		wg.Add(1)
		go func(i int, t Task) {
			defer wg.Done()
			start := time.Now()
			t()
			duration := time.Since(start)
			fmt.Printf("任务 #%d 耗时: %v\n", i+1, duration)
		}(i, task)
	}
	wg.Wait()
}
