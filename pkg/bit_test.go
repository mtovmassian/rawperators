package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBitSetFromInt64__0__0(t *testing.T) {
	// GIVEN
	zero := int64(0)
	// WHEN
	bit := NewBit().SetFromInt64(zero)
	// THEN
	assert.Equal(t, ZERO_INT8, bit.value)
}

func TestBitSetFromInt64__1__1(t *testing.T) {
	// GIVEN
	one := int64(1)
	// WHEN
	bit := NewBit().SetFromInt64(one)
	// THEN
	assert.Equal(t, ONE_INT8, bit.value)
}

func TestBitSetFromInt64__5__1(t *testing.T) {
	// GIVEN
	five := int64(5)
	// WHEN
	bit := NewBit().SetFromInt64(five)
	// THEN
	assert.Equal(t, ONE_INT8, bit.value)
}

func TestBitSetFromInt64__1000__1(t *testing.T) {
	// GIVEN
	thousand := int64(1000)
	// WHEN
	bit := NewBit().SetFromInt64(thousand)
	// THEN
	assert.Equal(t, ONE_INT8, bit.value)
}
