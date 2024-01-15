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

func TestRadarEmailSecurityDkimTimeseryListWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.DkimTimeseries.List(context.TODO(), cloudflare.RadarEmailSecurityDkimTimeseryListParams{
		AggInterval: cloudflare.F(cloudflare.RadarEmailSecurityDkimTimeseryListParamsAggInterval1h),
		Arc:         cloudflare.F([]cloudflare.RadarEmailSecurityDkimTimeseryListParamsArc{cloudflare.RadarEmailSecurityDkimTimeseryListParamsArcPass, cloudflare.RadarEmailSecurityDkimTimeseryListParamsArcNone, cloudflare.RadarEmailSecurityDkimTimeseryListParamsArcFail}),
		ASN:         cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   cloudflare.F([]cloudflare.RadarEmailSecurityDkimTimeseryListParamsDateRange{cloudflare.RadarEmailSecurityDkimTimeseryListParamsDateRange1d, cloudflare.RadarEmailSecurityDkimTimeseryListParamsDateRange2d, cloudflare.RadarEmailSecurityDkimTimeseryListParamsDateRange7d}),
		DateStart:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Dmarc:       cloudflare.F([]cloudflare.RadarEmailSecurityDkimTimeseryListParamsDmarc{cloudflare.RadarEmailSecurityDkimTimeseryListParamsDmarcPass, cloudflare.RadarEmailSecurityDkimTimeseryListParamsDmarcNone, cloudflare.RadarEmailSecurityDkimTimeseryListParamsDmarcFail}),
		Format:      cloudflare.F(cloudflare.RadarEmailSecurityDkimTimeseryListParamsFormatJson),
		Location:    cloudflare.F([]string{"string", "string", "string"}),
		Name:        cloudflare.F([]string{"string", "string", "string"}),
		Spf:         cloudflare.F([]cloudflare.RadarEmailSecurityDkimTimeseryListParamsSpf{cloudflare.RadarEmailSecurityDkimTimeseryListParamsSpfPass, cloudflare.RadarEmailSecurityDkimTimeseryListParamsSpfNone, cloudflare.RadarEmailSecurityDkimTimeseryListParamsSpfFail}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
