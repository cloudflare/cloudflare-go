// File generated from our OpenAPI spec by Stainless.

package cloudflare_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go"
	"github.com/cloudflare/cloudflare-sdk-go/internal/testutil"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

func TestIntelDNSGetWithOptionalParams(t *testing.T) {
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
	)
	_, err := client.Intel.DNS.Get(context.TODO(), cloudflare.IntelDNSGetParams{
		AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		IPV4:      cloudflare.F("string"),
		Page:      cloudflare.F(1.000000),
		PerPage:   cloudflare.F(20.000000),
		StartEndParams: cloudflare.F(cloudflare.IntelDNSGetParamsStartEndParams{
			End:   cloudflare.F(time.Now()),
			Start: cloudflare.F(time.Now()),
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
