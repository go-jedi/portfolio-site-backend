package filter

func Filter[T any](data []T, f func(T) bool) []T {
	result := make([]T, 0, len(data))

	for _, value := range data {
		if f(value) {
			result = append(result, value)
		}
	}

	return result
}
