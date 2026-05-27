package common

import "golang.org/x/exp/constraints"

func Max[T constraints.Ordered](v1, v2 T) T {
	if v1 > v2 {
		return v1
	}

	return v2
}

// generic method
func Abs[T constraints.Signed](v1 T) T {
	if v1 >= 0 {
		return v1
	}

	return -v1
}
