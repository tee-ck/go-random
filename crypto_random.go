package random

import (
	"crypto/rand"
	"encoding/binary"
	"math"
	"math/big"
	"reflect"
	"unicode/utf8"
)

const (
	IntBits = 32 << (^uint(0) >> 63) / 8
)

var (
	MaxBigInt    = big.NewInt(math.MaxInt)
	MaxBigInt16  = big.NewInt(math.MaxInt16)
	MaxBigInt32  = big.NewInt(math.MaxInt32)
	MaxBigInt64  = big.NewInt(math.MaxInt64)
	MaxBigUint16 = big.NewInt(math.MaxUint16)
	MaxBigUint32 = big.NewInt(math.MaxUint32)
	MaxBigUint64 = new(big.Int).SetUint64(math.MaxUint64)
)

type CryptoRandom struct {
	chars    []rune
	charsLen int64
}

func (c *CryptoRandom) Chars() Characters {
	return string(append(make([]rune, 0, c.charsLen), c.chars...))
}

func (c *CryptoRandom) RuneAtChars(n int) rune {
	return c.chars[n]
}

func (c *CryptoRandom) SetChars(chars Characters) {
	c.chars = []rune(chars)
	c.charsLen = int64(len(c.chars))
}

func (c *CryptoRandom) WithChars(chars Characters) *CryptoRandom {
	return &CryptoRandom{
		chars:    []rune(chars),
		charsLen: int64(utf8.RuneCountInString(chars)),
	}
}

func (c *CryptoRandom) BigInt(max *big.Int) (*big.Int, error) {
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		return nil, err
	}

	return n, nil
}

func (c *CryptoRandom) RandBigInt(min, max *big.Int) (*big.Int, error) {
	n, err := rand.Int(rand.Reader, new(big.Int).Sub(max, min))
	if err != nil {
		return nil, err
	}

	return new(big.Int).Add(min, n), nil
}

func (c *CryptoRandom) Bytes(size int) ([]byte, error) {
	buf := make([]byte, size)

	_, err := rand.Read(buf)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (c *CryptoRandom) Runes(size int) ([]rune, error) {
	buf := make([]rune, size)

	for i := 0; i < size; i++ {
		n, err := c.Int63()
		if err != nil {
			return nil, err
		}

		buf[i] = c.chars[n%c.charsLen]
	}

	return buf, nil
}

func (c *CryptoRandom) FastRunes(size int) ([]rune, error) {
	buf := make([]rune, size)

	for i := 0; i < size; i++ {
		n, err := c.FastInt63()
		if err != nil {
			return nil, err
		}

		buf[i] = c.chars[n%c.charsLen]
	}

	return buf, nil
}

func (c *CryptoRandom) RandString(size int) (string, error) {
	chars, err := c.Runes(size)
	if err != nil {
		return "", err
	}

	return string(chars), nil
}

func (c *CryptoRandom) FastRandString(size int) (string, error) {
	chars, err := c.FastRunes(size)
	if err != nil {
		return "", err
	}

	return string(chars), nil
}

func (c *CryptoRandom) Int() (int, error) {
	n, err := rand.Int(rand.Reader, MaxBigInt)
	if err != nil {
		return 0, err
	}

	return int(n.Int64()), nil
}

func (c *CryptoRandom) FastInt() (int, error) {
	buf := make([]byte, IntBits)
	_, err := rand.Read(buf)
	if err != nil {
		return 0, err
	}
	buf[0] = buf[0] % 0x80

	return int(binary.BigEndian.Uint64(buf)), nil
}

func (c *CryptoRandom) Int15() (int16, error) {
	n, err := rand.Int(rand.Reader, MaxBigInt16)
	if err != nil {
		return 0, err
	}

	return int16(n.Int64()), nil
}

func (c *CryptoRandom) Int31() (int32, error) {
	n, err := rand.Int(rand.Reader, MaxBigInt32)
	if err != nil {
		return 0, err
	}

	return int32(n.Int64()), nil
}

func (c *CryptoRandom) FastInt31() (int32, error) {
	buf := make([]byte, 4)
	_, err := rand.Read(buf)
	if err != nil {
		return 0, err
	}
	buf[0] = buf[0] % 0x80

	return int32(binary.BigEndian.Uint32(buf)), nil
}

func (c *CryptoRandom) Int63() (int64, error) {
	n, err := rand.Int(rand.Reader, MaxBigInt64)
	if err != nil {
		return 0, err
	}

	return n.Int64(), nil
}

func (c *CryptoRandom) FastInt63() (int64, error) {
	buf := make([]byte, 8)
	_, err := rand.Read(buf)
	if err != nil {
		return 0, err
	}
	buf[0] = buf[0] % 0x80

	return int64(binary.BigEndian.Uint64(buf)), nil
}

func (c *CryptoRandom) Uint16() (uint16, error) {
	n, err := rand.Int(rand.Reader, MaxBigUint16)
	if err != nil {
		return 0, err
	}

	return uint16(n.Uint64()), nil
}

func (c *CryptoRandom) FastUint16() (uint16, error) {
	buf := make([]byte, 2)

	_, err := rand.Read(buf)
	if err != nil {
		return 0, err
	}

	return binary.LittleEndian.Uint16(buf), nil
}

func (c *CryptoRandom) Uint32() (uint32, error) {
	n, err := rand.Int(rand.Reader, MaxBigUint32)
	if err != nil {
		return 0, err
	}

	return uint32(n.Uint64()), nil
}

func (c *CryptoRandom) FastUint32() (uint32, error) {
	buf := make([]byte, 4)

	_, err := rand.Read(buf)
	if err != nil {
		return 0, err
	}

	return binary.LittleEndian.Uint32(buf), nil
}

func (c *CryptoRandom) Uint64() (uint64, error) {
	n, err := rand.Int(rand.Reader, MaxBigUint64)
	if err != nil {
		return 0, err
	}

	return n.Uint64(), nil
}

func (c *CryptoRandom) FastUint64() (uint64, error) {
	buf := make([]byte, 8)

	_, err := rand.Read(buf)
	if err != nil {
		return 0, err
	}

	return binary.LittleEndian.Uint64(buf), nil
}

func (c *CryptoRandom) Float32() (float32, error) {
	n, err := c.Int31()
	if err != nil {
		return 0, err
	}

	return float32(n) / Float32ofMaxInt32, nil
}

func (c *CryptoRandom) FastFloat32() (float32, error) {
	n, err := c.FastInt31()
	if err != nil {
		return 0, err
	}

	return float32(n) / Float32ofMaxInt32, nil
}

func (c *CryptoRandom) Float64() (float64, error) {
	n, err := c.Int63()
	if err != nil {
		return 0, err
	}

	return float64(n) / Float64ofMaxInt64, nil
}

func (c *CryptoRandom) FastFloat64() (float64, error) {
	n, err := c.FastInt63()
	if err != nil {
		return 0, err
	}

	return float64(n) / Float64ofMaxInt64, nil
}

func (c *CryptoRandom) RandInt(min, max int) (int, error) {
	n, err := rand.Int(rand.Reader, MaxBigInt64)
	if err != nil {
		return 0, err
	}

	return min + (int(n.Int64()) % (max - min)), nil
}

func (c CryptoRandom) Choice(items interface{}) (interface{}, error) {
	kind := reflect.TypeOf(items).Kind()
	if kind == reflect.Slice || kind == reflect.Array {
		v := reflect.ValueOf(items)
		n, err := c.Int63()
		if err != nil {
			return nil, err
		}
		return v.Index(int(n) % v.Len()).Interface(), nil
	}

	return nil, nil
}

func (c CryptoRandom) FastChoice(items interface{}) (interface{}, error) {
	kind := reflect.TypeOf(items).Kind()
	if kind == reflect.Slice || kind == reflect.Array {
		v := reflect.ValueOf(items)
		n, err := c.FastInt63()
		if err != nil {
			return nil, err
		}
		return v.Index(int(n) % v.Len()).Interface(), nil
	}

	return nil, nil
}

func (c *CryptoRandom) ChoiceStrings(items []string) (string, error) {
	n, err := c.Int63()
	if err != nil {
		return "", nil
	}

	return items[n%int64(len(items))], nil
}

func (c *CryptoRandom) FastChoiceStrings(items []string) (string, error) {
	n, err := c.FastInt63()
	if err != nil {
		return "", nil
	}

	return items[n%int64(len(items))], nil
}
