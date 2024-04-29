// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/radar"
)

func TestTrafficAnomalyGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.TrafficAnomalies.Get(context.TODO(), radar.TrafficAnomalyGetParams{
		ASN:       cloudflare.F(int64(174)),
		DateEnd:   cloudflare.F(time.Now()),
		DateRange: cloudflare.F(radar.TrafficAnomalyGetParamsDateRange7d),
		DateStart: cloudflare.F(time.Now()),
		Format:    cloudflare.F(radar.TrafficAnomalyGetParamsFormatJson),
		Limit:     cloudflare.F(int64(5)),
		Location:  cloudflare.F("US"),
		Offset:    cloudflare.F(int64(0)),
		Status:    cloudflare.F(radar.TrafficAnomalyGetParamsStatusVerified),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
