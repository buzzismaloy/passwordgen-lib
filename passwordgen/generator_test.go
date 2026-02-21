package passwordgen

import (
	"errors"
	"testing"

	"github.com/buzzismaloy/passwordgen-lib/internal/rand"
)

func TestGeneratorLength(t *testing.T) {
	cfg := Config{
		Length:       16,
		UseDigits:    true,
		UseLowercase: true,
	}

	mock := rand.NewMockSource([]int{1, 2, 3, 4, 5})
	gen, _ := NewGenerator(cfg, mock)

	pass, err := gen.Generate()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(pass) != cfg.Length {
		t.Fatalf("expected length %d, got %d", cfg.Length, len(pass))
	}
}

func TestGeneratorIncludesClasses(t *testing.T) {
	cfg := Config{
		Length:       10,
		UseDigits:    true,
		UseLowercase: true,
	}

	mock := rand.NewMockSource([]int{0, 1, 2, 3, 4})
	gen, _ := NewGenerator(cfg, mock)

	pass, _ := gen.Generate()

	if !hasAny(pass, Digits) {
		t.Fatal("password must contain digit")
	}
	if !hasAny(pass, Lowercase) {
		t.Fatal("password must contain lowercase")
	}
}

func TestGeneratorErrorPropagation(t *testing.T) {
	cfg := Config{
		Length:       8,
		UseDigits:    true,
		UseLowercase: true,
	}

	// RNG that always fails
	bad := &FailSource{}
	gen, _ := NewGenerator(cfg, bad)

	_, err := gen.Generate()
	if err == nil {
		t.Fatal("expected error from RNG")
	}
}

func hasAny(s, charset string) bool {
	for _, c := range s {
		for _, x := range charset {
			if c == x {
				return true
			}
		}
	}
	return false
}

type FailSource struct{}

func (FailSource) Intn(n int) (int, error) {
	return 0, errors.New("rng failure")
}
