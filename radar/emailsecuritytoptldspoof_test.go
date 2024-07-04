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

func TestEmailSecurityTopTldSpoofGetWithOptionalParams(t *testing.T) {
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
		radar.EmailSecurityTopTldSpoofGetParamsSpoofSpoof,
		radar.EmailSecurityTopTldSpoofGetParams{
			ARC:         cloudflare.F([]radar.EmailSecurityTopTldSpoofGetParamsARC{radar.EmailSecurityTopTldSpoofGetParamsARCPass, radar.EmailSecurityTopTldSpoofGetParamsARCNone, radar.EmailSecurityTopTldSpoofGetParamsARCFail}),
			DateEnd:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange:   cloudflare.F([]string{"7d", "7d", "7d"}),
			DateStart:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DKIM:        cloudflare.F([]radar.EmailSecurityTopTldSpoofGetParamsDKIM{radar.EmailSecurityTopTldSpoofGetParamsDKIMPass, radar.EmailSecurityTopTldSpoofGetParamsDKIMNone, radar.EmailSecurityTopTldSpoofGetParamsDKIMFail}),
			DMARC:       cloudflare.F([]radar.EmailSecurityTopTldSpoofGetParamsDMARC{radar.EmailSecurityTopTldSpoofGetParamsDMARCPass, radar.EmailSecurityTopTldSpoofGetParamsDMARCNone, radar.EmailSecurityTopTldSpoofGetParamsDMARCFail}),
			Format:      cloudflare.F(radar.EmailSecurityTopTldSpoofGetParamsFormatJson),
			Limit:       cloudflare.F(int64(5)),
			Name:        cloudflare.F([]string{"string", "string", "string"}),
			SPF:         cloudflare.F([]radar.EmailSecurityTopTldSpoofGetParamsSPF{radar.EmailSecurityTopTldSpoofGetParamsSPFPass, radar.EmailSecurityTopTldSpoofGetParamsSPFNone, radar.EmailSecurityTopTldSpoofGetParamsSPFFail}),
			TldCategory: cloudflare.F(radar.EmailSecurityTopTldSpoofGetParamsTldCategoryClassic),
			TLSVersion:  cloudflare.F([]radar.EmailSecurityTopTldSpoofGetParamsTLSVersion{radar.EmailSecurityTopTldSpoofGetParamsTLSVersionTlSv1_0, radar.EmailSecurityTopTldSpoofGetParamsTLSVersionTlSv1_1, radar.EmailSecurityTopTldSpoofGetParamsTLSVersionTlSv1_2}),
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
