package pointy_test

import (
	"fmt"
	"testing"

	"github.com/openlyinc/pointy"
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

// This example uses the new generic construct to returns a pointer to a variable
// holding the `string` constant "point to me".
func ExamplePointer() {
	foo := pointy.Pointer("point to me")
	fmt.Println("foo contains value:", *foo)
}

func TestBool(t *testing.T) {
	var value bool = true
	var fallback bool = false

	result := pointy.Bool(true)

	assert.NotNil(t, result)
	assert.Exactly(t, &value, result)
	assert.Exactly(t, value, *result)

	assert.Exactly(t, value, pointy.BoolValue(result, fallback))
	assert.Exactly(t, fallback, pointy.BoolValue(nil, fallback))
}

func TestByte(t *testing.T) {
	var value byte = 'a'
	var fallback byte = 'b'

	result := pointy.Byte('a')

	assert.NotNil(t, result)
	assert.Exactly(t, &value, result)
	assert.Exactly(t, value, *result)

	assert.Exactly(t, value, pointy.ByteValue(result, fallback))
	assert.Exactly(t, fallback, pointy.ByteValue(nil, fallback))
}

func TestComplex128(t *testing.T) {
	var value complex128 = 42i
	var fallback complex128 = 10i

	result := pointy.Complex128(42i)

	assert.NotNil(t, result)
	assert.Exactly(t, &value, result)
	assert.Exactly(t, value, *result)

	assert.Exactly(t, value, pointy.Complex128Value(result, fallback))
	assert.Exactly(t, fallback, pointy.Complex128Value(nil, fallback))
}

func TestComplex64(t *testing.T) {
	var value complex64 = 42i
	var fallback complex64 = 10i

	result := pointy.Complex64(42i)

	assert.NotNil(t, result)
	assert.Exactly(t, &value, result)
	assert.Exactly(t, value, *result)

	assert.Exactly(t, value, pointy.Complex64Value(result, fallback))
	assert.Exactly(t, fallback, pointy.Complex64Value(nil, fallback))
}

func TestFloat32(t *testing.T) {
	var value float32 = 42.42
	var fallback float32 = 83.12

	result := pointy.Float32(42.42)

	assert.NotNil(t, result)
	assert.Exactly(t, &value, result)
	assert.Exactly(t, value, *result)

	assert.Exactly(t, value, pointy.Float32Value(result, fallback))
	assert.Exactly(t, fallback, pointy.Float32Value(nil, fallback))
}

func TestFloat64(t *testing.T) {
	var value float64 = 42.42
	var fallback float64 = 83.12

	result := pointy.Float64(42.42)

	assert.NotNil(t, result)
	assert.Exactly(t, &value, result)
	assert.Exactly(t, value, *result)

	assert.Exactly(t, value, pointy.Float64Value(result, fallback))
	assert.Exactly(t, fallback, pointy.Float64Value(nil, fallback))
}
func TestInt(t *testing.T) {
	var value int = 42
	var fallback int = 83

	result := pointy.Int(42)

	assert.NotNil(t, result)
	assert.Exactly(t, &value, result)
	assert.Exactly(t, value, *result)

	assert.Exactly(t, value, pointy.IntValue(result, fallback))
	assert.Exactly(t, fallback, pointy.IntValue(nil, fallback))
}

func TestInt8(t *testing.T) {
	var value int8 = 42
	var fallback int8 = 83

	result := pointy.Int8(42)

	assert.NotNil(t, result)
	assert.Exactly(t, &value, result)
	assert.Exactly(t, value, *result)

	assert.Exactly(t, value, pointy.Int8Value(result, fallback))
	assert.Exactly(t, fallback, pointy.Int8Value(nil, fallback))
}

func TestInt16(t *testing.T) {
	var value int16 = 42
	var fallback int16 = 83

	result := pointy.Int16(42)

	assert.NotNil(t, result)
	assert.Exactly(t, &value, result)
	assert.Exactly(t, value, *result)

	assert.Exactly(t, value, pointy.Int16Value(result, fallback))
	assert.Exactly(t, fallback, pointy.Int16Value(nil, fallback))
}

func TestInt32(t *testing.T) {
	var value int32 = 42
	var fallback int32 = 83

	result := pointy.Int32(42)

	assert.NotNil(t, result)
	assert.Exactly(t, &value, result)
	assert.Exactly(t, value, *result)

	assert.Exactly(t, value, pointy.Int32Value(result, fallback))
	assert.Exactly(t, fallback, pointy.Int32Value(nil, fallback))
}

func TestInt64(t *testing.T) {
	var value int64 = 42
	var fallback int64 = 83

	result := pointy.Int64(42)

	assert.NotNil(t, result)
	assert.Exactly(t, &value, result)
	assert.Exactly(t, value, *result)

	assert.Exactly(t, value, pointy.Int64Value(result, fallback))
	assert.Exactly(t, fallback, pointy.Int64Value(nil, fallback))
}

func TestUint(t *testing.T) {
	var value uint = 42
	var fallback uint = 83

	result := pointy.Uint(42)

	assert.NotNil(t, result)
	assert.Exactly(t, &value, result)
	assert.Exactly(t, value, *result)

	assert.Exactly(t, value, pointy.UintValue(result, fallback))
	assert.Exactly(t, fallback, pointy.UintValue(nil, fallback))
}

func TestUint8(t *testing.T) {
	var value uint8 = 42
	var fallback uint8 = 83

	result := pointy.Uint8(42)

	assert.NotNil(t, result)
	assert.Exactly(t, &value, result)
	assert.Exactly(t, value, *result)

	assert.Exactly(t, value, pointy.Uint8Value(result, fallback))
	assert.Exactly(t, fallback, pointy.Uint8Value(nil, fallback))
}

func TestUint16(t *testing.T) {
	var value uint16 = 42
	var fallback uint16 = 83

	result := pointy.Uint16(42)

	assert.NotNil(t, result)
	assert.Exactly(t, &value, result)
	assert.Exactly(t, value, *result)

	assert.Exactly(t, value, pointy.Uint16Value(result, fallback))
	assert.Exactly(t, fallback, pointy.Uint16Value(nil, fallback))
}

func TestUint32(t *testing.T) {
	var value uint32 = 42
	var fallback uint32 = 83

	result := pointy.Uint32(42)

	assert.NotNil(t, result)
	assert.Exactly(t, &value, result)
	assert.Exactly(t, value, *result)

	assert.Exactly(t, value, pointy.Uint32Value(result, fallback))
	assert.Exactly(t, fallback, pointy.Uint32Value(nil, fallback))
}

func TestUint64(t *testing.T) {
	var value uint64 = 42
	var fallback uint64 = 83

	result := pointy.Uint64(42)

	assert.NotNil(t, result)
	assert.Exactly(t, &value, result)
	assert.Exactly(t, value, *result)

	assert.Exactly(t, value, pointy.Uint64Value(result, fallback))
	assert.Exactly(t, fallback, pointy.Uint64Value(nil, fallback))
}

func TestString(t *testing.T) {
	var value string = "foo"
	var fallback string = "bar"

	result := pointy.String("foo")

	assert.NotNil(t, result)
	assert.Exactly(t, &value, result)
	assert.Exactly(t, value, *result)

	assert.Exactly(t, value, pointy.StringValue(result, fallback))
	assert.Exactly(t, fallback, pointy.StringValue(nil, fallback))
}

func TestRune(t *testing.T) {
	var value rune = 'a'
	var fallback rune = 'b'

	result := pointy.Rune('a')

	assert.NotNil(t, result)
	assert.Exactly(t, &value, result)
	assert.Exactly(t, value, *result)

	assert.Exactly(t, value, pointy.RuneValue(result, fallback))
	assert.Exactly(t, fallback, pointy.RuneValue(nil, fallback))
}

func TestPointer_bool(t *testing.T) {
	var value bool = true
	var fallback bool = false

	result := pointy.Pointer(true)

	assert.NotNil(t, result)
	assert.Exactly(t, &value, result)
	assert.Exactly(t, value, *result)

	assert.Exactly(t, value, pointy.PointerValue(result, fallback))
	assert.Exactly(t, fallback, pointy.PointerValue(nil, fallback))
}

func TestPointer_String(t *testing.T) {
	var value string = "foo"
	var fallback string = "bar"

	result := pointy.Pointer("foo")

	assert.NotNil(t, result)
	assert.Exactly(t, &value, result)
	assert.Exactly(t, value, *result)

	assert.Exactly(t, value, pointy.PointerValue(result, fallback))
	assert.Exactly(t, fallback, pointy.PointerValue(nil, fallback))
}

func TestPointer_Int(t *testing.T) {
	var value int = 42
	var fallback int = 83

	result := pointy.Pointer(42)

	assert.NotNil(t, result)
	assert.Exactly(t, &value, result)
	assert.Exactly(t, value, *result)

	assert.Exactly(t, value, pointy.PointerValue(result, fallback))
	assert.Exactly(t, fallback, pointy.PointerValue(nil, fallback))
}
