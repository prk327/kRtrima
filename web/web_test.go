package web

import (
	"testing"
)

func Testversion(t *testing.T) {
	want := "0.0.1"
	if got := version(); got != want {
		t.Errorf("Web() = %q, want %q", got, want)
	}
}
