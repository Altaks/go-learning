package main

import (
	"fmt"
	"sync"
	"time"
)

func validate(id int) {
	fmt.Printf("Worker %d validated\n", id)
}

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)

		go func() {
			// Should be after the worker call but whatever
			defer wg.Done()
			worker(i)
			validate(i)
		}()
	}

	wg.Wait()
}
