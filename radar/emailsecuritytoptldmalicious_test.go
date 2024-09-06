// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/radar"
)

func TestEmailSecurityTopTldMaliciousGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.Top.Tlds.Malicious.Get(
		context.TODO(),
		radar.EmailSecurityTopTldMaliciousGetParamsMaliciousMalicious,
		radar.EmailSecurityTopTldMaliciousGetParams{
			ARC:         cloudflare.F([]radar.EmailSecurityTopTldMaliciousGetParamsARC{radar.EmailSecurityTopTldMaliciousGetParamsARCPass, radar.EmailSecurityTopTldMaliciousGetParamsARCNone, radar.EmailSecurityTopTldMaliciousGetParamsARCFail}),
			DateEnd:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange:   cloudflare.F([]string{"7d", "7d", "7d"}),
			DateStart:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DKIM:        cloudflare.F([]radar.EmailSecurityTopTldMaliciousGetParamsDKIM{radar.EmailSecurityTopTldMaliciousGetParamsDKIMPass, radar.EmailSecurityTopTldMaliciousGetParamsDKIMNone, radar.EmailSecurityTopTldMaliciousGetParamsDKIMFail}),
			DMARC:       cloudflare.F([]radar.EmailSecurityTopTldMaliciousGetParamsDMARC{radar.EmailSecurityTopTldMaliciousGetParamsDMARCPass, radar.EmailSecurityTopTldMaliciousGetParamsDMARCNone, radar.EmailSecurityTopTldMaliciousGetParamsDMARCFail}),
			Format:      cloudflare.F(radar.EmailSecurityTopTldMaliciousGetParamsFormatJson),
			Limit:       cloudflare.F(int64(5)),
			Name:        cloudflare.F([]string{"string", "string", "string"}),
			SPF:         cloudflare.F([]radar.EmailSecurityTopTldMaliciousGetParamsSPF{radar.EmailSecurityTopTldMaliciousGetParamsSPFPass, radar.EmailSecurityTopTldMaliciousGetParamsSPFNone, radar.EmailSecurityTopTldMaliciousGetParamsSPFFail}),
			TldCategory: cloudflare.F(radar.EmailSecurityTopTldMaliciousGetParamsTldCategoryClassic),
			TLSVersion:  cloudflare.F([]radar.EmailSecurityTopTldMaliciousGetParamsTLSVersion{radar.EmailSecurityTopTldMaliciousGetParamsTLSVersionTlSv1_0, radar.EmailSecurityTopTldMaliciousGetParamsTLSVersionTlSv1_1, radar.EmailSecurityTopTldMaliciousGetParamsTLSVersionTlSv1_2}),
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
