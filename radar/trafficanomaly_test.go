// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go/v5"
	"github.com/cloudflare/cloudflare-go/v5/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/radar"
)

func TestTrafficAnomalyGetWithOptionalParams(t *testing.T) {
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
		DateRange: cloudflare.F("7d"),
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
