package Arithmetic

import "testing"

func TestBigInt(t *testing.T) {
	// Test cases for basic operations
	bi1 := NewBigInt(456)
	bi2 := NewBigInt(123)

	if bi1.String() != "456" {
		t.Errorf("Expected 456, got %s", bi1.String())
	}
	if bi2.String() != "123" {
		t.Errorf("Expected 123, got %s", bi2.String())
	}

	// Test addition
	sum := bi1.Add(bi2)
	if sum.String() != "579" {
		t.Errorf("Expected 579, got %s", sum.String())
	}

	// Test subtraction
	diff := bi1.Sub(bi2)
	if diff.String() != "333" {
		t.Errorf("Expected 333, got %s", diff.String())
	}

	// Test multiplication
	prod := bi1.Mul(bi2)
	if prod.String() != "56088" {
		t.Errorf("Expected 56088, got %s", prod.String())
	}

	// Test division
	quotient := bi1.Div(bi2)
	if quotient.String() != "3" {
		t.Errorf("Expected 3, got %s", quotient.String())
	}

	// Test modulus
	mod := bi1.Mod(bi2)
	if mod.String() != "87" {
		t.Errorf("Expected 87, got %s", mod.String())
	}

	// Test exponentiation
	exp := bi1.Power(3)
	if exp.String() != "94818816" {
		t.Errorf("Expected 94818816, got %s", exp.String())
	}

	// Test comparison
	bi3 := NewBigInt(123)
	if bi1.Cmp(bi3) != 1 {
		t.Errorf("Expected bi1 to be greater than bi3")
	}
	bi4 := NewBigInt(456)
	if bi1.Cmp(bi4) != 0 {
		t.Errorf("Expected bi1 and bi4 to be equal")
	}
	if bi2.Cmp(bi1) != -1 {
		t.Errorf("Expected bi2 to be less than bi1")
	}
}

func TestKaratsubaMulti(t *testing.T) {
	// Test cases for Karatsuba multiplication
	bi1 := NewBigInt(100000)
	bi2 := NewBigInt(50000)

	// Expected result from regular multiplication
	expected := bi1.Mul(bi2)

	// Test Karatsuba multiplication
	result := KaratsubaMulti(bi1, bi2)

	if result.String() != expected.String() {
		t.Errorf("Expected %s, got %s", expected.String(), result.String())
	}
}
