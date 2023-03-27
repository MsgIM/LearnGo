package constant_test

import "testing"

const (
	Monday = 1 + iota
	Tuesday
	Wednesday
)

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestConstant(t *testing.T) {
	t.Log(Monday, Tuesday)
}

func TestConstantTry(t *testing.T) {
	a := 1
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}

func TestBool(t *testing.T) {
	a := false
	t.Log("this is a bool value", a)
}
