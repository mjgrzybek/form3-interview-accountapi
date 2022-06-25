package address

func Of[T any](value T) *T {
	return &value
}
