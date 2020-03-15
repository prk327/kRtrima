package mongoDB

import (
    "testing"
)

func TestRun_mongoDB(t *testing.T) {
    want := "Connected to MongoDB!"
    if got := Run_mongoDB(); got != want {
        t.Errorf("Run_mongoDB() = %q, want %q", got, want)
    }
}
