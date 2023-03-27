package fib

import (
	"testing"
)

// func FibTestTry(t *testing.T) {
// 	t.Log("ftst")
// }

func TestFibSeqTry(t *testing.T) {
	t.Log("slslsl")
}

func TestTestFibEquTry(t *testing.T) {
	a := 1
	b := 1
	for i := 0; i < 5; i++ {
		t.Log("b:", b)
		tmp := a
		a = b
		b = tmp + a
	}
	t.Log(a, b)
}

func TestFibTryResult(t *testing.T) {
	a := 1
	b := 1

	for i := 0; i < 6; i++ {
		t.Log("b:", b)
		tmp := a
		a = b
		b = tmp + a
	}
	t.Log(a, b)
}
