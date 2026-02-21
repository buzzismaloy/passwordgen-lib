package passwordgen

const (
	Digits    = "0123456789"
	Lowercase = "abcdefghijklmnopqrstuvwxyz"
	Uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Symbols   = "@#&-_+!$%^*()[]{}"
)

func BuildCharSet(cfg Config) string {
	charset := ""

	if cfg.UseDigits {
		charset += Digits
	}

	if cfg.UseLowercase {
		charset += Lowercase
	}

	if cfg.UseUppercase {
		charset += Uppercase
	}

	if cfg.UseSymbols {
		charset += Symbols
	}

	return charset
}
