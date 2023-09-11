// File contains helper methods for accepting variants (pointers, values,
// slices, etc) of a particular type and returning them in another. A common use
// is pointer to values and back.
package cloudflare

// Ref is a helper routine that allocates a new value to store v and
// returns a pointer to it.
func Ref[T any](v T) *T { return &v }

// Deref is a helper routine that accepts a pointer and returns the value
// it points to, or the zero value if the pointer is nil.
func Deref[T any](v *T) T {
	if v != nil {
		return *v
	}
	var zero T
	return zero
}

// RefSlice converts a slice of values into a slice of pointers.
func RefSlice[T any](src []T) []*T {
	dst := make([]*T, len(src))
	for i := 0; i < len(src); i++ {
		dst[i] = &(src[i])
	}
	return dst
}

// DerefSlice converts a slice of pointers into a slice of values.
func DerefSlice[T any](src []*T) []T {
	dst := make([]T, len(src))
	for i := 0; i < len(src); i++ {
		if src[i] != nil {
			dst[i] = *(src[i])
		}
	}
	return dst
}

// RefMap converts a string map of values into a string map of pointers.
func RefMap[T any](src map[string]T) map[string]*T {
	dst := make(map[string]*T)
	for k, val := range src {
		v := val
		dst[k] = &v
	}
	return dst
}

// DerefMap converts a string map of pointers into a string map of values.
func DerefMap[T any](src map[string]*T) map[string]T {
	dst := make(map[string]T)
	for k, val := range src {
		if val != nil {
			dst[k] = *val
		}
	}
	return dst
}
