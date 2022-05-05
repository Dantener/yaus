package base62

import (
	"math"
)

const (
	alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	base     = uint32(len(alphabet))
)

func Encode(number uint32) string {
	encodedValue := make([]byte, 0)

	for number > 0 {
		rest := math.Mod(float64(number), float64(base))
		number /= base

		encodedValue = append([]byte{alphabet[int(rest)]}, encodedValue...)

	}

	return string(encodedValue)
}
