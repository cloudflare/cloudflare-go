// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package security_txt_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/security_txt"
)

func TestSecurityTXTUpdateWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.SecurityTXT.Update(context.TODO(), security_txt.SecurityTXTUpdateParams{
		ZoneID:             cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Acknowledgments:    cloudflare.F([]string{"https://example.com/hall-of-fame.html"}),
		Canonical:          cloudflare.F([]string{"https://www.example.com/.well-known/security.txt"}),
		Contact:            cloudflare.F([]string{"mailto:security@example.com", "tel:+1-201-555-0123", "https://example.com/security-contact.html"}),
		Enabled:            cloudflare.F(true),
		Encryption:         cloudflare.F([]string{"https://example.com/pgp-key.txt", "dns:5d2d37ab76d47d36._openpgpkey.example.com?type=OPENPGPKEY", "openpgp4fpr:5f2de5521c63a801ab59ccb603d49de44b29100f"}),
		Expires:            cloudflare.F(time.Now()),
		Hiring:             cloudflare.F([]string{"https://example.com/jobs.html"}),
		Policy:             cloudflare.F([]string{"https://example.com/disclosure-policy.html"}),
		PreferredLanguages: cloudflare.F("en, es, fr"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSecurityTXTDelete(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.SecurityTXT.Delete(context.TODO(), security_txt.SecurityTXTDeleteParams{
		ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSecurityTXTGet(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.SecurityTXT.Get(context.TODO(), security_txt.SecurityTXTGetParams{
		ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
