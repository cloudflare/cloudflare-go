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

func TestRadarEmailSecurityTopAseDkimGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Emails.Security.Top.Ases.Dkim.Get(
		context.TODO(),
		cloudflare.RadarEmailSecurityTopAseDkimGetParamsDkimPass,
		cloudflare.RadarEmailSecurityTopAseDkimGetParams{
			Arc:       cloudflare.F([]cloudflare.RadarEmailSecurityTopAseDkimGetParamsArc{cloudflare.RadarEmailSecurityTopAseDkimGetParamsArcPass, cloudflare.RadarEmailSecurityTopAseDkimGetParamsArcNone, cloudflare.RadarEmailSecurityTopAseDkimGetParamsArcFail}),
			Asn:       cloudflare.F([]string{"string", "string", "string"}),
			DateEnd:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange: cloudflare.F([]cloudflare.RadarEmailSecurityTopAseDkimGetParamsDateRange{cloudflare.RadarEmailSecurityTopAseDkimGetParamsDateRange1d, cloudflare.RadarEmailSecurityTopAseDkimGetParamsDateRange2d, cloudflare.RadarEmailSecurityTopAseDkimGetParamsDateRange7d}),
			DateStart: cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			Dmarc:     cloudflare.F([]cloudflare.RadarEmailSecurityTopAseDkimGetParamsDmarc{cloudflare.RadarEmailSecurityTopAseDkimGetParamsDmarcPass, cloudflare.RadarEmailSecurityTopAseDkimGetParamsDmarcNone, cloudflare.RadarEmailSecurityTopAseDkimGetParamsDmarcFail}),
			Format:    cloudflare.F(cloudflare.RadarEmailSecurityTopAseDkimGetParamsFormatJson),
			Limit:     cloudflare.F(int64(5)),
			Location:  cloudflare.F([]string{"string", "string", "string"}),
			Name:      cloudflare.F([]string{"string", "string", "string"}),
			Spf:       cloudflare.F([]cloudflare.RadarEmailSecurityTopAseDkimGetParamsSpf{cloudflare.RadarEmailSecurityTopAseDkimGetParamsSpfPass, cloudflare.RadarEmailSecurityTopAseDkimGetParamsSpfNone, cloudflare.RadarEmailSecurityTopAseDkimGetParamsSpfFail}),
		},
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
