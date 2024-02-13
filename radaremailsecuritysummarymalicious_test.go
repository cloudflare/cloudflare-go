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

func TestRadarEmailSecuritySummaryMaliciousListWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.Summaries.Malicious.List(context.TODO(), cloudflare.RadarEmailSecuritySummaryMaliciousListParams{
		Arc:       cloudflare.F([]cloudflare.RadarEmailSecuritySummaryMaliciousListParamsArc{cloudflare.RadarEmailSecuritySummaryMaliciousListParamsArcPass, cloudflare.RadarEmailSecuritySummaryMaliciousListParamsArcNone, cloudflare.RadarEmailSecuritySummaryMaliciousListParamsArcFail}),
		Asn:       cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange: cloudflare.F([]cloudflare.RadarEmailSecuritySummaryMaliciousListParamsDateRange{cloudflare.RadarEmailSecuritySummaryMaliciousListParamsDateRange1d, cloudflare.RadarEmailSecuritySummaryMaliciousListParamsDateRange2d, cloudflare.RadarEmailSecuritySummaryMaliciousListParamsDateRange7d}),
		DateStart: cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:      cloudflare.F([]cloudflare.RadarEmailSecuritySummaryMaliciousListParamsDKIM{cloudflare.RadarEmailSecuritySummaryMaliciousListParamsDKIMPass, cloudflare.RadarEmailSecuritySummaryMaliciousListParamsDKIMNone, cloudflare.RadarEmailSecuritySummaryMaliciousListParamsDKIMFail}),
		Dmarc:     cloudflare.F([]cloudflare.RadarEmailSecuritySummaryMaliciousListParamsDmarc{cloudflare.RadarEmailSecuritySummaryMaliciousListParamsDmarcPass, cloudflare.RadarEmailSecuritySummaryMaliciousListParamsDmarcNone, cloudflare.RadarEmailSecuritySummaryMaliciousListParamsDmarcFail}),
		Format:    cloudflare.F(cloudflare.RadarEmailSecuritySummaryMaliciousListParamsFormatJson),
		Location:  cloudflare.F([]string{"string", "string", "string"}),
		Name:      cloudflare.F([]string{"string", "string", "string"}),
		SPF:       cloudflare.F([]cloudflare.RadarEmailSecuritySummaryMaliciousListParamsSPF{cloudflare.RadarEmailSecuritySummaryMaliciousListParamsSPFPass, cloudflare.RadarEmailSecuritySummaryMaliciousListParamsSPFNone, cloudflare.RadarEmailSecuritySummaryMaliciousListParamsSPFFail}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
