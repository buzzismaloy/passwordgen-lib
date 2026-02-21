package rand

import "testing"

func TestCryptoSourceIntn(t *testing.T) {
	cs := CryptoSource{}
	v, err := cs.Intn(10)

	if err != nil {
		t.Fatal(err)
	}
	if v < 0 || v >= 10 {
		t.Fatalf("out of bounds: %d", v)
	}
}
