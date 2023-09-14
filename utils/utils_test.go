package utils

import "testing"

func Test_add(t *testing.T) {
	v := add(5, 3)
	if v != 8 {
		t.Errorf("expected: 8, got: %d", v)
	}
}

func Test_abs(t *testing.T) {
	v := abs(-50)
	if v != 50 {
		t.Errorf("expected: 50, got: %f", v)
	}

	v = abs(50)
	if v != 50 {
		t.Errorf("expected: 50, got: %f", v)
	}
}

func Test_mul(t *testing.T) {
	v := mul(5, 3)
	if v != 15 {
		t.Errorf("expected: 15, got: %d", v)
	}
}
