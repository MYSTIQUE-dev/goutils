// Useful slice manipulation utility functions.

package slices

func FindIndex[T comparable](slice []T, item T) int {
	for i := 0; i < len(slice); i++ {
		if slice[i] == item {
			return i
		}
	}

	return -1
}

func Concat[T any](s1 []T, s2 []T) []T {
	return append(s1, s2...)
}

func Contains[T comparable](slice []T, item T) bool {
	return FindIndex(slice, item) > -1
}

func Clone[T any](slice []T) []T {
	cloned := make([]T, len(slice))
	copy(cloned, slice)

	return cloned
}

func Insert[T any](slice []T, i int, items ...T) []T {
	return append(slice[:i], append(items, slice[i+1:]...)...)
}

func Append[T any](slice []T, items ...T) []T {
	return append(slice, items...)
}

func Prepend[T any](slice []T, items ...T) []T {
	return append(items, slice...)
}

func Remove[T any](slice []T, i int) []T {
	return append(slice[:i], slice[i+1:]...)
}

func RemoveItems[T comparable](slice []T, items ...T) []T {
	for i := 0; i < len(items); i++ {
		if ii := FindIndex(slice, items[i]); ii > -1 {
			slice = Remove(slice, ii)
		}
	}

	return slice
}

func Reverse[T any](slice []T) []T {
	length := len(slice)
	reversed := make([]T, length)

	for i := 0; i < length; i++ {
		reversed = append(reversed, slice[length-i-1])
	}

	return reversed
}

func MakeFilled[T any](item T, length int) []T {
	filled := make([]T, length)

	for i := 0; i < length; i++ {
		filled = append(filled, item)
	}

	return filled
}
