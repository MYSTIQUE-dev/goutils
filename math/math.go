package math

import "golang.org/x/exp/constraints"

// Generic min function.
func Min[T constraints.Ordered](a, b T) T {
    if a < b {
        return a
    }
    return b
}

// Generic max function.
func Max[T constraints.Ordered](a, b T) T {
    if a > b {
        return a
    }
    return b
}
