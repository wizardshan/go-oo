package slice

func Walker[T any](slice []T, f func(T)) {
	for _, e := range slice {
		f(e)
	}
}

func Filter[T any](slice []T, f func(T) bool) []T {
	var s []T
	for _, e := range slice {
		if f(e) {
			s = append(s, e)
		}
	}
	return s
}
