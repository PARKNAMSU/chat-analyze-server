package tools

type CustomSlice[T any] struct {
	Slice []T
}

func (c CustomSlice[T]) Reverse() []T {
	for i := 0; i < len(c.Slice)/2; i++ {
		j := len(c.Slice) - i - 1
		c.Slice[i], c.Slice[j] = c.Slice[j], c.Slice[i]
	}
	return c.Slice
}
