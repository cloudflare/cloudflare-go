// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/radar"
)

func TestBGPTimeseriesWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.BGP.Timeseries(context.TODO(), radar.BGPTimeseriesParams{
		AggInterval: cloudflare.F(radar.BGPTimeseriesParamsAggInterval15m),
		ASN:         cloudflare.F([]string{"string"}),
		DateEnd:     cloudflare.F([]time.Time{time.Now()}),
		DateRange:   cloudflare.F([]string{"7d"}),
		DateStart:   cloudflare.F([]time.Time{time.Now()}),
		Format:      cloudflare.F(radar.BGPTimeseriesParamsFormatJson),
		Name:        cloudflare.F([]string{"string"}),
		Prefix:      cloudflare.F([]string{"1.1.1.0/24"}),
		UpdateType:  cloudflare.F([]radar.BGPTimeseriesParamsUpdateType{radar.BGPTimeseriesParamsUpdateTypeAnnouncement}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
