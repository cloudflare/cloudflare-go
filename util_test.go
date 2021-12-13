package cloudflare

import (
	"os"
	"testing"
)

func TestGetAPIURLDefault(t *testing.T) {
	if GetAPIUrl() != "https://api.cloudflare.com/client/v4" {
		t.Error("Expected https://api.cloudflare.com/client/v4, got", GetAPIUrl())
	}
}

func TestGetAPIURLEnv(t *testing.T) {
	os.Setenv("CLOUDFLARE_API_HOSTNAME", "other.host.cloudflare.com")
	defer os.Unsetenv("CLOUDFLARE_API_HOSTNAME")
	if GetAPIUrl() != "https://other.host.cloudflare.com/client/v4" {
		t.Error("Expected https://other.host.cloudflare.com/client/v4, got", GetAPIUrl())
	}
}
