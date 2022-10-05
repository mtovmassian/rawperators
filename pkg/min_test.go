package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMin(t *testing.T) {
	// GIVEN
	eightySeven := int64(87)
	ninetyNine := int64(99)
	// THEN
	assert.Equal(t, Min(eightySeven, ninetyNine), eightySeven)
	assert.Equal(t, Min(ninetyNine, eightySeven), eightySeven)
	assert.Equal(t, Min(eightySeven, eightySeven), eightySeven)
	assert.Equal(t, Min(ninetyNine, ninetyNine), ninetyNine)
	assert.Equal(t, Min(-eightySeven, -ninetyNine), -ninetyNine)
	assert.Equal(t, Min(-ninetyNine, -eightySeven), -ninetyNine)
}
