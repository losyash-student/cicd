package main

import (
	"testing"
)

func TestSumm(t *testing.T) {
	ans := summ(2, -2)
	if ans != 0 {
		t.Errorf("summ(2, -2) = %d; want 0", ans)
	}

}

func TestMinus(t *testing.T) {
	ans := minus(2, -2)
	if ans != 4 {
		t.Errorf("minus(2, -2) = %d; want 0-4", ans)
	}

}
