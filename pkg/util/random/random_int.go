package random

import (
	"math"
	"math/rand/v2"
)

func GetRandomInt(len int) int {
	return rand.IntN(9 * int(math.Pow(10, float64(len-1)))) + int(math.Pow(10, float64(len-1)))
}