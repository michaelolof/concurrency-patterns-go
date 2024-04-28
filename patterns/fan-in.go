package patterns

import "sync"

func FanIn(inputs ...<-chan int) <-chan int {

	var wg sync.WaitGroup
	out := make(chan int)

	wg.Add(len(inputs))
	for _, input := range inputs {
		go func(ch <-chan int) {
			for {
				value, ok := <-ch
				if !ok {
					wg.Done()
					break
				}

				out <- value
			}
		}(input)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
