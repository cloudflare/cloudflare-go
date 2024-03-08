// File generated from our OpenAPI spec by Stainless.

package cloudflare_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/cloudflare-go/internal/testutil"
	"github.com/cloudflare/cloudflare-go/option"
)

func TestRadarEmailSecurityTopTldGetWithOptionalParams(t *testing.T) {
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
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.Radar.Email.Security.Top.Tlds.Get(context.TODO(), cloudflare.RadarEmailSecurityTopTldGetParams{
		ARC:         cloudflare.F([]cloudflare.RadarEmailSecurityTopTldGetParamsARC{cloudflare.RadarEmailSecurityTopTldGetParamsARCPass, cloudflare.RadarEmailSecurityTopTldGetParamsARCNone, cloudflare.RadarEmailSecurityTopTldGetParamsARCFail}),
		DateEnd:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   cloudflare.F([]cloudflare.RadarEmailSecurityTopTldGetParamsDateRange{cloudflare.RadarEmailSecurityTopTldGetParamsDateRange1d, cloudflare.RadarEmailSecurityTopTldGetParamsDateRange2d, cloudflare.RadarEmailSecurityTopTldGetParamsDateRange7d}),
		DateStart:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:        cloudflare.F([]cloudflare.RadarEmailSecurityTopTldGetParamsDKIM{cloudflare.RadarEmailSecurityTopTldGetParamsDKIMPass, cloudflare.RadarEmailSecurityTopTldGetParamsDKIMNone, cloudflare.RadarEmailSecurityTopTldGetParamsDKIMFail}),
		DMARC:       cloudflare.F([]cloudflare.RadarEmailSecurityTopTldGetParamsDMARC{cloudflare.RadarEmailSecurityTopTldGetParamsDMARCPass, cloudflare.RadarEmailSecurityTopTldGetParamsDMARCNone, cloudflare.RadarEmailSecurityTopTldGetParamsDMARCFail}),
		Format:      cloudflare.F(cloudflare.RadarEmailSecurityTopTldGetParamsFormatJson),
		Limit:       cloudflare.F(int64(5)),
		Name:        cloudflare.F([]string{"string", "string", "string"}),
		SPF:         cloudflare.F([]cloudflare.RadarEmailSecurityTopTldGetParamsSPF{cloudflare.RadarEmailSecurityTopTldGetParamsSPFPass, cloudflare.RadarEmailSecurityTopTldGetParamsSPFNone, cloudflare.RadarEmailSecurityTopTldGetParamsSPFFail}),
		TldCategory: cloudflare.F(cloudflare.RadarEmailSecurityTopTldGetParamsTldCategoryClassic),
		TLSVersion:  cloudflare.F([]cloudflare.RadarEmailSecurityTopTldGetParamsTLSVersion{cloudflare.RadarEmailSecurityTopTldGetParamsTLSVersionTlSv1_0, cloudflare.RadarEmailSecurityTopTldGetParamsTLSVersionTlSv1_1, cloudflare.RadarEmailSecurityTopTldGetParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
