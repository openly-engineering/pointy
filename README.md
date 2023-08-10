# pointy

Simple helper functions to provide a shorthand to get a pointer to a variable holding a constant...because it's annoying when you have to do it hundreds of times in unit tests:

```golang

val := 42
pointerToVal := &val
// vs.
pointerToVal := pointy.Int(42) // if using Go 1.17 or earlier w/o generics
pointerToVal := pointy.Pointer(42) // if using Go 1.18+ w/ generics
```

### New in release 2.0.0

🚨 Breaking change

Package has changed to `go.openly.dev`. Please use 
```
import "go.openly.dev/pointy"
```

### New in release 1.2.0

Generic implementation of the pointer-to-value and value-to-pointer functions. *Requires Go 1.18+.*
The type-specific functions are still available for backwards-compatibility. 

```golang
pointerToInt := pointy.Pointer(42) 
pointerToString := pointy.Pointer("foo") 
// then later in your code..
intValue := pointy.PointerValue(pointerToInt, 99) 
stringValue := pointy.PointerValue(pointerToString, "bar") 
```

Convenience functions to safely compare pointers by their dereferenced values:

```golang
// when both values are pointers
a := pointy.Int(1)
b := pointy.Int(1)
if pointy.PointersValueEqual(a, b) {
	fmt.Println("a and b contain equal dereferenced values")
}

// or if just one is a pointer
a := pointy.Int(1)
b := 1
if pointy.PointerValueEqual(a, b) {
	fmt.Println("a and b contain equal dereferenced values")
}
```

### New in release 1.1.0

Additional helper functions have been added to safely dereference pointers
or return a fallback value:

```golang
val := 42
pointerToVal := &val
// then later in your code..
myVal := pointy.IntValue(pointerToVal, 99) // returns 42 (or 99 if pointerToVal was nil)
```

## GoDoc

[https://godoc.org/github.com/openly-engineering/pointy](https://pkg.go.dev/github.com/openly-engineering/pointy)

## Installation

`go get go.openly.dev/pointy`

## Example

```golang
package main

import (
	"fmt"

	"go.openly.dev/pointy"
)

func main() {
	foo := pointy.Pointer(2018)
	fmt.Println("foo is a pointer to:", *foo)

	bar := pointy.Pointer("point to me")
	fmt.Println("bar is a pointer to:", *bar)

	// get the value back out (new in v1.1.0)
	barVal := pointy.PointerValue(bar, "empty!")
	fmt.Println("bar's value is:", barVal)
}
```

## Available Functions

`Pointer[T any](x T) *T`  
`PointerValue[T any](p *T, fallback T) T`  
`Bool(x bool) *bool`  
`BoolValue(p *bool, fallback bool) bool`  
`Byte(x byte) *byte`  
`ByteValue(p *byte, fallback byte) byte`  
`Complex128(x complex128) *complex128`  
`Complex128Value(p *complex128, fallback complex128) complex128`  
`Complex64(x complex64) *complex64`  
`Complex64Value(p *complex64, fallback complex64) complex64`  
`Float32(x float32) *float32`  
`Float32Value(p *float32, fallback float32) float32`  
`Float64(x float64) *float64`  
`Float64Value(p *float64, fallback float64) float64`  
`Int(x int) *int`  
`IntValue(p *int, fallback int) int`  
`Int8(x int8) *int8`  
`Int8Value(p *int8, fallback int8) int8`  
`Int16(x int16) *int16`  
`Int16Value(p *int16, fallback int16) int16`  
`Int32(x int32) *int32`  
`Int32Value(p *int32, fallback int32) int32`  
`Int64(x int64) *int64`  
`Int64Value(p *int64, fallback int64) int64`  
`Uint(x uint) *uint`  
`UintValue(p *uint, fallback uint) uint`  
`Uint8(x uint8) *uint8`  
`Uint8Value(p *uint8, fallback uint8) uint8`  
`Uint16(x uint16) *uint16`  
`Uint16Value(p *uint16, fallback uint16) uint16`  
`Uint32(x uint32) *uint32`  
`Uint32Value(p *uint32, fallback uint32) uint32`  
`Uint64(x uint64) *uint64`  
`Uint64Value(p *uint64, fallback uint64) uint64`  
`String(x string) *string`  
`StringValue(p *string, fallback string) string`  
`Rune(x rune) *rune`  
`RuneValue(p *rune, fallback rune) rune`  
`PointersValueEqual[T comparable](a *T, b *T) bool`  
`PointerValueEqual[T comparable](a *T, b T) bool`  
## Motivation

Creating pointers to literal constant values is useful, especially in unit tests. Go doesn't support simply using the address operator (&) to reference the location of e.g. `value := &int64(42)` so we're forced to [create](https://stackoverflow.com/questions/35146286/find-address-of-constant-in-go/35146856#35146856) [little](https://stackoverflow.com/questions/34197248/how-can-i-store-reference-to-the-result-of-an-operation-in-go/34197367#34197367) [workarounds](https://stackoverflow.com/questions/30716354/how-do-i-do-a-literal-int64-in-go/30716481#30716481). A common solution is to create a helper function:

```golang
func createInt64Pointer(x int64) *int64 {
    return &x
}
// now you can create a pointer to 42 inline
value := createInt64Pointer(42)
```

This package provides a library of these simple little helper functions for every native Go primitive.

Made @ Openly. [Join us](https://careers.openly.com/) and use Go to build cool stuff.
