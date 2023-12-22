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

func TestRadarEmailSummaryThreatCategoryListWithOptionalParams(t *testing.T) {
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
	_, err := client.Radars.Emails.Summaries.ThreatCategories.List(context.TODO(), cloudflare.RadarEmailSummaryThreatCategoryListParams{
		Arc:       cloudflare.F([]cloudflare.RadarEmailSummaryThreatCategoryListParamsArc{cloudflare.RadarEmailSummaryThreatCategoryListParamsArcPass, cloudflare.RadarEmailSummaryThreatCategoryListParamsArcNone, cloudflare.RadarEmailSummaryThreatCategoryListParamsArcFail}),
		ASN:       cloudflare.F([]string{"15169", "15169", "15169"}),
		DateEnd:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange: cloudflare.F([]cloudflare.RadarEmailSummaryThreatCategoryListParamsDateRange{cloudflare.RadarEmailSummaryThreatCategoryListParamsDateRange1d, cloudflare.RadarEmailSummaryThreatCategoryListParamsDateRange7d, cloudflare.RadarEmailSummaryThreatCategoryListParamsDateRange14d}),
		DateStart: cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Dkim:      cloudflare.F([]cloudflare.RadarEmailSummaryThreatCategoryListParamsDkim{cloudflare.RadarEmailSummaryThreatCategoryListParamsDkimPass, cloudflare.RadarEmailSummaryThreatCategoryListParamsDkimNone, cloudflare.RadarEmailSummaryThreatCategoryListParamsDkimFail}),
		Dmarc:     cloudflare.F([]cloudflare.RadarEmailSummaryThreatCategoryListParamsDmarc{cloudflare.RadarEmailSummaryThreatCategoryListParamsDmarcPass, cloudflare.RadarEmailSummaryThreatCategoryListParamsDmarcNone, cloudflare.RadarEmailSummaryThreatCategoryListParamsDmarcFail}),
		Format:    cloudflare.F(cloudflare.RadarEmailSummaryThreatCategoryListParamsFormatJson),
		Location:  cloudflare.F([]string{"US,CA", "US,CA", "US,CA"}),
		Name:      cloudflare.F([]string{"main_series", "main_series", "main_series"}),
		Spf:       cloudflare.F([]cloudflare.RadarEmailSummaryThreatCategoryListParamsSpf{cloudflare.RadarEmailSummaryThreatCategoryListParamsSpfPass, cloudflare.RadarEmailSummaryThreatCategoryListParamsSpfNone, cloudflare.RadarEmailSummaryThreatCategoryListParamsSpfFail}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
