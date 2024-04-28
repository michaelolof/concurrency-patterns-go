package patterns

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func WorkerPool(idx int, jobs <-chan int, results chan<- int) {

	for job := range jobs {

		go func(job int) {

			dur := rand.Intn(12-1) + 1
			time.Sleep(time.Duration(dur) * time.Second)

			// Start work on job
			fmt.Println("worker", idx, "started job", job)
			val := job * 2

			// Pipe value into result
			results <- val
			fmt.Println("worker", idx, "finished job", job, "result", val, "in", dur, "seconds")

		}(job)
	}

}

func WgWorkerPool(idx int, jobs <-chan int, results chan<- int) {

	var wg sync.WaitGroup

	for job := range jobs {
		wg.Add(1)

		go func(job int) {
			defer wg.Done()

			dur := rand.Intn(12-1) + 1
			time.Sleep(time.Duration(dur) * time.Second)

			// Start work on job
			fmt.Println("worker", idx, "started job", job)
			val := job * 2

			// Pipe value into result
			results <- val
			fmt.Println("worker", idx, "finished job", job, "result", val, "in", dur, "seconds")
		}(job)
	}

	wg.Wait()
}
