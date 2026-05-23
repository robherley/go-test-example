package lowcov

import "testing"

// Only Add is tested here, leaving the rest of the package uncovered
// so this module reports low coverage.
func TestAdd(t *testing.T) {
	if got := Add(2, 3); got != 5 {
		t.Fatalf("Add(2, 3) = %d, want 5", got)
	}
}
