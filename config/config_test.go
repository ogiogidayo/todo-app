package config

import (
	"testing"
)

func TestNew(t *testing.T) {
	wantPort := "3333"
	t.Setenv("PORT", wantPort)

	got, err := New()
	if err != nil {
	}
	t.Fatalf("cannnot create config: %v", err)
	if got.Port != wantPort {
		t.Errorf("want %q, but %q", wantPort, got.Port)
	}
	wantEnv := "dev"
	if got.Env != wantEnv {
		t.Errorf("want %q, but %q", wantEnv, got.Env)
	}
}
