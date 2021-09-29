# go-random

## A simple random data generator based on Golang standard libraries "math/rand" and "crypto/rand"

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
> go test -bench=. -benchmem
goos: windows
goarch: amd64
pkg: github.com/tee-ck/go-random
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
BenchmarkRandom_Bytes32-8                       17337657                65.19 ns/op           32 B/op          1 allocs/op
BenchmarkRandom_Int-8                           254578542                4.696 ns/op           0 B/op          0 allocs/op
BenchmarkRandom_Int31-8                         283269266                4.216 ns/op           0 B/op          0 allocs/op
BenchmarkRandom_Int63-8                         289433644                4.147 ns/op           0 B/op          0 allocs/op
BenchmarkRandom_Uint32-8                        286437813                4.161 ns/op           0 B/op          0 allocs/op
BenchmarkRandom_Uint64-8                        262137454                4.613 ns/op           0 B/op          0 allocs/op
BenchmarkRandom_Float32-8                       209240331                5.749 ns/op           0 B/op          0 allocs/op
BenchmarkRandom_FastFloat32-8                   292067305                4.140 ns/op           0 B/op          0 allocs/op
BenchmarkRandom_Float64-8                       238042197                5.078 ns/op           0 B/op          0 allocs/op
BenchmarkRandom_FastFloat64-8                   292005900                4.099 ns/op           0 B/op          0 allocs/op
BenchmarkRandom_Runes32-8                        2531392               473.0 ns/op           128 B/op          1 allocs/op
BenchmarkRandom_RandString32-8                   1651809               730.1 ns/op           176 B/op          2 allocs/op
BenchmarkRandom_RandString64-8                    857338              1397 ns/op             336 B/op          2 allocs/op
BenchmarkRandom_Choice-8                        10568961               109.9 ns/op            40 B/op          2 allocs/op
BenchmarkRandom_ChoiceStrings-8                 217578728                5.391 ns/op           0 B/op          0 allocs/op
BenchmarkCryptoRandom_Bytes32-8                  5317098               224.9 ns/op            32 B/op          1 allocs/op
BenchmarkCryptoRandom_Int15-8                    4251613               269.9 ns/op            48 B/op          3 allocs/op
BenchmarkCryptoRandom_Int31-8                    4259911               278.2 ns/op            48 B/op          3 allocs/op
BenchmarkCryptoRandom_FastInt31-8                7028402               171.3 ns/op             4 B/op          1 allocs/op
BenchmarkCryptoRandom_Int63-8                    4268498               276.9 ns/op            48 B/op          3 allocs/op
BenchmarkCryptoRandom_FastInt63-8                6684483               179.2 ns/op             8 B/op          1 allocs/op
BenchmarkCryptoRandom_Uint16-8                   4407358               273.4 ns/op            48 B/op          3 allocs/op
BenchmarkCryptoRandom_Uint32-8                   4281760               278.4 ns/op            48 B/op          3 allocs/op
BenchmarkCryptoRandom_FastUint32-8               7056573               168.4 ns/op             4 B/op          1 allocs/op
BenchmarkCryptoRandom_Uint64-8                   4198670               274.6 ns/op            48 B/op          3 allocs/op
BenchmarkCryptoRandom_FastUint64-8               6742682               177.6 ns/op             8 B/op          1 allocs/op
BenchmarkCryptoRandom_Float32-8                  4299981               277.5 ns/op            48 B/op          3 allocs/op
BenchmarkCryptoRandom_FastFloat32-8              6820102               170.4 ns/op             4 B/op          1 allocs/op
BenchmarkCryptoRandom_Float64-8                  4315041               275.3 ns/op            48 B/op          3 allocs/op
BenchmarkCryptoRandom_FastFloat64-8              6284011               179.7 ns/op             8 B/op          1 allocs/op
BenchmarkCryptoRandom_Runes32-8                   127357              9247 ns/op            1664 B/op         97 allocs/op
BenchmarkCryptoRandom_FastRunes32-8               193069              6062 ns/op             384 B/op         33 allocs/op
BenchmarkCryptoRandom_RandString32-8              123489              9478 ns/op            1712 B/op         98 allocs/op
BenchmarkCryptoRandom_FastRandString32-8          188001              6344 ns/op             432 B/op         34 allocs/op
BenchmarkCryptoRandom_RandString64-8               63314             19040 ns/op            3408 B/op        194 allocs/op
BenchmarkCryptoRandom_FastRandString64-8           94004             12806 ns/op             848 B/op         66 allocs/op
BenchmarkCryptoRandom_Choice-8                   3062632               400.0 ns/op            88 B/op          5 allocs/op
BenchmarkCryptoRandom_FastChoice-8               4006525               302.1 ns/op            48 B/op          3 allocs/op
BenchmarkCryptoRandom_ChoiceStrings-8            4178482               292.7 ns/op            48 B/op          3 allocs/op
BenchmarkCryptoRandom_FastChoiceStrings-8        6266858               187.5 ns/op             8 B/op          1 allocs/op
PASS
```

# License

The project is licensed under MIT license.

