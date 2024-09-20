// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v3/option"
	"github.com/cloudflare/cloudflare-go/v3/radar"
)

func TestEmailSecurityTopTldSpamGetWithOptionalParams(t *testing.T) {
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
		radar.EmailSecurityTopTldSpamGetParamsSpamSpam,
		radar.EmailSecurityTopTldSpamGetParams{
			ARC:         cloudflare.F([]radar.EmailSecurityTopTldSpamGetParamsARC{radar.EmailSecurityTopTldSpamGetParamsARCPass, radar.EmailSecurityTopTldSpamGetParamsARCNone, radar.EmailSecurityTopTldSpamGetParamsARCFail}),
			DateEnd:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange:   cloudflare.F([]string{"7d", "7d", "7d"}),
			DateStart:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DKIM:        cloudflare.F([]radar.EmailSecurityTopTldSpamGetParamsDKIM{radar.EmailSecurityTopTldSpamGetParamsDKIMPass, radar.EmailSecurityTopTldSpamGetParamsDKIMNone, radar.EmailSecurityTopTldSpamGetParamsDKIMFail}),
			DMARC:       cloudflare.F([]radar.EmailSecurityTopTldSpamGetParamsDMARC{radar.EmailSecurityTopTldSpamGetParamsDMARCPass, radar.EmailSecurityTopTldSpamGetParamsDMARCNone, radar.EmailSecurityTopTldSpamGetParamsDMARCFail}),
			Format:      cloudflare.F(radar.EmailSecurityTopTldSpamGetParamsFormatJson),
			Limit:       cloudflare.F(int64(5)),
			Name:        cloudflare.F([]string{"string", "string", "string"}),
			SPF:         cloudflare.F([]radar.EmailSecurityTopTldSpamGetParamsSPF{radar.EmailSecurityTopTldSpamGetParamsSPFPass, radar.EmailSecurityTopTldSpamGetParamsSPFNone, radar.EmailSecurityTopTldSpamGetParamsSPFFail}),
			TldCategory: cloudflare.F(radar.EmailSecurityTopTldSpamGetParamsTldCategoryClassic),
			TLSVersion:  cloudflare.F([]radar.EmailSecurityTopTldSpamGetParamsTLSVersion{radar.EmailSecurityTopTldSpamGetParamsTLSVersionTlSv1_0, radar.EmailSecurityTopTldSpamGetParamsTLSVersionTlSv1_1, radar.EmailSecurityTopTldSpamGetParamsTLSVersionTlSv1_2}),
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
