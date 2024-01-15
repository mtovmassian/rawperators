package pkg

import "math"

const ZERO_INT64 = int64(0)
const ONE_INT64 = int64(1)

type Bit struct {
	value int64
}

func (bit Bit) SetFromInt64(value int64) *Bit {
	valueSquare := int64(math.Pow(float64(value), 2))
	bit.value = (valueSquare + 2) % (valueSquare + 1)
	return &bit
}

func (bit Bit) SetTrue() *Bit {
	bit.value = ONE_INT64
	return &bit
}

func (bit Bit) SetFalse() *Bit {
	bit.value = ZERO_INT64
	return &bit
}

func NewBit() *Bit {
	bit := Bit{}
	bit.SetFalse()
	return &bit
}
