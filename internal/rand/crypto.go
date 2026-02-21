package rand

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
)

var ErrCryptoInvalidBound = errors.New("[ERROR]: invalid bound for crypto generation")

type RandomSource interface {
	Intn(n int) (int, error)
}

type CryptoSource struct{}

func (CryptoSource) Intn(n int) (int, error) {
	if n <= 0 {
		return 0, fmt.Errorf("Crypto failure: %w", ErrCryptoInvalidBound)
	}
	max := big.NewInt(int64(n))

	v, err := rand.Int(rand.Reader, max)
	if err != nil {
		return 0, err
	}

	return int(v.Int64()), nil
}
