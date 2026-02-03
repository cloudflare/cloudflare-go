// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/radar"
)

func TestEmailSecurityTopTLDMaliciousGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.Top.TLDs.Malicious.Get(
		context.TODO(),
		radar.EmailSecurityTopTLDMaliciousGetParamsMaliciousMalicious,
		radar.EmailSecurityTopTLDMaliciousGetParams{
			ARC:         cloudflare.F([]radar.EmailSecurityTopTLDMaliciousGetParamsARC{radar.EmailSecurityTopTLDMaliciousGetParamsARCPass}),
			DateEnd:     cloudflare.F([]time.Time{time.Now()}),
			DateRange:   cloudflare.F([]string{"7d"}),
			DateStart:   cloudflare.F([]time.Time{time.Now()}),
			DKIM:        cloudflare.F([]radar.EmailSecurityTopTLDMaliciousGetParamsDKIM{radar.EmailSecurityTopTLDMaliciousGetParamsDKIMPass}),
			DMARC:       cloudflare.F([]radar.EmailSecurityTopTLDMaliciousGetParamsDMARC{radar.EmailSecurityTopTLDMaliciousGetParamsDMARCPass}),
			Format:      cloudflare.F(radar.EmailSecurityTopTLDMaliciousGetParamsFormatJson),
			Limit:       cloudflare.F(int64(1)),
			Name:        cloudflare.F([]string{"main_series"}),
			SPF:         cloudflare.F([]radar.EmailSecurityTopTLDMaliciousGetParamsSPF{radar.EmailSecurityTopTLDMaliciousGetParamsSPFPass}),
			TLDCategory: cloudflare.F(radar.EmailSecurityTopTLDMaliciousGetParamsTLDCategoryClassic),
			TLSVersion:  cloudflare.F([]radar.EmailSecurityTopTLDMaliciousGetParamsTLSVersion{radar.EmailSecurityTopTLDMaliciousGetParamsTLSVersionTlSv1_0}),
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
