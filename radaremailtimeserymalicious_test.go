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

func TestRadarEmailTimeseryMaliciousListWithOptionalParams(t *testing.T) {
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
	_, err := client.Radars.Emails.Timeseries.Malicious.List(context.TODO(), cloudflare.RadarEmailTimeseryMaliciousListParams{
		AggInterval: cloudflare.F(cloudflare.RadarEmailTimeseryMaliciousListParamsAggInterval15m),
		Arc:         cloudflare.F([]cloudflare.RadarEmailTimeseryMaliciousListParamsArc{cloudflare.RadarEmailTimeseryMaliciousListParamsArcPass, cloudflare.RadarEmailTimeseryMaliciousListParamsArcNone, cloudflare.RadarEmailTimeseryMaliciousListParamsArcFail}),
		ASN:         cloudflare.F([]string{"15169", "15169", "15169"}),
		DateEnd:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   cloudflare.F([]cloudflare.RadarEmailTimeseryMaliciousListParamsDateRange{cloudflare.RadarEmailTimeseryMaliciousListParamsDateRange1d, cloudflare.RadarEmailTimeseryMaliciousListParamsDateRange7d, cloudflare.RadarEmailTimeseryMaliciousListParamsDateRange14d}),
		DateStart:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Dkim:        cloudflare.F([]cloudflare.RadarEmailTimeseryMaliciousListParamsDkim{cloudflare.RadarEmailTimeseryMaliciousListParamsDkimPass, cloudflare.RadarEmailTimeseryMaliciousListParamsDkimNone, cloudflare.RadarEmailTimeseryMaliciousListParamsDkimFail}),
		Dmarc:       cloudflare.F([]cloudflare.RadarEmailTimeseryMaliciousListParamsDmarc{cloudflare.RadarEmailTimeseryMaliciousListParamsDmarcPass, cloudflare.RadarEmailTimeseryMaliciousListParamsDmarcNone, cloudflare.RadarEmailTimeseryMaliciousListParamsDmarcFail}),
		Format:      cloudflare.F(cloudflare.RadarEmailTimeseryMaliciousListParamsFormatJson),
		Location:    cloudflare.F([]string{"US,CA", "US,CA", "US,CA"}),
		Name:        cloudflare.F([]string{"main_series", "main_series", "main_series"}),
		Spf:         cloudflare.F([]cloudflare.RadarEmailTimeseryMaliciousListParamsSpf{cloudflare.RadarEmailTimeseryMaliciousListParamsSpfPass, cloudflare.RadarEmailTimeseryMaliciousListParamsSpfNone, cloudflare.RadarEmailTimeseryMaliciousListParamsSpfFail}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
