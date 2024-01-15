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

func TestRadarEmailSecurityMaliciousTimeseryListWithOptionalParams(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithAPIToken("my-cloudflare-api-token"),
		option.WithUserServiceKey("my-cloudflare-user-service-key"),
	)
	_, err := client.Radar.Email.Security.MaliciousTimeseries.List(context.TODO(), cloudflare.RadarEmailSecurityMaliciousTimeseryListParams{
		AggInterval: cloudflare.F(cloudflare.RadarEmailSecurityMaliciousTimeseryListParamsAggInterval1h),
		Arc:         cloudflare.F([]cloudflare.RadarEmailSecurityMaliciousTimeseryListParamsArc{cloudflare.RadarEmailSecurityMaliciousTimeseryListParamsArcPass, cloudflare.RadarEmailSecurityMaliciousTimeseryListParamsArcNone, cloudflare.RadarEmailSecurityMaliciousTimeseryListParamsArcFail}),
		ASN:         cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   cloudflare.F([]cloudflare.RadarEmailSecurityMaliciousTimeseryListParamsDateRange{cloudflare.RadarEmailSecurityMaliciousTimeseryListParamsDateRange1d, cloudflare.RadarEmailSecurityMaliciousTimeseryListParamsDateRange2d, cloudflare.RadarEmailSecurityMaliciousTimeseryListParamsDateRange7d}),
		DateStart:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Dkim:        cloudflare.F([]cloudflare.RadarEmailSecurityMaliciousTimeseryListParamsDkim{cloudflare.RadarEmailSecurityMaliciousTimeseryListParamsDkimPass, cloudflare.RadarEmailSecurityMaliciousTimeseryListParamsDkimNone, cloudflare.RadarEmailSecurityMaliciousTimeseryListParamsDkimFail}),
		Dmarc:       cloudflare.F([]cloudflare.RadarEmailSecurityMaliciousTimeseryListParamsDmarc{cloudflare.RadarEmailSecurityMaliciousTimeseryListParamsDmarcPass, cloudflare.RadarEmailSecurityMaliciousTimeseryListParamsDmarcNone, cloudflare.RadarEmailSecurityMaliciousTimeseryListParamsDmarcFail}),
		Format:      cloudflare.F(cloudflare.RadarEmailSecurityMaliciousTimeseryListParamsFormatJson),
		Location:    cloudflare.F([]string{"string", "string", "string"}),
		Name:        cloudflare.F([]string{"string", "string", "string"}),
		Spf:         cloudflare.F([]cloudflare.RadarEmailSecurityMaliciousTimeseryListParamsSpf{cloudflare.RadarEmailSecurityMaliciousTimeseryListParamsSpfPass, cloudflare.RadarEmailSecurityMaliciousTimeseryListParamsSpfNone, cloudflare.RadarEmailSecurityMaliciousTimeseryListParamsSpfFail}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
