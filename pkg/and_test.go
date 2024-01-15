package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnd__0_0__0(t *testing.T) {
	// GIVEN
	bit1 := NewBit().SetFalse()
	bit2 := NewBit().SetFalse()
	//
	bitAnd := And(bit1, bit2)
	// THEN
	assert.Equal(t, bitAnd.value, ZERO_INT64)
}

func TestAnd__0_1__0(t *testing.T) {
	// GIVEN
	bit1 := NewBit().SetFalse()
	bit2 := NewBit().SetTrue()
	//
	bitAnd := And(bit1, bit2)
	// THEN
	assert.Equal(t, bitAnd.value, ZERO_INT64)
}

func TestAnd__1_0__0(t *testing.T) {
	// GIVEN
	bit1 := NewBit().SetTrue()
	bit2 := NewBit().SetFalse()
	//
	bitAnd := And(bit1, bit2)
	// THEN
	assert.Equal(t, bitAnd.value, ZERO_INT64)
}

func TestAnd__1_1__1(t *testing.T) {
	// GIVEN
	bit1 := NewBit().SetTrue()
	bit2 := NewBit().SetTrue()
	//
	bitAnd := And(bit1, bit2)
	// THEN
	assert.Equal(t, bitAnd.value, ONE_INT64)
}
