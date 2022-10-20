package channels

// AsChan is used to convert multi items into channel.
func AsChan[T any](values ...T) <-chan T {
	out := make(chan T)
	go func() {
		for _, v := range values {
			out <- v
		}
		close(out)
	}()
	return out
}
