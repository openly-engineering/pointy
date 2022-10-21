package pointy

// PointersValueEqual returns true if both pointer parameters are nil or contain the same dereferenced value.
func PointersValueEqual[T comparable](a *T, b *T) bool {
	if a == nil && b == nil {
		return true
	}
	if a != nil && b != nil && *a == *b {
		return true
	}

	return false
}

// PointerValueEqual returns true if the pointer parameter is not nil and contains the same dereferenced value as the value parameter.
func PointerValueEqual[T comparable](a *T, b T) bool {
	if a == nil {
		return false
	}
	if *a == b {
		return true
	}

	return false
}
