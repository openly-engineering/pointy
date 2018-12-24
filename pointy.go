// Package pointy is a set of simple helper functions to provide a shorthand to
// get a pointer to a variable holding a constant.
package pointy

// Bool returns a pointer to a variable holding the supplied bool constant
func Bool(x bool) *bool {
	return &x
}

// Byte returns a pointer to a variable holding the supplied byte constant
func Byte(x byte) *byte {
	return &x
}

// Complex128 returns a pointer to a variable holding the supplied complex128 constant
func Complex128(x complex128) *complex128 {
	return &x
}

// Complex64 returns a pointer to a variable holding the supplied complex64 constant
func Complex64(x complex64) *complex64 {
	return &x
}

// Float32 returns a pointer to a variable holding the supplied float32 constant
func Float32(x float32) *float32 {
	return &x
}

// Float64 returns a pointer to a variable holding the supplied float64 constant
func Float64(x float64) *float64 {
	return &x
}

// Int returns a pointer to a variable holding the supplied int constant
func Int(x int) *int {
	return &x
}

// Int8 returns a pointer to a variable holding the supplied int8 constant
func Int8(x int8) *int8 {
	return &x
}

// Int16 returns a pointer to a variable holding the supplied int16 constant
func Int16(x int16) *int16 {
	return &x
}

// Int32 returns a pointer to a variable holding the supplied int32 constant
func Int32(x int32) *int32 {
	return &x
}

// Int64 returns a pointer to a variable holding the supplied int64 constant
func Int64(x int64) *int64 {
	return &x
}

// Uint returns a pointer to a variable holding the supplied uint constant
func Uint(x uint) *uint {
	return &x
}

// Uint8 returns a pointer to a variable holding the supplied uint8 constant
func Uint8(x uint8) *uint8 {
	return &x
}

// Uint16 returns a pointer to a variable holding the supplied uint16 constant
func Uint16(x uint16) *uint16 {
	return &x
}

// Uint32 returns a pointer to a variable holding the supplied uint32 constant
func Uint32(x uint32) *uint32 {
	return &x
}

// Uint64 returns a pointer to a variable holding the supplied uint64 constant
func Uint64(x uint64) *uint64 {
	return &x
}

// String returns a pointer to a variable holding the supplied string constant
func String(x string) *string {
	return &x
}

// Rune returns a pointer to a variable holding the supplied rune constant
func Rune(x rune) *rune {
	return &x
}
