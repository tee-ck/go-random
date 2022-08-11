# go-random
A simple random data generator based on Golang standard libraries "math/rand" and "crypto/rand"

# requirements
go version > 1.18

# Installation

```shell
go get -u github.com/tee-ck/go-random
```

# Usage

Based on math/rand

```go
package main

import (
	"fmt"
	"github.com/tee-ck/go-random"
	"time"
)

func main() {
	// generate a random string with default characters and length given (32 characters)
	fmt.Println(random.RandString(32))
	// output: 9FgvfIBSlIQUvRqqB5WcGwFbofO8OjKO

	// generate a random string with characters given and length given (24 characters) by function chain
	fmt.Println(random.WithChars("0123456789").RandString(24))
	// output: 413864904880075041479896

	// generate a random string with seed given and length given (24 characters) by function chain
	fmt.Println(random.WithSeed(time.Now().UnixNano()).RandString(24))
	// output: D3nQ5hwH9HIW03neQERsEgIC

	// generate a random int between 10 and 100
	fmt.Println(random.RandInt(10, 100))
	// output: 74

	// generate a byte slice ([]byte) with length given (12)
	fmt.Println(random.Bytes(12))
	// output: [1 207 75 27 249 186 141 7 72 44 85 88]

	// generate a rune slice ([]rune) with default characters and length given (12)
	fmt.Println(random.Runes(12))
	// output: [89 116 89 98 105 90 69 70 110 65 119 48]

	// generate a random 64 bit integer
	fmt.Println(random.Int63())
	// output: 5928505690140062246

	// generate a random 32 bit integer
	fmt.Println(random.Int31())
	// output: 588574091

	// generate a random 32 bit floating point
	fmt.Println(random.Float32())
	// output: 0.16250405

	// generate a random 64 bit floating point
	fmt.Println(random.Float64())
	// output: 0.10604198575158426
}
```

Based on crypto/rand

- crypto functions usually require error handling

```go
package main

import (
	"fmt"
	"github.com/tee-ck/go-random"
	"log"
)

func main() {
	// first, create a crypto randomizer
	c := random.Crypto()

	// generate a random string with default characters and length given (32 characters)
	v, err := c.RandString(32)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(v)
	// output: 7kRBEJcGl0tBzYYCojM1j7N9Y3mrlPOv

	// generate a random string with characters given and length given (24 characters) by function chain
	v, err = c.WithChars("0123456789abcdef").RandString(24)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(v)
	// output: 890ce6af8a62cdd2b8566b77

	// generate a random string with characters given and length given (24 characters) by function chain
	v, err = c.WithChars(random.CharsRomanLettersUpper).RandString(24)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(v)
	// output: MQESAKHTHRABBOAXPJBYVTOO

	// or you can use crypto randomizer by function chain
	v, err = random.Crypto().RandString(24)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(v)
	// output: 95d4457abcebb294e9388674

	/// ...
	// same with above math/rand, just need to do the extra error handling
}
```

# Benchmarks

```shell
PS C:\Programming\.projects\.public\go-random> go version
go version go1.19 windows/amd64
PS C:\Programming\.projects\.public\go-random> go test -bench . -benchmem
goos: windows
goarch: amd64
pkg: github.com/tee-ck/go-random
cpu: 12th Gen Intel(R) Core(TM) i5-12600K
BenchmarkRandom_Bytes32-16                      39999866                28.81 ns/op           32 B/op          1 allocs/op
BenchmarkRandom_Int-16                          551555500                2.147 ns/op           0 B/op          0 allocs/op
BenchmarkRandom_Int31-16                        668201673                1.774 ns/op           0 B/op          0 allocs/op
BenchmarkRandom_Int63-16                        664688479                1.783 ns/op           0 B/op          0 allocs/op
BenchmarkRandom_Uint32-16                       662852685                1.768 ns/op           0 B/op          0 allocs/op
BenchmarkRandom_Uint64-16                       507398289                2.332 ns/op           0 B/op          0 allocs/op
BenchmarkRandom_Float32-16                      330874932                3.513 ns/op           0 B/op          0 allocs/op
BenchmarkRandom_FastFloat32-16                  664313151                1.772 ns/op           0 B/op          0 allocs/op
BenchmarkRandom_Float64-16                      481848950                2.508 ns/op           0 B/op          0 allocs/op
BenchmarkRandom_Runes32-16                       8571152               138.1 ns/op           128 B/op          1 allocs/op
BenchmarkRandom_RandString32-16                  4600060               256.8 ns/op           176 B/op          2 allocs/op
BenchmarkRandom_RandString64-16                  2546078               469.9 ns/op           336 B/op          2 allocs/op
BenchmarkRandom_Choice-16                       396812930                2.958 ns/op           0 B/op          0 allocs/op
BenchmarkRandom_CryptoChoice-16                 10788126               111.3 ns/op            32 B/op          2 allocs/op
BenchmarkCryptoRandom_Bytes32-16                 8807842               135.2 ns/op            56 B/op          2 allocs/op
BenchmarkCryptoRandom_Int-16                     7339480               163.3 ns/op            72 B/op          4 allocs/op
BenchmarkCryptoRandom_FastInt-16                10792628               111.2 ns/op            32 B/op          2 allocs/op
BenchmarkCryptoRandom_Int15-16                   7547150               157.8 ns/op            72 B/op          4 allocs/op
BenchmarkCryptoRandom_Int31-16                   7287877               162.3 ns/op            72 B/op          4 allocs/op
BenchmarkCryptoRandom_FastInt31-16              11311236               106.5 ns/op            28 B/op          2 allocs/op
BenchmarkCryptoRandom_Int63-16                   7450393               162.6 ns/op            72 B/op          4 allocs/op
BenchmarkCryptoRandom_FastInt63-16              10666524               110.7 ns/op            32 B/op          2 allocs/op
BenchmarkCryptoRandom_Uint16-16                  7669504               157.9 ns/op            72 B/op          4 allocs/op
BenchmarkCryptoRandom_Uint32-16                  7414854               161.4 ns/op            72 B/op          4 allocs/op
BenchmarkCryptoRandom_FastUint32-16             11327775               104.8 ns/op            28 B/op          2 allocs/op
BenchmarkCryptoRandom_Uint64-16                  7500018               160.8 ns/op            72 B/op          4 allocs/op
BenchmarkCryptoRandom_FastUint64-16             11316878               109.7 ns/op            32 B/op          2 allocs/op
BenchmarkCryptoRandom_Float32-16                 7318242               160.8 ns/op            72 B/op          4 allocs/op
BenchmarkCryptoRandom_FastFloat32-16            11555961               105.5 ns/op            28 B/op          2 allocs/op
BenchmarkCryptoRandom_Float64-16                 7375858               161.3 ns/op            72 B/op          4 allocs/op
BenchmarkCryptoRandom_FastFloat64-16            10811073               110.8 ns/op            32 B/op          2 allocs/op
BenchmarkCryptoRandom_Runes32-16                  222214              5305 ns/op            2432 B/op        129 allocs/op
BenchmarkCryptoRandom_FastRunes32-16              324321              3616 ns/op            1152 B/op         65 allocs/op
BenchmarkCryptoRandom_RandString32-16             216213              5467 ns/op            2480 B/op        130 allocs/op
BenchmarkCryptoRandom_FastRandString32-16         319977              3743 ns/op            1200 B/op         66 allocs/op
BenchmarkCryptoRandom_RandString64-16             110599             11077 ns/op            4944 B/op        258 allocs/op
BenchmarkCryptoRandom_FastRandString64-16         161073              7473 ns/op            2384 B/op        130 allocs/op
BenchmarkCryptoRandom_Choice-16                  5638254               212.1 ns/op           112 B/op          6 allocs/op
BenchmarkCryptoRandom_FastChoice-16              7537574               160.8 ns/op            72 B/op          4 allocs/op
BenchmarkCryptoRandom_ChoiceStrings-16           7323315               163.4 ns/op            72 B/op          4 allocs/op
BenchmarkCryptoRandom_FastChoiceStrings-16      10793443               111.6 ns/op            32 B/op          2 allocs/op
PASS
ok      github.com/tee-ck/go-random     55.695s
```

# License

The project is licensed under MIT license.

