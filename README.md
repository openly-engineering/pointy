# pointy

Simple helper functions to provide a shorthand to get a pointer to a variable holding a constant...because it's annoying when you have to do it hundreds of times in unit tests:

```golang
val := 42
pointerToVal := &val
// vs.
pointerToVal := pointy.Int(42)
```

## GoDoc

https://godoc.org/github.com/mwielbut/pointy

## Installation

`go get github.com/mwielbut/pointy`

## Example

```golang
package main

import (
	"fmt"

	"github.com/mwielbut/pointy"
)

func main() {
	foo := pointy.Int64(2018)
	fmt.Println("foo is a pointer to:", *foo)

	bar := pointy.String("point to me")
	fmt.Println("bar is a pointer to:", *bar)
}
```

## Available Functions

`pointy.Bool(true)`  
`pointy.Byte('a')`  
`pointy.Complex64(42i)`  
`pointy.Complex128(42i)`  
`pointy.Float32(42.0)`  
`pointy.Float64(42.0)`  
`pointy.Int(42)`  
`pointy.Int8(42)`  
`pointy.Int16(42)`  
`pointy.Int32(42)`  
`pointy.Int64(42)`  
`pointy.Uint(42)`  
`pointy.Uint8(42)`  
`pointy.Uint16(42)`  
`pointy.Uint32(42)`  
`pointy.Uint64(42)`  
`pointy.String("foo")`  
`pointy.Rune('a')`

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
