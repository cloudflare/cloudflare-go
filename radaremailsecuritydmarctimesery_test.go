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

func TestRadarEmailSecurityDmarcTimeseryListWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.DmarcTimeseries.List(context.TODO(), cloudflare.RadarEmailSecurityDmarcTimeseryListParams{
		AggInterval: cloudflare.F(cloudflare.RadarEmailSecurityDmarcTimeseryListParamsAggInterval1h),
		Arc:         cloudflare.F([]cloudflare.RadarEmailSecurityDmarcTimeseryListParamsArc{cloudflare.RadarEmailSecurityDmarcTimeseryListParamsArcPass, cloudflare.RadarEmailSecurityDmarcTimeseryListParamsArcNone, cloudflare.RadarEmailSecurityDmarcTimeseryListParamsArcFail}),
		ASN:         cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   cloudflare.F([]cloudflare.RadarEmailSecurityDmarcTimeseryListParamsDateRange{cloudflare.RadarEmailSecurityDmarcTimeseryListParamsDateRange1d, cloudflare.RadarEmailSecurityDmarcTimeseryListParamsDateRange2d, cloudflare.RadarEmailSecurityDmarcTimeseryListParamsDateRange7d}),
		DateStart:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Dkim:        cloudflare.F([]cloudflare.RadarEmailSecurityDmarcTimeseryListParamsDkim{cloudflare.RadarEmailSecurityDmarcTimeseryListParamsDkimPass, cloudflare.RadarEmailSecurityDmarcTimeseryListParamsDkimNone, cloudflare.RadarEmailSecurityDmarcTimeseryListParamsDkimFail}),
		Format:      cloudflare.F(cloudflare.RadarEmailSecurityDmarcTimeseryListParamsFormatJson),
		Location:    cloudflare.F([]string{"string", "string", "string"}),
		Name:        cloudflare.F([]string{"string", "string", "string"}),
		Spf:         cloudflare.F([]cloudflare.RadarEmailSecurityDmarcTimeseryListParamsSpf{cloudflare.RadarEmailSecurityDmarcTimeseryListParamsSpfPass, cloudflare.RadarEmailSecurityDmarcTimeseryListParamsSpfNone, cloudflare.RadarEmailSecurityDmarcTimeseryListParamsSpfFail}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
