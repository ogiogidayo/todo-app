package testutil

import (
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"io"
	"net/http"
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

func AssertResponse(t *testing.T, got *http.Response, status int, body []byte) {
	t.Helper()
	t.Cleanup(func() { _ = got.Body.Close() })
	gb, err := io.ReadAll(got.Body)
	if err != nil {
		t.Fatal(err)
	}
	if got.StatusCode != status {
		t.Fatalf("want status %d, but got %d, body %q", status, got.StatusCode, gb)
	}

	if len(gb) == 0 && len(body) == 0 {
		return
	}

	AssertJSON(t, body, gb)
}
