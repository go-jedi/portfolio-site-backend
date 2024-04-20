package mapfn

func Map[T, U any](list []T, f func(T) U) []U {
	result := make([]U, 0, len(list))

	for _, value := range list {
		result = append(result, f(value))
	}

	return result
}
