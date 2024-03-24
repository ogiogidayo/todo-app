package testutil

import (
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func AssertJSON(t *testing.T, want, got []byte) {
	t.Helper()

	var jw, jg any
	if err := json.Unmarshal(want, &jw); err != nil {
		t.Fatalf("cannot unmarshal want %q: %v", want, err)
	}
	if err := json.Unmarshal(got, &jg); err != nil {
		t.Fatalf("cannot unmarshal want %q: %v", got, err)
	}
	if diff := cmp.Diff(jw, jg); diff != "" {
		t.Fatalf("got differs: (-got +want)\n%s", diff)
	}
}
