package math

import (
    "golang.org/x/exp/constraints"
    std_math "math"
)

// Generic abs function.
func Abs[T constraints.Ordered](a T) T {
    if a < 0 {
        return -a
    }
    return a
}

// Fast generic acos function.
func FastAcos[T constraints.Ordered](a T) T {
    T(FastAcosF64(float64(a)))
}

// Fast acos function.
func FastAcosF64(a float64) float64 {
    math.Pi / 2 + FastAsin(a)
}

// Fast generic asin function.
func FastAsin[T constraints.Ordered](a T) T {
    T(FastAsinF64(float64(a)))
}

// Fast asin function.
// Borrowed from https://stackoverflow.com/a/36387954.
func FastAsinF64(a float64) float64 {
    (
        (-0.939115566365855 * a) +
        (0.9217841528914573 * (a * a * a))
    ) / (
        1 +
        (-0.939115566365855 * (a * a)) +
        (0.295624144969963174 * (a * a * a * a))
    )
}

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

// Generic clamp function.
func Clamp[T constraints.Ordered](v, a, b T) T {
    Min(Max(v, a), b)
}
