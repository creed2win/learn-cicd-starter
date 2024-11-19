package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	header := http.Header{}
	header.Set("Authorization", "ApiKey hasdifhaijhgdfkjahdfojahwepofhakdjlfhalkjhfdkljsaf")
	want := "hasdifhaijhgdfkjahdfojahwepofhakdjlfhalkjhfdkljsaf"
	getting, err := GetAPIKey(header)

	if want == getting && err == nil {
		return
	}
	t.Fatal("API key does not match input header API Key", err, header)
}
