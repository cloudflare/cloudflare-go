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

func TestRadarEmailSecuritySPFListWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Emails.Security.SPF.List(context.TODO(), cloudflare.RadarEmailSecuritySPFListParams{
		AggInterval: cloudflare.F(cloudflare.RadarEmailSecuritySPFListParamsAggInterval1h),
		Arc:         cloudflare.F([]cloudflare.RadarEmailSecuritySPFListParamsArc{cloudflare.RadarEmailSecuritySPFListParamsArcPass, cloudflare.RadarEmailSecuritySPFListParamsArcNone, cloudflare.RadarEmailSecuritySPFListParamsArcFail}),
		Asn:         cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   cloudflare.F([]cloudflare.RadarEmailSecuritySPFListParamsDateRange{cloudflare.RadarEmailSecuritySPFListParamsDateRange1d, cloudflare.RadarEmailSecuritySPFListParamsDateRange2d, cloudflare.RadarEmailSecuritySPFListParamsDateRange7d}),
		DateStart:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:        cloudflare.F([]cloudflare.RadarEmailSecuritySPFListParamsDKIM{cloudflare.RadarEmailSecuritySPFListParamsDKIMPass, cloudflare.RadarEmailSecuritySPFListParamsDKIMNone, cloudflare.RadarEmailSecuritySPFListParamsDKIMFail}),
		Dmarc:       cloudflare.F([]cloudflare.RadarEmailSecuritySPFListParamsDmarc{cloudflare.RadarEmailSecuritySPFListParamsDmarcPass, cloudflare.RadarEmailSecuritySPFListParamsDmarcNone, cloudflare.RadarEmailSecuritySPFListParamsDmarcFail}),
		Format:      cloudflare.F(cloudflare.RadarEmailSecuritySPFListParamsFormatJson),
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
