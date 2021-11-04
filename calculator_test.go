package calculator

import "testing"

func TestMultiply(t *testing.T) {
	if got, want := Multiply(2, 2), 4; got != want {
		t.Errorf("add method produced wrong result. expected: %d, got: %d", want, got)
	}

	if got, want := Multiply(2, -2), -4; got != want {
		t.Errorf("add method produced wrong result. expected: %d, got: %d", want, got)
	}
}

