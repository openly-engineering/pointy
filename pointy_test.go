package pointy_test

import (
	"fmt"
	"testing"

	"github.com/mwielbut/pointy"
	"github.com/stretchr/testify/assert"
)

// This example returns a pointer to a variable holding the `int64` constant `2018`.
func ExampleInt64() {
	foo := pointy.Int64(2018)
	fmt.Println("foo contains value:", *foo)
}

// This example returns a pointer to a variable holding the `string` constant
// "point to me".
func ExampleString() {
	bar := pointy.String("point to me")
	fmt.Println("bar contains value:", *bar)
}

func TestBool(t *testing.T) {
	var value bool = true

	result := pointy.Bool(true)

	assert.NotNil(t, result)
	assert.Exactly(t, &value, result)
	assert.Exactly(t, value, *result)
}

func TestByte(t *testing.T) {
	var value byte = 'a'

	result := pointy.Byte('a')

	assert.NotNil(t, result)
	assert.Exactly(t, &value, result)
	assert.Exactly(t, value, *result)
}

func TestComplex128(t *testing.T) {
	var value complex128 = 42i

	result := pointy.Complex128(42i)

	assert.NotNil(t, result)
	assert.Exactly(t, &value, result)
	assert.Exactly(t, value, *result)
}

func TestComplex64(t *testing.T) {
	var value complex64 = 42i

	result := pointy.Complex64(42i)

	assert.NotNil(t, result)
	assert.Exactly(t, &value, result)
	assert.Exactly(t, value, *result)
}

func TestFloat32(t *testing.T) {
	var value float32 = 42.42

	result := pointy.Float32(42.42)

	assert.NotNil(t, result)
	assert.Exactly(t, &value, result)
	assert.Exactly(t, value, *result)
}

func TestFloat64(t *testing.T) {
	var value float64 = 42.42

	result := pointy.Float64(42.42)

	assert.NotNil(t, result)
	assert.Exactly(t, &value, result)
	assert.Exactly(t, value, *result)
}
func TestInt(t *testing.T) {
	var value int = 42

	result := pointy.Int(42)

	assert.NotNil(t, result)
	assert.Exactly(t, &value, result)
	assert.Exactly(t, value, *result)
}

func TestInt8(t *testing.T) {
	var value int8 = 42

	result := pointy.Int8(42)

	assert.NotNil(t, result)
	assert.Exactly(t, &value, result)
	assert.Exactly(t, value, *result)
}

func TestInt16(t *testing.T) {
	var value int16 = 42

	result := pointy.Int16(42)

	assert.NotNil(t, result)
	assert.Exactly(t, &value, result)
	assert.Exactly(t, value, *result)
}

func TestInt32(t *testing.T) {
	var value int32 = 42

	result := pointy.Int32(42)

	assert.NotNil(t, result)
	assert.Exactly(t, &value, result)
	assert.Exactly(t, value, *result)
}

func TestInt64(t *testing.T) {
	var value int64 = 42

	result := pointy.Int64(42)

	assert.NotNil(t, result)
	assert.Exactly(t, &value, result)
	assert.Exactly(t, value, *result)
}
func TestUint(t *testing.T) {
	var value uint = 42

	result := pointy.Uint(42)

	assert.NotNil(t, result)
	assert.Exactly(t, &value, result)
	assert.Exactly(t, value, *result)
}

func TestUint8(t *testing.T) {
	var value uint8 = 42

	result := pointy.Uint8(42)

	assert.NotNil(t, result)
	assert.Exactly(t, &value, result)
	assert.Exactly(t, value, *result)
}

func TestUint16(t *testing.T) {
	var value uint16 = 42

	result := pointy.Uint16(42)

	assert.NotNil(t, result)
	assert.Exactly(t, &value, result)
	assert.Exactly(t, value, *result)
}

func TestUint32(t *testing.T) {
	var value uint32 = 42

	result := pointy.Uint32(42)

	assert.NotNil(t, result)
	assert.Exactly(t, &value, result)
	assert.Exactly(t, value, *result)
}

func TestUint64(t *testing.T) {
	var value uint64 = 42

	result := pointy.Uint64(42)

	assert.NotNil(t, result)
	assert.Exactly(t, &value, result)
	assert.Exactly(t, value, *result)
}
func TestString(t *testing.T) {
	var value string = "foo"

	result := pointy.String("foo")

	assert.NotNil(t, result)
	assert.Exactly(t, &value, result)
	assert.Exactly(t, value, *result)
}
func TestRune(t *testing.T) {
	var value rune = 'a'

	result := pointy.Rune('a')

	assert.NotNil(t, result)
	assert.Exactly(t, &value, result)
	assert.Exactly(t, value, *result)
}
