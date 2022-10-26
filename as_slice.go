package channels

// AsSlice is used to convert channel to a slice.
// Be careful using this func because it can lead to memory explosion.
func AsSlice[T any](c <-chan T) []T {
	out := make([]T, 0, cap(c))
	for v := range c {
		out = append(out, v)
	}
	return out
}
