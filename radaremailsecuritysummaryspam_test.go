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

func TestRadarEmailSecuritySummarySpamGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.Summaries.Spams.Get(context.TODO(), cloudflare.RadarEmailSecuritySummarySpamGetParams{
		Arc:       cloudflare.F([]cloudflare.RadarEmailSecuritySummarySpamGetParamsArc{cloudflare.RadarEmailSecuritySummarySpamGetParamsArcPass, cloudflare.RadarEmailSecuritySummarySpamGetParamsArcNone, cloudflare.RadarEmailSecuritySummarySpamGetParamsArcFail}),
		Asn:       cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange: cloudflare.F([]cloudflare.RadarEmailSecuritySummarySpamGetParamsDateRange{cloudflare.RadarEmailSecuritySummarySpamGetParamsDateRange1d, cloudflare.RadarEmailSecuritySummarySpamGetParamsDateRange2d, cloudflare.RadarEmailSecuritySummarySpamGetParamsDateRange7d}),
		DateStart: cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:      cloudflare.F([]cloudflare.RadarEmailSecuritySummarySpamGetParamsDKIM{cloudflare.RadarEmailSecuritySummarySpamGetParamsDKIMPass, cloudflare.RadarEmailSecuritySummarySpamGetParamsDKIMNone, cloudflare.RadarEmailSecuritySummarySpamGetParamsDKIMFail}),
		Dmarc:     cloudflare.F([]cloudflare.RadarEmailSecuritySummarySpamGetParamsDmarc{cloudflare.RadarEmailSecuritySummarySpamGetParamsDmarcPass, cloudflare.RadarEmailSecuritySummarySpamGetParamsDmarcNone, cloudflare.RadarEmailSecuritySummarySpamGetParamsDmarcFail}),
		Format:    cloudflare.F(cloudflare.RadarEmailSecuritySummarySpamGetParamsFormatJson),
		Location:  cloudflare.F([]string{"string", "string", "string"}),
		Name:      cloudflare.F([]string{"string", "string", "string"}),
		SPF:       cloudflare.F([]cloudflare.RadarEmailSecuritySummarySpamGetParamsSPF{cloudflare.RadarEmailSecuritySummarySpamGetParamsSPFPass, cloudflare.RadarEmailSecuritySummarySpamGetParamsSPFNone, cloudflare.RadarEmailSecuritySummarySpamGetParamsSPFFail}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
