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

func TestRadarEmailSecurityTopTldSpoofGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.Top.Tlds.Spoof.Get(
		context.TODO(),
		cloudflare.RadarEmailSecurityTopTldSpoofGetParamsSpoofSpoof,
		cloudflare.RadarEmailSecurityTopTldSpoofGetParams{
			ARC:         cloudflare.F([]cloudflare.RadarEmailSecurityTopTldSpoofGetParamsARC{cloudflare.RadarEmailSecurityTopTldSpoofGetParamsARCPass, cloudflare.RadarEmailSecurityTopTldSpoofGetParamsARCNone, cloudflare.RadarEmailSecurityTopTldSpoofGetParamsARCFail}),
			DateEnd:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange:   cloudflare.F([]cloudflare.RadarEmailSecurityTopTldSpoofGetParamsDateRange{cloudflare.RadarEmailSecurityTopTldSpoofGetParamsDateRange1d, cloudflare.RadarEmailSecurityTopTldSpoofGetParamsDateRange2d, cloudflare.RadarEmailSecurityTopTldSpoofGetParamsDateRange7d}),
			DateStart:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DKIM:        cloudflare.F([]cloudflare.RadarEmailSecurityTopTldSpoofGetParamsDKIM{cloudflare.RadarEmailSecurityTopTldSpoofGetParamsDKIMPass, cloudflare.RadarEmailSecurityTopTldSpoofGetParamsDKIMNone, cloudflare.RadarEmailSecurityTopTldSpoofGetParamsDKIMFail}),
			DMARC:       cloudflare.F([]cloudflare.RadarEmailSecurityTopTldSpoofGetParamsDMARC{cloudflare.RadarEmailSecurityTopTldSpoofGetParamsDMARCPass, cloudflare.RadarEmailSecurityTopTldSpoofGetParamsDMARCNone, cloudflare.RadarEmailSecurityTopTldSpoofGetParamsDMARCFail}),
			Format:      cloudflare.F(cloudflare.RadarEmailSecurityTopTldSpoofGetParamsFormatJson),
			Limit:       cloudflare.F(int64(5)),
			Name:        cloudflare.F([]string{"string", "string", "string"}),
			SPF:         cloudflare.F([]cloudflare.RadarEmailSecurityTopTldSpoofGetParamsSPF{cloudflare.RadarEmailSecurityTopTldSpoofGetParamsSPFPass, cloudflare.RadarEmailSecurityTopTldSpoofGetParamsSPFNone, cloudflare.RadarEmailSecurityTopTldSpoofGetParamsSPFFail}),
			TldCategory: cloudflare.F(cloudflare.RadarEmailSecurityTopTldSpoofGetParamsTldCategoryClassic),
			TLSVersion:  cloudflare.F([]cloudflare.RadarEmailSecurityTopTldSpoofGetParamsTLSVersion{cloudflare.RadarEmailSecurityTopTldSpoofGetParamsTLSVersionTlSv1_0, cloudflare.RadarEmailSecurityTopTldSpoofGetParamsTLSVersionTlSv1_1, cloudflare.RadarEmailSecurityTopTldSpoofGetParamsTLSVersionTlSv1_2}),
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
