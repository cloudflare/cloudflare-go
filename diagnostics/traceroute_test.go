// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package diagnostics_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/diagnostics"
	"github.com/cloudflare/cloudflare-go/v4/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v4/option"
)

func TestTracerouteNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Diagnostics.Traceroutes.New(context.TODO(), diagnostics.TracerouteNewParams{
		AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Targets:   cloudflare.F([]string{"203.0.113.1", "cloudflare.com"}),
		Colos:     cloudflare.F([]string{"den", "sin"}),
		Options: cloudflare.F(diagnostics.TracerouteNewParamsOptions{
			MaxTTL:        cloudflare.F(int64(15)),
			PacketType:    cloudflare.F(diagnostics.TracerouteNewParamsOptionsPacketTypeIcmp),
			PacketsPerTTL: cloudflare.F(int64(0)),
			Port:          cloudflare.F(int64(0)),
			WaitTime:      cloudflare.F(int64(1)),
		}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
