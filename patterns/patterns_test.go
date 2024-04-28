package patterns

import (
	"fmt"
	"testing"
	"time"
)

func TestGenerator(t *testing.T) {

	gen := Generator()
	one := <-gen
	time.Sleep(time.Second * 10)
	two := <-gen

	fmt.Println("One:", one, "Two:", two)
}

func TestFanIn(t *testing.T) {

	one := SetGenerator([]int{0, 2, 4, 6, 8})
	two := SetGenerator([]int{1, 3, 5, 7, 9})

	out := FanIn(one, two)

	for value := range out {
		fmt.Println("Value is:", value)
	}

	fmt.Println("done")
}

func TestFanOut(t *testing.T) {

	set := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

	in := SetGenerator(set)

	out1 := FanOut(in)
	out2 := FanOut(in)
	out3 := FanOut(in)
	out4 := FanOut(in)

	for range set {
		select {
		case val1 := <-out1:
			fmt.Println("#1 processed", val1)
		case val2 := <-out2:
			fmt.Println("#2 processed", val2)
		case val3 := <-out3:
			fmt.Println("#3 processed", val3)
		case val4 := <-out4:
			fmt.Println("#4 processed", val4)
		}
	}

	fmt.Println("done")
}

func TestWorkerPool(t *testing.T) {

	const totalJobs int = 9
	const totalWorkers int = 3 // Ideally you should call [runtime.NumCPU] function

	jobs := make(chan int, totalJobs)
	results := make(chan int, totalJobs)

	now := time.Now()

	// Spin off workers
	for i := 1; i <= totalWorkers; i++ {
		go WorkerPool(i, jobs, results)
	}

	// Send in jobs to the worker pool
	for i := 1; i <= totalJobs; i++ {
		jobs <- i
	}
	close(jobs)

	// Retrieve results from the worker pool
	for i := 1; i <= totalJobs; i++ {
		fmt.Println("Result recieved", <-results)
	}
	close(results)

	dur := time.Since(now)
	fmt.Println("Finished in", dur, "milliseconds")

	fmt.Println("done")

}
