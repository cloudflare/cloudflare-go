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

func TestRadarTrafficAnomalyGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.TrafficAnomalies.Get(context.TODO(), cloudflare.RadarTrafficAnomalyGetParams{
		ASN:       cloudflare.F(int64(0)),
		DateEnd:   cloudflare.F(time.Now()),
		DateRange: cloudflare.F(cloudflare.RadarTrafficAnomalyGetParamsDateRange7d),
		DateStart: cloudflare.F(time.Now()),
		Format:    cloudflare.F(cloudflare.RadarTrafficAnomalyGetParamsFormatJson),
		Limit:     cloudflare.F(int64(5)),
		Location:  cloudflare.F("US"),
		Offset:    cloudflare.F(int64(0)),
		Status:    cloudflare.F(cloudflare.RadarTrafficAnomalyGetParamsStatusVerified),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
