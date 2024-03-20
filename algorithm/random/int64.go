package randomutil

import (
	"crypto/rand"
	"math/big"
)

func RandInt64(min, max int64) int64 {

	maxBigInt := big.NewInt(max)

	i, _ := rand.Int(rand.Reader, maxBigInt)

	iInt64 := i.Int64()
	if iInt64 < min {
		iInt64 = RandInt64(min, max)
	}
	return iInt64
}
