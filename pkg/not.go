package pkg

func Not(bit *Bit) Bit {
	return *NewBit().SetFromInt64(
		1 - bit.value,
	)
}
