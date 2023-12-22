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

func TestRadarEmailTimeserySpfListWithOptionalParams(t *testing.T) {
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
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithEmail("dev@cloudflare.com"),
	)
	_, err := client.Radars.Emails.Timeseries.Spfs.List(context.TODO(), cloudflare.RadarEmailTimeserySpfListParams{
		AggInterval: cloudflare.F(cloudflare.RadarEmailTimeserySpfListParamsAggInterval15m),
		Arc:         cloudflare.F([]cloudflare.RadarEmailTimeserySpfListParamsArc{cloudflare.RadarEmailTimeserySpfListParamsArcPass, cloudflare.RadarEmailTimeserySpfListParamsArcNone, cloudflare.RadarEmailTimeserySpfListParamsArcFail}),
		ASN:         cloudflare.F([]string{"15169", "15169", "15169"}),
		DateEnd:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   cloudflare.F([]cloudflare.RadarEmailTimeserySpfListParamsDateRange{cloudflare.RadarEmailTimeserySpfListParamsDateRange1d, cloudflare.RadarEmailTimeserySpfListParamsDateRange7d, cloudflare.RadarEmailTimeserySpfListParamsDateRange14d}),
		DateStart:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Dkim:        cloudflare.F([]cloudflare.RadarEmailTimeserySpfListParamsDkim{cloudflare.RadarEmailTimeserySpfListParamsDkimPass, cloudflare.RadarEmailTimeserySpfListParamsDkimNone, cloudflare.RadarEmailTimeserySpfListParamsDkimFail}),
		Dmarc:       cloudflare.F([]cloudflare.RadarEmailTimeserySpfListParamsDmarc{cloudflare.RadarEmailTimeserySpfListParamsDmarcPass, cloudflare.RadarEmailTimeserySpfListParamsDmarcNone, cloudflare.RadarEmailTimeserySpfListParamsDmarcFail}),
		Format:      cloudflare.F(cloudflare.RadarEmailTimeserySpfListParamsFormatJson),
		Location:    cloudflare.F([]string{"US,CA", "US,CA", "US,CA"}),
		Name:        cloudflare.F([]string{"main_series", "main_series", "main_series"}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
