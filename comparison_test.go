package pointy_test

import (
	"fmt"
	"testing"

	"github.com/openlyinc/pointy"
	"github.com/stretchr/testify/assert"
)

// This example demonstrates comparing two pointers for dereferenced value equality.
func ExamplePointersValueEqual() {
	a := pointy.Int(1)
	b := pointy.Int(2)
	result := pointy.PointersValueEqual(a, b)
	fmt.Printf("comparing value quality returns: %v\n", result)
}

func TestPointersValueEqual_float64(t *testing.T) {

	a := pointy.Float64(1.0)
	b := pointy.Float64(2.0)

	assert.False(t, pointy.PointersValueEqual(a, b), "expected both values to be not equal")
	assert.False(t, pointy.PointersValueEqual(a, nil), "expected both values to be not equal")
	assert.False(t, pointy.PointersValueEqual(nil, b), "expected both values to be not equal")
	var aNil, bNil *float64
	assert.True(t, pointy.PointersValueEqual(aNil, bNil), "expected nil values to be equal")

	a1 := pointy.Float64(1.0)
	assert.True(t, pointy.PointersValueEqual(a, a1), "expected both values to be equal")
}

func TestPointersValueEqual_string(t *testing.T) {

	a := pointy.String("a")
	b := pointy.String("b")

	assert.False(t, pointy.PointersValueEqual(a, b), "expected both values to be not equal")
	assert.False(t, pointy.PointersValueEqual(a, nil), "expected both values to be not equal")
	assert.False(t, pointy.PointersValueEqual(nil, b), "expected both values to be not equal")
	var aNil, bNil *float64
	assert.True(t, pointy.PointersValueEqual(aNil, bNil), "expected nil values to be equal")

	a1 := pointy.String("a")
	assert.True(t, pointy.PointersValueEqual(a, a1), "expected both values to be equal")
}
