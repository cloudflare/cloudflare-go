// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_firewall_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/dns_firewall"
	"github.com/cloudflare/cloudflare-go/v4/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v4/option"
)

func TestAnalyticsReportBytimeGetWithOptionalParams(t *testing.T) {
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
	_, err := client.DNSFirewall.Analytics.Reports.Bytimes.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		dns_firewall.AnalyticsReportBytimeGetParams{
			AccountID:  cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Dimensions: cloudflare.F("queryType"),
			Filters:    cloudflare.F("responseCode==NOERROR,queryType==A"),
			Limit:      cloudflare.F(int64(100)),
			Metrics:    cloudflare.F("queryCount,uncachedCount"),
			Since:      cloudflare.F(time.Now()),
			Sort:       cloudflare.F("+responseCode,-queryName"),
			TimeDelta:  cloudflare.F(dns_firewall.AnalyticsReportBytimeGetParamsTimeDeltaAll),
			Until:      cloudflare.F(time.Now()),
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
