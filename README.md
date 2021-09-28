# go-random

## A simple random data generator based on golang standard libraries "math/rand" and "crypto/rand"

# Installation

```shell
go get -u github.com/tee-ck/go-random
```

# Usage

## Based on math/rand

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

## Based on crypto/rand

### crypto functions usually require error handling

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

# License

The project is licensed under MIT license.

