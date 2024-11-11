package main

import (
	"fmt"
	"sync"
	"time"

	"limited-goroutine/store/stored_data"
)

func main() {
    items := stored_data.CreateItems(50)

    fmt.Println("Processing start At", time.Now())

    var wg sync.WaitGroup
    sem := make(chan struct{}, 5)
    for _, item := range *items {
        wg.Add(1)
        sem <- struct{}{}
        go func(i stored_data.Item) {
            defer wg.Done()
            process(i)
            <-sem
        }(item)
    }

    wg.Wait() //用來避免 main 比最後一個 goroutine 還要早結束
    fmt.Println("Processing end At", time.Now())
}
func process(i stored_data.Item) {
	fmt.Println("Processing item", i.Id)
	time.Sleep(1 * time.Second)
	fmt.Println("Processed item", i.Id)
}