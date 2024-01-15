package pkg

func And(bit1 *Bit, bit2 *Bit) Bit {
	return *NewBit().SetFromInt64(
		bit1.value * bit2.value,
	)
}
