package pkg

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNot__0__1(t *testing.T) {
	// GIVEN
	bit := NewBit().SetTrue()
	// WHEN
	notBit := Not(bit)
	// THEN
	assert.Equal(t, notBit.value, ZERO_INT64)
}

func TestNot__1__0(t *testing.T) {
	// GIVEN
	bit := NewBit().SetFalse()
	fmt.Printf("%d", bit.value)
	// WHEN
	notBit := Not(bit)
	// THEN
	assert.Equal(t, notBit.value, ONE_INT64)
}
