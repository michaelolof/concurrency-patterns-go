package patterns

// The generator pattern is used to generate a sequence of values
// which is used to produce some output
func Generator() <-chan int {
	ch := make(chan int)

	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()

	return ch
}

func SetGenerator[T any](set []T) <-chan T {
	ch := make(chan T)

	go func() {
		defer close(ch)

		for _, m := range set {
			ch <- m
		}
	}()

	return ch
}
