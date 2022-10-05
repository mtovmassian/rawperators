package pkg

func Min(numberA int64, numberB int64) int64 {
	return int64(
		0.5 * float64(
			(numberA+numberB)-Abs(numberA-numberB),
		),
	)
}
