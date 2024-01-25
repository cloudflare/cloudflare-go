// File generated from our OpenAPI spec by Stainless.

package cloudflare_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-sdk-go"
	"github.com/cloudflare/cloudflare-sdk-go/internal/testutil"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

func TestAccountDiagnosticTracerouteDiagnosticsTracerouteWithOptionalParams(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithAPIToken("my-cloudflare-api-token"),
		option.WithUserServiceKey("my-cloudflare-user-service-key"),
	)
	_, err := client.Accounts.Diagnostics.Traceroutes.DiagnosticsTraceroute(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.AccountDiagnosticTracerouteDiagnosticsTracerouteParams{
			Targets: cloudflare.F([]string{"203.0.113.1", "cloudflare.com"}),
			Colos:   cloudflare.F([]string{"den", "sin"}),
			Options: cloudflare.F(cloudflare.AccountDiagnosticTracerouteDiagnosticsTracerouteParamsOptions{
				MaxTtl:        cloudflare.F(int64(15)),
				PacketType:    cloudflare.F(cloudflare.AccountDiagnosticTracerouteDiagnosticsTracerouteParamsOptionsPacketTypeIcmp),
				PacketsPerTtl: cloudflare.F(int64(0)),
				Port:          cloudflare.F(int64(0)),
				WaitTime:      cloudflare.F(int64(1)),
			}),
		},
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
