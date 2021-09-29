package random

import (
	"testing"
)

func BenchmarkRandom_Bytes32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bytes(32)

	}
}

func BenchmarkRandom_Int(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Int()
	}
}

func BenchmarkRandom_Int31(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Int31()
	}
}

func BenchmarkRandom_Int63(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Int63()
	}
}

func BenchmarkRandom_Uint32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Uint32()
	}
}

func BenchmarkRandom_Uint64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Uint64()
	}
}

func BenchmarkRandom_Float32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Float32()
	}
}

func BenchmarkRandom_FastFloat32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FastFloat32()
	}
}

func BenchmarkRandom_Float64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Float64()
	}
}

func BenchmarkRandom_FastFloat64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FastFloat64()
	}
}

func BenchmarkRandom_Runes32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Runes(32)
	}
}

func BenchmarkRandom_RandString32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandString(32)
	}
}

func BenchmarkRandom_RandString64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandString(64)
	}
}

func BenchmarkRandom_Choice(b *testing.B) {
	items := []string{
		"apple",
		"banana",
		"mango",
		"dragon fruit",
		"lemon",
		"orange",
		"peach",
	}

	for i := 0; i < b.N; i++ {
		Choice(items)
	}
}

func BenchmarkRandom_ChoiceStrings(b *testing.B) {
	items := []string{
		"apple",
		"banana",
		"mango",
		"dragon fruit",
		"lemon",
		"orange",
		"peach",
	}

	for i := 0; i < b.N; i++ {
		ChoiceStrings(items)
	}
}

func BenchmarkCryptoRandom_Bytes32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := crypto.Bytes(32); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkCryptoRandom_Int(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := crypto.Int(); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkCryptoRandom_FastInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := crypto.FastInt(); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkCryptoRandom_Int15(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := crypto.Int15(); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkCryptoRandom_Int31(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := crypto.Int31(); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkCryptoRandom_FastInt31(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := crypto.FastInt31(); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkCryptoRandom_Int63(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := crypto.Int63(); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkCryptoRandom_FastInt63(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := crypto.FastInt63(); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkCryptoRandom_Uint16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := crypto.Uint16(); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkCryptoRandom_Uint32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := crypto.Uint32(); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkCryptoRandom_FastUint32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := crypto.FastUint32(); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkCryptoRandom_Uint64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := crypto.Uint64(); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkCryptoRandom_FastUint64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := crypto.FastUint64(); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkCryptoRandom_Float32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := crypto.Float32(); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkCryptoRandom_FastFloat32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := crypto.FastFloat32(); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkCryptoRandom_Float64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := crypto.Float64(); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkCryptoRandom_FastFloat64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := crypto.FastFloat64(); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkCryptoRandom_Runes32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := crypto.Runes(32); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkCryptoRandom_FastRunes32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := crypto.FastRunes(32); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkCryptoRandom_RandString32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := crypto.RandString(32); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkCryptoRandom_FastRandString32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := crypto.FastRandString(32); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkCryptoRandom_RandString64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := crypto.RandString(64); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkCryptoRandom_FastRandString64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := crypto.FastRandString(64); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkCryptoRandom_Choice(b *testing.B) {
	items := []string{
		"apple",
		"banana",
		"mango",
		"dragon fruit",
		"lemon",
		"orange",
		"peach",
	}

	for i := 0; i < b.N; i++ {
		if _, err := crypto.Choice(items); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkCryptoRandom_FastChoice(b *testing.B) {
	items := []string{
		"apple",
		"banana",
		"mango",
		"dragon fruit",
		"lemon",
		"orange",
		"peach",
	}

	for i := 0; i < b.N; i++ {
		if _, err := crypto.FastChoice(items); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkCryptoRandom_ChoiceStrings(b *testing.B) {
	items := []string{
		"apple",
		"banana",
		"mango",
		"dragon fruit",
		"lemon",
		"orange",
		"peach",
	}

	for i := 0; i < b.N; i++ {
		if _, err := crypto.ChoiceStrings(items); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkCryptoRandom_FastChoiceStrings(b *testing.B) {
	items := []string{
		"apple",
		"banana",
		"mango",
		"dragon fruit",
		"lemon",
		"orange",
		"peach",
	}

	for i := 0; i < b.N; i++ {
		if _, err := crypto.FastChoiceStrings(items); err != nil {
			b.Fatal(err)
		}
	}
}
