// bigint.go

package Arithmetic

import (
	"errors"
	"math/big"
)

// BigInt represents a arbitrary-length integer.
type BigInt struct {
	value *big.Int
}

// NewBigInt creates a BigInt from an int.
func NewBigInt(n int64) *BigInt {
	return &BigInt{big.NewInt(n)}
}

// ParseBigInt parses a string into a BigInt.
func ParseBigInt(s string) (*BigInt, error) {
	val, success := new(big.Int).SetString(s, 10)
	if !success {
		return nil, errors.New("failed to parse string as BigInt")
	}
	return &BigInt{val}, nil
}

// String returns the string representation of the BigInt.
func (bi *BigInt) String() string {
	return bi.value.String()
}

// Int64 returns the int64 representation of the BigInt.
func (bi *BigInt) Int64() (int64, error) {
	if bi.value.IsInt64() {
		return bi.value.Int64(), nil
	}
	return 0, errors.New("cannot convert BigInt to int64")
}

// Add adds two BigInts.
func (bi *BigInt) Add(other *BigInt) *BigInt {
	result := &BigInt{new(big.Int).Add(bi.value, other.value)}
	return result
}

// Sub subtracts one BigInt from another.
func (bi *BigInt) Sub(other *BigInt) *BigInt {
	result := &BigInt{new(big.Int).Sub(bi.value, other.value)}
	return result
}

// Mul multiplies two BigInts.
func (bi *BigInt) Mul(other *BigInt) *BigInt {
	result := &BigInt{new(big.Int).Mul(bi.value, other.value)}
	return result
}

// Div divides one BigInt by another.
func (bi *BigInt) Div(other *BigInt) *BigInt {
	if other.value.Sign() == 0 {
		return nil // division by zero
	}
	result := &BigInt{new(big.Int).Div(bi.value, other.value)}
	return result
}

// Mod returns the modulus of one BigInt by another.
func (bi *BigInt) Mod(other *BigInt) *BigInt {
	if other.value.Sign() == 0 {
		return nil // division by zero
	}
	result := &BigInt{new(big.Int).Mod(bi.value, other.value)}
	return result
}

// Cmp compares two BigInts.
func (bi *BigInt) Cmp(other *BigInt) int {
	return bi.value.Cmp(other.value)
}

// Power computes the exponentiation of a BigInt.
func (bi *BigInt) Power(exp int64) *BigInt {
	result := &BigInt{new(big.Int).Exp(bi.value, big.NewInt(exp), nil)}
	return result
}

// KaratsubaMulti multiplies two BigInts using Karatsuba algorithm.
func KaratsubaMulti(x, y *BigInt) *BigInt {
	// Base case: if either of the numbers has only one digit, perform regular multiplication.
	if x.value.BitLen() <= 64 || y.value.BitLen() <= 64 {
		return x.Mul(y)
	}

	// Determine the length of the numbers and the half length.
	n := max(x.value.BitLen(), y.value.BitLen())
	m := (n + 1) / 2

	// Split x and y into two parts.
	xHigh, xLow := splitBigInt(x, m)
	yHigh, yLow := splitBigInt(y, m)

	// Recursive calls to compute intermediate values.
	z0 := KaratsubaMulti(xLow, yLow)
	z2 := KaratsubaMulti(xHigh, yHigh)
	z1 := KaratsubaMulti(xLow.Add(xHigh), yLow.Add(yHigh)).Sub(z0).Sub(z2)

	// Combine intermediate values to get the final result.
	result := z0.Add(z1.Power(2 * int64(m))).Add(z2.Power(2 * int64(m)))
	return result
}

// Helper function to split BigInt into two parts.
func splitBigInt(bi *BigInt, n int) (*BigInt, *BigInt) {
	mask := new(big.Int).Lsh(big.NewInt(1), uint(n)).Sub(big.NewInt(1), big.NewInt(1))
	high := new(big.Int).Rsh(bi.value, uint(n))
	low := new(big.Int).And(bi.value, mask)
	return &BigInt{high}, &BigInt{low}
}

// Helper function to find maximum of two numbers.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
