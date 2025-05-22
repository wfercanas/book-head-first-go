package arithmetic

import "testing"

func TestAdd(t *testing.T) {
	if Add(1, 2) != 3 {
		t.Error("1 + 2 must be equal to 3")
	}
}

func TestSubstract(t *testing.T) {
	if Substract(8, 4) != 4 {
		t.Error("8 - 4 must be equal to 4")
	}
}
