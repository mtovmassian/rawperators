package pkg

const ZERO_INT64 = int64(0)
const ZERO_INT8 = int8(0)
const ONE_INT64 = int64(1)
const ONE_INT8 = int8(1)

type Bit struct {
	value int8
}

func (bit Bit) SetFromInt64(value int64) *Bit {
	bit.value = int8(
		Min(ONE_INT64, Max(ZERO_INT64, Abs(value))),
	)
	return &bit
}

func NewBit() *Bit {
	return &Bit{int8(ZERO_INT8)}
}
