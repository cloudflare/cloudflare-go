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

func TestRadarEmailSecurityArcTimeseryListWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.ArcTimeseries.List(context.TODO(), cloudflare.RadarEmailSecurityArcTimeseryListParams{
		AggInterval: cloudflare.F(cloudflare.RadarEmailSecurityArcTimeseryListParamsAggInterval1h),
		ASN:         cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   cloudflare.F([]cloudflare.RadarEmailSecurityArcTimeseryListParamsDateRange{cloudflare.RadarEmailSecurityArcTimeseryListParamsDateRange1d, cloudflare.RadarEmailSecurityArcTimeseryListParamsDateRange2d, cloudflare.RadarEmailSecurityArcTimeseryListParamsDateRange7d}),
		DateStart:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Dkim:        cloudflare.F([]cloudflare.RadarEmailSecurityArcTimeseryListParamsDkim{cloudflare.RadarEmailSecurityArcTimeseryListParamsDkimPass, cloudflare.RadarEmailSecurityArcTimeseryListParamsDkimNone, cloudflare.RadarEmailSecurityArcTimeseryListParamsDkimFail}),
		Dmarc:       cloudflare.F([]cloudflare.RadarEmailSecurityArcTimeseryListParamsDmarc{cloudflare.RadarEmailSecurityArcTimeseryListParamsDmarcPass, cloudflare.RadarEmailSecurityArcTimeseryListParamsDmarcNone, cloudflare.RadarEmailSecurityArcTimeseryListParamsDmarcFail}),
		Format:      cloudflare.F(cloudflare.RadarEmailSecurityArcTimeseryListParamsFormatJson),
		Location:    cloudflare.F([]string{"string", "string", "string"}),
		Name:        cloudflare.F([]string{"string", "string", "string"}),
		Spf:         cloudflare.F([]cloudflare.RadarEmailSecurityArcTimeseryListParamsSpf{cloudflare.RadarEmailSecurityArcTimeseryListParamsSpfPass, cloudflare.RadarEmailSecurityArcTimeseryListParamsSpfNone, cloudflare.RadarEmailSecurityArcTimeseryListParamsSpfFail}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
