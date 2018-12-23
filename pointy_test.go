package pointy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBool(t *testing.T) {
	var value bool = true

	result := Bool(true)

	assert.NotNil(t, result)
	assert.Exactly(t, &value, result)
	assert.Exactly(t, value, *result)
}

func TestByte(t *testing.T) {
	var value byte = 'a'

	result := Byte('a')

	assert.NotNil(t, result)
	assert.Exactly(t, &value, result)
	assert.Exactly(t, value, *result)
}

func TestComplex128(t *testing.T) {
	var value complex128 = 42i

	result := Complex128(42i)

	assert.NotNil(t, result)
	assert.Exactly(t, &value, result)
	assert.Exactly(t, value, *result)
}

func TestComplex64(t *testing.T) {
	var value complex64 = 42i

	result := Complex64(42i)

	assert.NotNil(t, result)
	assert.Exactly(t, &value, result)
	assert.Exactly(t, value, *result)
}

func TestFloat32(t *testing.T) {
	var value float32 = 42.42

	result := Float32(42.42)

	assert.NotNil(t, result)
	assert.Exactly(t, &value, result)
	assert.Exactly(t, value, *result)
}

func TestFloat64(t *testing.T) {
	var value float64 = 42.42

	result := Float64(42.42)

	assert.NotNil(t, result)
	assert.Exactly(t, &value, result)
	assert.Exactly(t, value, *result)
}
func TestInt(t *testing.T) {
	var value int = 42

	result := Int(42)

	assert.NotNil(t, result)
	assert.Exactly(t, &value, result)
	assert.Exactly(t, value, *result)
}

func TestInt8(t *testing.T) {
	var value int8 = 42

	result := Int8(42)

	assert.NotNil(t, result)
	assert.Exactly(t, &value, result)
	assert.Exactly(t, value, *result)
}

func TestInt16(t *testing.T) {
	var value int16 = 42

	result := Int16(42)

	assert.NotNil(t, result)
	assert.Exactly(t, &value, result)
	assert.Exactly(t, value, *result)
}

func TestInt32(t *testing.T) {
	var value int32 = 42

	result := Int32(42)

	assert.NotNil(t, result)
	assert.Exactly(t, &value, result)
	assert.Exactly(t, value, *result)
}

func TestInt64(t *testing.T) {
	var value int64 = 42

	result := Int64(42)

	assert.NotNil(t, result)
	assert.Exactly(t, &value, result)
	assert.Exactly(t, value, *result)
}
func TestUint(t *testing.T) {
	var value uint = 42

	result := Uint(42)

	assert.NotNil(t, result)
	assert.Exactly(t, &value, result)
	assert.Exactly(t, value, *result)
}

func TestUint8(t *testing.T) {
	var value uint8 = 42

	result := Uint8(42)

	assert.NotNil(t, result)
	assert.Exactly(t, &value, result)
	assert.Exactly(t, value, *result)
}

func TestUint16(t *testing.T) {
	var value uint16 = 42

	result := Uint16(42)

	assert.NotNil(t, result)
	assert.Exactly(t, &value, result)
	assert.Exactly(t, value, *result)
}

func TestUint32(t *testing.T) {
	var value uint32 = 42

	result := Uint32(42)

	assert.NotNil(t, result)
	assert.Exactly(t, &value, result)
	assert.Exactly(t, value, *result)
}

func TestUint64(t *testing.T) {
	var value uint64 = 42

	result := Uint64(42)

	assert.NotNil(t, result)
	assert.Exactly(t, &value, result)
	assert.Exactly(t, value, *result)
}
func TestString(t *testing.T) {
	var value string = "foo"

	result := String("foo")

	assert.NotNil(t, result)
	assert.Exactly(t, &value, result)
	assert.Exactly(t, value, *result)
}
func TestRune(t *testing.T) {
	var value rune = 'a'

	result := Rune('a')

	assert.NotNil(t, result)
	assert.Exactly(t, &value, result)
	assert.Exactly(t, value, *result)
}
