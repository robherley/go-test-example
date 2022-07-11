package success_test

import (
	"fmt"
	"testing"
)

func TestSuccess(t *testing.T) {
	tests := []struct {
		a int
	}{
		{1},
		{2},
		{3},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("Subtest(%d)", tc.a), func(t *testing.T) {
			t.Logf("hello from subtest #%d", tc.a)
		})
	}
}
