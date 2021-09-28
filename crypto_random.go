package random

import (
	"crypto/rand"
	"math/big"
	"unicode/utf8"
)

var (
	MaxInt16  = big.NewInt(32767)
	MaxInt32  = big.NewInt(2147483647)
	MaxInt64  = big.NewInt(9223372036854775807)
	MaxUint16 = big.NewInt(65535)
	MaxUint32 = big.NewInt(4294967295)
	MaxUint64 = MaxInt16.Add(MaxInt16.Add(MaxInt64, MaxInt64), big.NewInt(1))
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

func (c *CryptoRandom) RandString(size int) (string, error) {
	chars, err := c.Runes(size)
	if err != nil {
		return "", err
	}

	return string(chars), nil
}

func (c *CryptoRandom) Int63() (int64, error) {
	n, err := rand.Int(rand.Reader, MaxInt64)
	if err != nil {
		return 0, err
	}

	return n.Int64(), nil
}

func (c *CryptoRandom) Int31() (int32, error) {
	n, err := rand.Int(rand.Reader, MaxInt32)
	if err != nil {
		return 0, err
	}

	return int32(n.Int64()), nil
}

func (c *CryptoRandom) Int15() (int16, error) {
	n, err := rand.Int(rand.Reader, MaxInt16)
	if err != nil {
		return 0, err
	}

	return int16(n.Int64()), nil
}

func (c *CryptoRandom) Uint16() (uint16, error) {
	n, err := rand.Int(rand.Reader, MaxUint16)
	if err != nil {
		return 0, err
	}

	return uint16(n.Uint64()), nil
}

func (c *CryptoRandom) Uint32() (uint32, error) {
	n, err := rand.Int(rand.Reader, MaxUint32)
	if err != nil {
		return 0, err
	}

	return uint32(n.Uint64()), nil
}

func (c *CryptoRandom) Uint64() (uint64, error) {
	n, err := rand.Int(rand.Reader, MaxUint64)
	if err != nil {
		return 0, err
	}

	return n.Uint64(), nil
}

func (c *CryptoRandom) RandInt(min, max int) (int, error) {
	n, err := rand.Int(rand.Reader, MaxInt64)
	if err != nil {
		return 0, err
	}

	return min + (int(n.Int64()) % (max - min)), nil
}
