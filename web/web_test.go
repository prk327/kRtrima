package web

import (
	"testing"
)

func TestVersion(t *testing.T) {
	want := "0.0.1"
	if got := version(); got != want {
		t.Errorf("Web() = %q, want %q", got, want)
	}
}
