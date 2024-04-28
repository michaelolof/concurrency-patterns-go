package patterns

func FanOut[T any](in <-chan T) <-chan T {

	rtn := make(chan T)

	go func() {
		defer close(rtn)

		for val := range in {
			rtn <- val
		}
	}()

	return rtn
}
