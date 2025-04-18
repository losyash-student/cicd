package calc

import (
	"testing"
)

func TestSumm(t *testing.T) {
	ans := Summ(2, -2)
	if ans != 0 {
		t.Errorf("summ(2, -2) = %d; want 0", ans)
	}

}

func TestMinus(t *testing.T) {
	ans := Minus(2, -2)
	if ans != 4 {
		t.Errorf("minus(2, -2) = %d; want 0-4", ans)
	}

}
