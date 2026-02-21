# passwordgen-lib

A lightweight Go library for generating **cryptographically secure passwords** with flexible configuration and pluggable randomness sources.

This project is designed as a reusable library with strong defaults, extensibility, and full unit test coverage.

---

## Features

- Configurable password length
- Enable/disable character classes:
  - digits
  - lowercase letters
  - uppercase letters
  - symbols
- Cryptographically secure randomness using `crypto/rand`
- Pluggable RNG interface (ideal for testing and custom entropy sources)
- Guaranteed inclusion of at least one character from each enabled set
- Full unit test suite
- Clean functional options API

---

## Installation

### Method 1: Using `go get`

```bash
go get github.com/buzzismaloy/passwordgen-lib
```

### Method 2: Using `go mod edit`

```bash
go mod edit -require=github.com/buzzismaloy/passwordgen-lib@latest
go mod tidy
```

### Method 3: Manually editing `go.mod`

Add this line to your `go.mod` file:

```
require github.com/buzzismaloy/passwordgen-lib v0.0.0
```

Then run:

```bash
go mod tidy
```

---

## Quick Start

```go
package main

import (
	"fmt"
	"log"

	passwordgen "github.com/buzzismaloy/passwordgen-lib/passwordgen"
)

func main() {
	cfg := passwordgen.NewConfig(
		passwordgen.WithLength(16),
		passwordgen.WithDigits(true),
		passwordgen.WithLowercase(true),
		passwordgen.WithUppercase(true),
		passwordgen.WithSymbols(true),
	)

	gen, err := passwordgen.NewGenerator(*cfg, nil)
	if err != nil {
		log.Fatal(err)
	}

	pass, err := gen.Generate()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(pass)
}
```

---

## Example Output

```
gA9@f3Q!2xP7b$
```

---

## Configuration

### Config Structure

```go
type Config struct {
	Length       int
	UseDigits    bool
	UseLowercase bool
	UseUppercase bool
	UseSymbols   bool
}
```

### Default Configuration

```
Length:       8
Digits:       enabled
Lowercase:    enabled
Uppercase:    disabled
Symbols:      disabled
```

### Functional Options API

```go
passwordgen.WithLength(int)
passwordgen.WithDigits(bool)
passwordgen.WithLowercase(bool)
passwordgen.WithUppercase(bool)
passwordgen.WithSymbols(bool)
```

Example :

```go
cfg := passwordgen.NewConfig(
	passwordgen.WithLength(32),
	passwordgen.WithUppercase(true),
	passwordgen.WithSymbols(true),
)
```

### Configuration Validation

The library validates configuration before password generation.

**Validation Rules**:

- Password length must be 4 ≤ length ≤ 128
- At least one character class must be enabled

**Possible Errors**:

```go
ErrInvalidLength
ErrNoCharacterSet
```

## Password Generator

```go
gen, err := passwordgen.NewGenerator(cfg, nil)
password, err := gen.Generate()
```

Generator Guarantees:

- At least one character from each enabled class

- Remaining characters are randomly selected

- Cryptographically secure shuffle of the final password

## Random Sources

### Default RNG (crypto/rand)

```go
gen, _ := passwordgen.NewGenerator(cfg, nil)
```

### Custom RNG (useful for tests)

```go
import "github.com/buzzismaloy/passwordgen-lib/internal/rand"

mock := rand.NewMockSource([]int{1, 2, 3, 4})
gen, _ := passwordgen.NewGenerator(cfg, mock)
```

### Random Source Interface

```go
type RandomSource interface {
	Intn(n int) (int, error)
}
```

You can implement your own entropy source (e.g., deterministic RNG, hardware RNG, fuzzing RNG).

## Character Sets

Built-in character classes:

```go
Digits    = "0123456789"
Lowercase = "abcdefghijklmnopqrstuvwxyz"
Uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
Symbols   = "@#&-_+!$%^*()[]{}"
```

---

## Testing

Run all tests:

```bash
go test ./...
```

**Covered Components**:

- Config defaults and validation

- Charset builder

- Password generator logic

- Crypto RNG

- Mock RNG

- Error propagation

---

## Project Structure

```
passwordgen-lib/
 ├─ passwordgen/
 │   ├─ config.go
 │   ├─ charset.go
 │   ├─ generator.go
 │   ├─ config_test.go
 │   ├─ charset_test.go
 │   ├─ generator_test.go
 ├─ internal/rand/
 │   ├─ crypto.go
 │   ├─ mock.go
 │   ├─ crypto_test.go
 ├─ go.mod
 └─ README.md
```

## License

[![License](https://img.shields.io/github/license/buzzismaloy/passwordgen-lib)](https://github.com/buzzismaloy/passwordgen-lib/blob/main/LICENSE)

## About repo
![Go](https://img.shields.io/badge/Go-1.25-blue)
![Repo Size](https://img.shields.io/github/repo-size/buzzismaloy/passwordgen-lib)

## Author
[![GitHub](https://img.shields.io/badge/GitHub-buzzismaloy-black?logo=github)](https://github.com/buzzismaloy)
