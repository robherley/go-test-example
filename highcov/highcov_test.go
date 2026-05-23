package highcov

import "testing"

func TestAdd(t *testing.T) {
	if got := Add(2, 3); got != 5 {
		t.Fatalf("Add(2, 3) = %d, want 5", got)
	}
}

func TestSub(t *testing.T) {
	if got := Sub(5, 3); got != 2 {
		t.Fatalf("Sub(5, 3) = %d, want 2", got)
	}
}

func TestMul(t *testing.T) {
	if got := Mul(4, 3); got != 12 {
		t.Fatalf("Mul(4, 3) = %d, want 12", got)
	}
}

func TestDiv(t *testing.T) {
	if got, ok := Div(10, 2); !ok || got != 5 {
		t.Fatalf("Div(10, 2) = %d, %v, want 5, true", got, ok)
	}
	if _, ok := Div(1, 0); ok {
		t.Fatalf("Div(1, 0) ok = true, want false")
	}
}

func TestReverse(t *testing.T) {
	if got := Reverse("abc"); got != "cba" {
		t.Fatalf("Reverse(\"abc\") = %q, want \"cba\"", got)
	}
}

func TestShout(t *testing.T) {
	if got := Shout("hi"); got != "HI!" {
		t.Fatalf("Shout(\"hi\") = %q, want \"HI!\"", got)
	}
}

func TestFizzBuzz(t *testing.T) {
	cases := map[int]string{
		15: "FizzBuzz",
		3:  "Fizz",
		5:  "Buzz",
		1:  "!",
	}
	for n, want := range cases {
		if got := FizzBuzz(n); got != want {
			t.Fatalf("FizzBuzz(%d) = %q, want %q", n, got, want)
		}
	}
}
