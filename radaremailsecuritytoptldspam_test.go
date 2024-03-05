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

func TestRadarEmailSecurityTopTldSpamGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.Top.Tlds.Spam.Get(
		context.TODO(),
		cloudflare.RadarEmailSecurityTopTldSpamGetParamsSpamSpam,
		cloudflare.RadarEmailSecurityTopTldSpamGetParams{
			ARC:         cloudflare.F([]cloudflare.RadarEmailSecurityTopTldSpamGetParamsARC{cloudflare.RadarEmailSecurityTopTldSpamGetParamsARCPass, cloudflare.RadarEmailSecurityTopTldSpamGetParamsARCNone, cloudflare.RadarEmailSecurityTopTldSpamGetParamsARCFail}),
			DateEnd:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange:   cloudflare.F([]cloudflare.RadarEmailSecurityTopTldSpamGetParamsDateRange{cloudflare.RadarEmailSecurityTopTldSpamGetParamsDateRange1d, cloudflare.RadarEmailSecurityTopTldSpamGetParamsDateRange2d, cloudflare.RadarEmailSecurityTopTldSpamGetParamsDateRange7d}),
			DateStart:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DKIM:        cloudflare.F([]cloudflare.RadarEmailSecurityTopTldSpamGetParamsDKIM{cloudflare.RadarEmailSecurityTopTldSpamGetParamsDKIMPass, cloudflare.RadarEmailSecurityTopTldSpamGetParamsDKIMNone, cloudflare.RadarEmailSecurityTopTldSpamGetParamsDKIMFail}),
			DMARC:       cloudflare.F([]cloudflare.RadarEmailSecurityTopTldSpamGetParamsDMARC{cloudflare.RadarEmailSecurityTopTldSpamGetParamsDMARCPass, cloudflare.RadarEmailSecurityTopTldSpamGetParamsDMARCNone, cloudflare.RadarEmailSecurityTopTldSpamGetParamsDMARCFail}),
			Format:      cloudflare.F(cloudflare.RadarEmailSecurityTopTldSpamGetParamsFormatJson),
			Limit:       cloudflare.F(int64(5)),
			Name:        cloudflare.F([]string{"string", "string", "string"}),
			SPF:         cloudflare.F([]cloudflare.RadarEmailSecurityTopTldSpamGetParamsSPF{cloudflare.RadarEmailSecurityTopTldSpamGetParamsSPFPass, cloudflare.RadarEmailSecurityTopTldSpamGetParamsSPFNone, cloudflare.RadarEmailSecurityTopTldSpamGetParamsSPFFail}),
			TldCategory: cloudflare.F(cloudflare.RadarEmailSecurityTopTldSpamGetParamsTldCategoryClassic),
			TLSVersion:  cloudflare.F([]cloudflare.RadarEmailSecurityTopTldSpamGetParamsTLSVersion{cloudflare.RadarEmailSecurityTopTldSpamGetParamsTLSVersionTlSv1_0, cloudflare.RadarEmailSecurityTopTldSpamGetParamsTLSVersionTlSv1_1, cloudflare.RadarEmailSecurityTopTldSpamGetParamsTLSVersionTlSv1_2}),
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
