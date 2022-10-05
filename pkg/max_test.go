package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMax(t *testing.T) {
	// GIVEN
	eightySeven := int64(87)
	ninetyNine := int64(99)
	// THEN
	assert.Equal(t, ninetyNine, Max(eightySeven, ninetyNine))
	assert.Equal(t, ninetyNine, Max(ninetyNine, eightySeven))
	assert.Equal(t, eightySeven, Max(eightySeven, eightySeven))
	assert.Equal(t, ninetyNine, Max(ninetyNine, ninetyNine))
	assert.Equal(t, -eightySeven, Max(-eightySeven, -ninetyNine))
	assert.Equal(t, -eightySeven, Max(-ninetyNine, -eightySeven))
}
