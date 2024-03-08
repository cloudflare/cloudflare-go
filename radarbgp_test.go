// File generated from our OpenAPI spec by Stainless.

package cloudflare_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/cloudflare-go/internal/testutil"
	"github.com/cloudflare/cloudflare-go/option"
)

func TestRadarBGPTimeseriesWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.BGP.Timeseries(context.TODO(), cloudflare.RadarBGPTimeseriesParams{
		AggInterval: cloudflare.F(cloudflare.RadarBGPTimeseriesParamsAggInterval1h),
		ASN:         cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   cloudflare.F([]cloudflare.RadarBGPTimeseriesParamsDateRange{cloudflare.RadarBGPTimeseriesParamsDateRange1d, cloudflare.RadarBGPTimeseriesParamsDateRange2d, cloudflare.RadarBGPTimeseriesParamsDateRange7d}),
		DateStart:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:      cloudflare.F(cloudflare.RadarBGPTimeseriesParamsFormatJson),
		Name:        cloudflare.F([]string{"string", "string", "string"}),
		Prefix:      cloudflare.F([]string{"string", "string", "string"}),
		UpdateType:  cloudflare.F([]cloudflare.RadarBGPTimeseriesParamsUpdateType{cloudflare.RadarBGPTimeseriesParamsUpdateTypeAnnouncement, cloudflare.RadarBGPTimeseriesParamsUpdateTypeWithdrawal}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
