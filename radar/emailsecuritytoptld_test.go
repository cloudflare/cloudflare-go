// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/cloudflare-go/internal/testutil"
	"github.com/cloudflare/cloudflare-go/option"
	"github.com/cloudflare/cloudflare-go/radar"
)

func TestEmailSecurityTopTldGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.Top.Tlds.Get(context.TODO(), radar.EmailSecurityTopTldGetParams{
		ARC:         cloudflare.F([]radar.EmailSecurityTopTldGetParamsARC{radar.EmailSecurityTopTldGetParamsARCPass, radar.EmailSecurityTopTldGetParamsARCNone, radar.EmailSecurityTopTldGetParamsARCFail}),
		DateEnd:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   cloudflare.F([]string{"7d", "7d", "7d"}),
		DateStart:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DKIM:        cloudflare.F([]radar.EmailSecurityTopTldGetParamsDKIM{radar.EmailSecurityTopTldGetParamsDKIMPass, radar.EmailSecurityTopTldGetParamsDKIMNone, radar.EmailSecurityTopTldGetParamsDKIMFail}),
		DMARC:       cloudflare.F([]radar.EmailSecurityTopTldGetParamsDMARC{radar.EmailSecurityTopTldGetParamsDMARCPass, radar.EmailSecurityTopTldGetParamsDMARCNone, radar.EmailSecurityTopTldGetParamsDMARCFail}),
		Format:      cloudflare.F(radar.EmailSecurityTopTldGetParamsFormatJson),
		Limit:       cloudflare.F(int64(5)),
		Name:        cloudflare.F([]string{"string", "string", "string"}),
		SPF:         cloudflare.F([]radar.EmailSecurityTopTldGetParamsSPF{radar.EmailSecurityTopTldGetParamsSPFPass, radar.EmailSecurityTopTldGetParamsSPFNone, radar.EmailSecurityTopTldGetParamsSPFFail}),
		TldCategory: cloudflare.F(radar.EmailSecurityTopTldGetParamsTldCategoryClassic),
		TLSVersion:  cloudflare.F([]radar.EmailSecurityTopTldGetParamsTLSVersion{radar.EmailSecurityTopTldGetParamsTLSVersionTlSv1_0, radar.EmailSecurityTopTldGetParamsTLSVersionTlSv1_1, radar.EmailSecurityTopTldGetParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
