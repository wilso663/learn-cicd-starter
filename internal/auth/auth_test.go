package auth

import (
	"testing"
	"net/http"
	"errors"
)

func TestGetAPIKey(t *testing.T) {
	t.Run("Test No Header Error", func(t *testing.T) {
		headers := http.Header{}

		_, err := GetAPIKey(headers)
		if err == nil {
			t.Fatalf("expected error, got nil")
		}
		if !errors.Is(err, ErrNoAuthHeaderIncluded) {
			t.Fatalf("expected ErrNoAuthHeaderIncluded, got %v", err)
		}
	})

	t.Run("Test Malformed Header Error", func(t *testing.T) {
		headers := http.Header{}
		headers.Set("Authorization", "Bearer sometoken")

		_, err := GetAPIKey(headers)

		if err == nil {
			t.Fatalf("expected error for malformed header, got nil")
		}

		if err.Error() != "malformed authorization header" {
			t.Fatalf("expected malformed header error, got %v", err)
		}
	})

	t.Run("Test Valid Header", func(t *testing.T) {
		headers := http.Header{}

		headers.Set("Authorization", "ApiKey abc123")
		apiKey, err := GetAPIKey(headers)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if apiKey != "abc123" {
			t.Fatalf("expected api key 'abc123', got %s", apiKey)
		}
	})


}