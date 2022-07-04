package functional

func Map[T, R any](arr []T, f func(e T) R) []R {
	mapped := make([]R, len(arr))
	for i, e := range arr {
		mapped[i] = f(e)
	}
	return mapped
}
