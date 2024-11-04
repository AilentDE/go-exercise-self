package main

import (
	"fmt"
	"sync"
	"time"
)

type job struct {
	count int
}

func (j *job) syncWorker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}
func (j *job) syncJob() {
	var wg sync.WaitGroup

	for i := 1; i <= j.count; i++ {
		wg.Add(1)
		go j.syncWorker(i, &wg)
	}
	wg.Wait() // Wait for all workers to finish
}

func (j *job) channelWorker(id int, ch chan<- int) {
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	ch <- id
}
func (j *job) channelJob() {
	chs := make([]chan int, j.count)
	for i:=1; i<=j.count; i++ {
		chs[i-1] = make(chan int)
		go j.channelWorker(i, chs[i-1])
	}

	// Wait for all workers to finish
	for _, ch := range chs {
		id := <-ch
		fmt.Printf("Worker %d done\n", id)
	}
}

func main() {
	j := job{count: 5}

	fmt.Printf("[sync]Start %d workers\n", j.count)
	startTime := time.Now()
	j.syncJob()
	fmt.Printf("[sync]Job done in %v\n", time.Since(startTime))

	fmt.Println(
		"------------------------------------------------------------",
	)

	fmt.Printf("[channel]Start %d workers\n", j.count)
	startTime = time.Now()
	j.channelJob()
	fmt.Printf("[channel]Job done in %v\n", time.Since(startTime))
}