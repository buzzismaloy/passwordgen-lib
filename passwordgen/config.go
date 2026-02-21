package passwordgen

import "errors"

var ErrInvalidLength = errors.New("[ERROR]: Invalid length of password.")
var ErrNoCharacterSet = errors.New("[ERROR]: no characters sets are used to generate password.")

// package constants that define the min and max password length
const (
	MinPassLength = 4
	MaxPassLength = 128
	DefaultLength = 8
)

// Config - structure that contains configs
type Config struct {
	Length       int
	UseDigits    bool
	UseLowercase bool
	UseUppercase bool
	UseSymbols   bool
}

// DefaultConfig -  function that returns defaul configs
func DefaultConfig() *Config {
	return &Config{
		Length:       DefaultLength,
		UseDigits:    true,
		UseLowercase: true,
		UseUppercase: false,
		UseSymbols:   false,
	}
}

// Option - type of function-option
type Option func(*Config)

// WithLength - function that sets the length in configuration structure
func WithLength(l int) Option {
	return func(c *Config) {
		/*if l < MinPassLength {
			l = MinPassLength
		}
		if l > MaxPassLength {
			l = MaxPassLength
		}*/
		c.Length = l
	}
}

// WithDigits - function that sets the use of digits in configuration
func WithDigits(f bool) Option {
	return func(c *Config) {
		c.UseDigits = f
	}
}

// WithLowercase - function that sets the use of lowercase letters in configuration
func WithLowercase(f bool) Option {
	return func(c *Config) {
		c.UseLowercase = f
	}
}

// WithUppercase - function that sets the use of uppercase letters in configuration
func WithUppercase(f bool) Option {
	return func(c *Config) {
		c.UseUppercase = f
	}
}

// WithSymbols - function that sets the use of special symbols in configuration
func WithSymbols(f bool) Option {
	return func(c *Config) {
		c.UseSymbols = f
	}
}

// NewConfig - function that creates new configuration with received options and returns it
func NewConfig(opts ...Option) *Config {
	cfg := DefaultConfig()
	for _, opt := range opts {
		opt(cfg)
	}
	return cfg
}

// ValidateConfig - function that validates the length of password and if any set of characters is used
// if length doesn't fit in the borders - the default length(which equals to 8) is set
func ValidateConfig(cfg Config) error {
	if cfg.Length < MinPassLength || cfg.Length > MaxPassLength {
		return ErrInvalidLength
	}

	if !cfg.UseDigits && !cfg.UseLowercase && !cfg.UseUppercase && !cfg.UseSymbols {
		return ErrNoCharacterSet
	}

	return nil
}
