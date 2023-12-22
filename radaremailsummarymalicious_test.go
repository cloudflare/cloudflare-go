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

func TestRadarEmailSummaryMaliciousListWithOptionalParams(t *testing.T) {
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
	_, err := client.Radars.Emails.Summaries.Malicious.List(context.TODO(), cloudflare.RadarEmailSummaryMaliciousListParams{
		Arc:       cloudflare.F([]cloudflare.RadarEmailSummaryMaliciousListParamsArc{cloudflare.RadarEmailSummaryMaliciousListParamsArcPass, cloudflare.RadarEmailSummaryMaliciousListParamsArcNone, cloudflare.RadarEmailSummaryMaliciousListParamsArcFail}),
		ASN:       cloudflare.F([]string{"15169", "15169", "15169"}),
		DateEnd:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange: cloudflare.F([]cloudflare.RadarEmailSummaryMaliciousListParamsDateRange{cloudflare.RadarEmailSummaryMaliciousListParamsDateRange1d, cloudflare.RadarEmailSummaryMaliciousListParamsDateRange7d, cloudflare.RadarEmailSummaryMaliciousListParamsDateRange14d}),
		DateStart: cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Dkim:      cloudflare.F([]cloudflare.RadarEmailSummaryMaliciousListParamsDkim{cloudflare.RadarEmailSummaryMaliciousListParamsDkimPass, cloudflare.RadarEmailSummaryMaliciousListParamsDkimNone, cloudflare.RadarEmailSummaryMaliciousListParamsDkimFail}),
		Dmarc:     cloudflare.F([]cloudflare.RadarEmailSummaryMaliciousListParamsDmarc{cloudflare.RadarEmailSummaryMaliciousListParamsDmarcPass, cloudflare.RadarEmailSummaryMaliciousListParamsDmarcNone, cloudflare.RadarEmailSummaryMaliciousListParamsDmarcFail}),
		Format:    cloudflare.F(cloudflare.RadarEmailSummaryMaliciousListParamsFormatJson),
		Location:  cloudflare.F([]string{"US,CA", "US,CA", "US,CA"}),
		Name:      cloudflare.F([]string{"main_series", "main_series", "main_series"}),
		Spf:       cloudflare.F([]cloudflare.RadarEmailSummaryMaliciousListParamsSpf{cloudflare.RadarEmailSummaryMaliciousListParamsSpfPass, cloudflare.RadarEmailSummaryMaliciousListParamsSpfNone, cloudflare.RadarEmailSummaryMaliciousListParamsSpfFail}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
