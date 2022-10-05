package pkg

import (
	"math"
)

func Abs(number int64) int64 {
	return int64(
		math.Sqrt(
			math.Pow(float64(number), 2),
		),
	)
}
