// File generated from our OpenAPI spec by Stainless.

package cloudflare_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/cloudflare-go/internal/testutil"
	"github.com/cloudflare/cloudflare-go/option"
)

func TestSSLCertificatePackOrderNewWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.SSL.CertificatePacks.Order.New(context.TODO(), cloudflare.SSLCertificatePackOrderNewParams{
		ZoneID:               cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		CertificateAuthority: cloudflare.F(cloudflare.SSLCertificatePackOrderNewParamsCertificateAuthorityLetsEncrypt),
		Hosts:                cloudflare.F([]string{"example.com", "*.example.com", "www.example.com"}),
		Type:                 cloudflare.F(cloudflare.SSLCertificatePackOrderNewParamsTypeAdvanced),
		ValidationMethod:     cloudflare.F(cloudflare.SSLCertificatePackOrderNewParamsValidationMethodTXT),
		ValidityDays:         cloudflare.F(cloudflare.SSLCertificatePackOrderNewParamsValidityDays14),
		CloudflareBranding:   cloudflare.F(false),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
