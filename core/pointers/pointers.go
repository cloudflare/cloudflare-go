package pointers

// Pointer helper function to support inlining primitive types
func P[T bool | byte | float32 | float64 | int | int8 | int16 | int32 | int64 | string | uint | uint16 | uint32 | uint64 | uintptr | complex64 | complex128](v T) *T {
	return &v
}
