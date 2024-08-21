package castopod_test

import (
	"testing"

	castopod "git.cmcode.dev/cmcode/go-castopod/pkg/lib"
)

func TestNewToken(t *testing.T) {
	t.Parallel()

	got := castopod.NewToken()
	if len(got) != 64 {
		t.Logf("expected 64 digit token, got %v", len(got))
		t.Fail()
	}
}
