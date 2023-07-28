package utils

func Ternary[T any](condition bool, l1, l2 T) T {
	if condition {
		return l1
	}
	return l2
}
