package reduce

func Reduce[T any](list []T, accumulator func(T, T) T, init T) T {
	for _, value := range list {
		init = accumulator(init, value)
	}

	return init
}
