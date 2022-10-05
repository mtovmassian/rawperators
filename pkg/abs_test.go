package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbs(t *testing.T) {
	// GIVEN
	zero := int64(0)
	one := int64(1)
	twelve := int64(12)
	// THEN
	assert.Equal(t, zero, Abs(zero))
	assert.Equal(t, zero, Abs(-zero))
	assert.Equal(t, one, Abs(one))
	assert.Equal(t, one, Abs(-one))
	assert.Equal(t, int64(twelve), Abs(twelve))
	assert.Equal(t, int64(twelve), Abs(-twelve))
}
