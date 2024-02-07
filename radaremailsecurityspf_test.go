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

func TestRadarEmailSecuritySpfListWithOptionalParams(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("My User Service Key"),
	)
	_, err := client.Radar.Emails.Security.Spf.List(context.TODO(), cloudflare.RadarEmailSecuritySpfListParams{
		AggInterval: cloudflare.F(cloudflare.RadarEmailSecuritySpfListParamsAggInterval1h),
		Arc:         cloudflare.F([]cloudflare.RadarEmailSecuritySpfListParamsArc{cloudflare.RadarEmailSecuritySpfListParamsArcPass, cloudflare.RadarEmailSecuritySpfListParamsArcNone, cloudflare.RadarEmailSecuritySpfListParamsArcFail}),
		Asn:         cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   cloudflare.F([]cloudflare.RadarEmailSecuritySpfListParamsDateRange{cloudflare.RadarEmailSecuritySpfListParamsDateRange1d, cloudflare.RadarEmailSecuritySpfListParamsDateRange2d, cloudflare.RadarEmailSecuritySpfListParamsDateRange7d}),
		DateStart:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Dkim:        cloudflare.F([]cloudflare.RadarEmailSecuritySpfListParamsDkim{cloudflare.RadarEmailSecuritySpfListParamsDkimPass, cloudflare.RadarEmailSecuritySpfListParamsDkimNone, cloudflare.RadarEmailSecuritySpfListParamsDkimFail}),
		Dmarc:       cloudflare.F([]cloudflare.RadarEmailSecuritySpfListParamsDmarc{cloudflare.RadarEmailSecuritySpfListParamsDmarcPass, cloudflare.RadarEmailSecuritySpfListParamsDmarcNone, cloudflare.RadarEmailSecuritySpfListParamsDmarcFail}),
		Format:      cloudflare.F(cloudflare.RadarEmailSecuritySpfListParamsFormatJson),
		Location:    cloudflare.F([]string{"string", "string", "string"}),
		Name:        cloudflare.F([]string{"string", "string", "string"}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
