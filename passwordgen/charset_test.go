package passwordgen

import "testing"

func TestBuildCharSet(t *testing.T) {
	cfg := Config{
		UseDigits:    true,
		UseLowercase: true,
		UseUppercase: true,
		UseSymbols:   true,
	}

	cs := BuildCharSet(cfg)

	if len(cs) == 0 {
		t.Fatal("charset must not be empty")
	}

	if !containsAll(cs, Digits) {
		t.Fatal("digits missing")
	}
	if !containsAll(cs, Lowercase) {
		t.Fatal("lowercase missing")
	}
	if !containsAll(cs, Uppercase) {
		t.Fatal("uppercase missing")
	}
	if !containsAll(cs, Symbols) {
		t.Fatal("symbols missing")
	}
}

func containsAll(src, chars string) bool {
	for _, c := range chars {
		if !containsRune(src, c) {
			return false
		}
	}
	return true
}

func containsRune(s string, r rune) bool {
	for _, c := range s {
		if c == r {
			return true
		}
	}
	return false
}
