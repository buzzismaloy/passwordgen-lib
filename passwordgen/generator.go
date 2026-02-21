package passwordgen

import (
	"github.com/buzzismaloy/passwordgen-lib/internal/rand"
)

type Generator struct {
	cfg     Config
	rng     rand.RandomSource
	charset string
}

func NewGenerator(cfg Config, rngSrc rand.RandomSource) (*Generator, error) {
	if rngSrc == nil {
		rngSrc = rand.CryptoSource{}
	}

	if err := ValidateConfig(cfg); err != nil {
		return nil, err
	}

	charset := BuildCharSet(cfg)

	return &Generator{
		cfg:     cfg,
		rng:     rngSrc,
		charset: charset,
	}, nil
}

func randomChar(charset string, rng rand.RandomSource) (byte, error) {
	idx, err := rng.Intn(len(charset))

	if err != nil {
		return 0, err
	}

	return charset[idx], nil
}

func (g *Generator) Generate() (string, error) {
	var password []byte

	if g.cfg.UseDigits {
		c, err := randomChar(Digits, g.rng)
		if err != nil {
			return "", err
		}
		password = append(password, c)
	}

	if g.cfg.UseLowercase {
		c, err := randomChar(Lowercase, g.rng)
		if err != nil {
			return "", err
		}
		password = append(password, c)
	}

	if g.cfg.UseUppercase {
		c, err := randomChar(Uppercase, g.rng)
		if err != nil {
			return "", err
		}
		password = append(password, c)
	}

	if g.cfg.UseSymbols {
		c, err := randomChar(Symbols, g.rng)
		if err != nil {
			return "", err
		}
		password = append(password, c)
	}

	// Fill remaining length
	for len(password) < g.cfg.Length {
		c, err := randomChar(g.charset, g.rng)
		if err != nil {
			return "", err
		}
		password = append(password, c)
	}

	// Shuffle
	for i := len(password) - 1; i > 0; i-- {
		j, err := g.rng.Intn(i + 1)
		if err != nil {
			return "", err
		}
		password[i], password[j] = password[j], password[i]
	}

	return string(password), nil
}
