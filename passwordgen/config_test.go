package passwordgen

import "testing"

func TestDefaultConfig(t *testing.T) {
	cfg := DefaultConfig()

	if cfg.Length != DefaultLength {
		t.Fatalf("expected default length %d, got %d", DefaultLength, cfg.Length)
	}
	if !cfg.UseDigits || !cfg.UseLowercase {
		t.Fatal("default digits and lowercase must be enabled")
	}
	if cfg.UseUppercase || cfg.UseSymbols {
		t.Fatal("uppercase and symbols must be disabled by default")
	}
}

func TestValidateConfigInvalidLength(t *testing.T) {
	cfg := Config{Length: 2, UseDigits: true}
	err := ValidateConfig(cfg)

	if err == nil {
		t.Fatal("expected ErrInvalidLength")
	}
	if err != ErrInvalidLength {
		t.Fatalf("expected ErrInvalidLength, got %v", err)
	}
}

func TestValidateConfigNoCharset(t *testing.T) {
	cfg := Config{Length: 8}
	err := ValidateConfig(cfg)

	if err != ErrNoCharacterSet {
		t.Fatalf("expected ErrNoCharacterSet, got %v", err)
	}
}

func TestValidateConfigValid(t *testing.T) {
	cfg := Config{Length: 8, UseDigits: true}
	err := ValidateConfig(cfg)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
