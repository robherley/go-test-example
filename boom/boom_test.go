package boom_test

import "testing"

func TestFatal(t *testing.T) {
	t.Fatal("this was a failure")
}

func TestPanic(t *testing.T) {
	panic("this was a panic")
}
