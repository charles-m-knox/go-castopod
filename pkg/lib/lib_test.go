package castopod_test

import (
	"testing"

	castopod "github.com/charles-m-knox/go-castopod/pkg/lib"
)

func TestNewToken(t *testing.T) {
	t.Parallel()

	gotSecret, gotHash := castopod.NewToken()

	if len(gotSecret) != 8 {
		t.Logf("expected 8 digit token, gotSecret %v", len(gotSecret))
		t.Fail()
	}
	if len(gotHash) != 64 {
		t.Logf("expected 64 digit token, gotHash %v", len(gotHash))
		t.Fail()
	}
}
