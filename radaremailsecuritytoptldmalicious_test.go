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

func TestRadarEmailSecurityTopTldMaliciousGetWithOptionalParams(t *testing.T) {
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
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("v1.0-144c9defac04969c7bfad8ef-631a41d003a32d25fe878081ef365c49503f7fada600da935e2851a1c7326084b85cbf6429c4b859de8475731dc92a9c329631e6d59e6c73da7b198497172b4cefe071d90d0f5d2719"),
	)
	_, err := client.Radar.Email.Security.Top.Tlds.Malicious.Get(
		context.TODO(),
		cloudflare.RadarEmailSecurityTopTldMaliciousGetParamsMaliciousMalicious,
		cloudflare.RadarEmailSecurityTopTldMaliciousGetParams{
			ARC:         cloudflare.F([]cloudflare.RadarEmailSecurityTopTldMaliciousGetParamsARC{cloudflare.RadarEmailSecurityTopTldMaliciousGetParamsARCPass, cloudflare.RadarEmailSecurityTopTldMaliciousGetParamsARCNone, cloudflare.RadarEmailSecurityTopTldMaliciousGetParamsARCFail}),
			DateEnd:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange:   cloudflare.F([]cloudflare.RadarEmailSecurityTopTldMaliciousGetParamsDateRange{cloudflare.RadarEmailSecurityTopTldMaliciousGetParamsDateRange1d, cloudflare.RadarEmailSecurityTopTldMaliciousGetParamsDateRange2d, cloudflare.RadarEmailSecurityTopTldMaliciousGetParamsDateRange7d}),
			DateStart:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DKIM:        cloudflare.F([]cloudflare.RadarEmailSecurityTopTldMaliciousGetParamsDKIM{cloudflare.RadarEmailSecurityTopTldMaliciousGetParamsDKIMPass, cloudflare.RadarEmailSecurityTopTldMaliciousGetParamsDKIMNone, cloudflare.RadarEmailSecurityTopTldMaliciousGetParamsDKIMFail}),
			DMARC:       cloudflare.F([]cloudflare.RadarEmailSecurityTopTldMaliciousGetParamsDMARC{cloudflare.RadarEmailSecurityTopTldMaliciousGetParamsDMARCPass, cloudflare.RadarEmailSecurityTopTldMaliciousGetParamsDMARCNone, cloudflare.RadarEmailSecurityTopTldMaliciousGetParamsDMARCFail}),
			Format:      cloudflare.F(cloudflare.RadarEmailSecurityTopTldMaliciousGetParamsFormatJson),
			Limit:       cloudflare.F(int64(5)),
			Name:        cloudflare.F([]string{"string", "string", "string"}),
			SPF:         cloudflare.F([]cloudflare.RadarEmailSecurityTopTldMaliciousGetParamsSPF{cloudflare.RadarEmailSecurityTopTldMaliciousGetParamsSPFPass, cloudflare.RadarEmailSecurityTopTldMaliciousGetParamsSPFNone, cloudflare.RadarEmailSecurityTopTldMaliciousGetParamsSPFFail}),
			TldCategory: cloudflare.F(cloudflare.RadarEmailSecurityTopTldMaliciousGetParamsTldCategoryClassic),
			TLSVersion:  cloudflare.F([]cloudflare.RadarEmailSecurityTopTldMaliciousGetParamsTLSVersion{cloudflare.RadarEmailSecurityTopTldMaliciousGetParamsTLSVersionTlSv1_0, cloudflare.RadarEmailSecurityTopTldMaliciousGetParamsTLSVersionTlSv1_1, cloudflare.RadarEmailSecurityTopTldMaliciousGetParamsTLSVersionTlSv1_2}),
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
