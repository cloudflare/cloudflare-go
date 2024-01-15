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

func TestRadarAttackLayer7SummaryHTTPVersionListWithOptionalParams(t *testing.T) {
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithAPIToken("my-cloudflare-api-token"),
		option.WithUserServiceKey("my-cloudflare-user-service-key"),
	)
	_, err := client.Radar.Attacks.Layer7.Summaries.HTTPVersion.List(context.TODO(), cloudflare.RadarAttackLayer7SummaryHTTPVersionListParams{
		ASN:               cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:           cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:         cloudflare.F([]cloudflare.RadarAttackLayer7SummaryHTTPVersionListParamsDateRange{cloudflare.RadarAttackLayer7SummaryHTTPVersionListParamsDateRange1d, cloudflare.RadarAttackLayer7SummaryHTTPVersionListParamsDateRange2d, cloudflare.RadarAttackLayer7SummaryHTTPVersionListParamsDateRange7d}),
		DateStart:         cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:            cloudflare.F(cloudflare.RadarAttackLayer7SummaryHTTPVersionListParamsFormatJson),
		HTTPMethod:        cloudflare.F([]cloudflare.RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethod{cloudflare.RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethodGet, cloudflare.RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethodPost, cloudflare.RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethodDelete}),
		IPVersion:         cloudflare.F([]cloudflare.RadarAttackLayer7SummaryHTTPVersionListParamsIPVersion{cloudflare.RadarAttackLayer7SummaryHTTPVersionListParamsIPVersionIPv4, cloudflare.RadarAttackLayer7SummaryHTTPVersionListParamsIPVersionIPv6}),
		Location:          cloudflare.F([]string{"string", "string", "string"}),
		MitigationProduct: cloudflare.F([]cloudflare.RadarAttackLayer7SummaryHTTPVersionListParamsMitigationProduct{cloudflare.RadarAttackLayer7SummaryHTTPVersionListParamsMitigationProductDdos, cloudflare.RadarAttackLayer7SummaryHTTPVersionListParamsMitigationProductWaf, cloudflare.RadarAttackLayer7SummaryHTTPVersionListParamsMitigationProductBotManagement}),
		Name:              cloudflare.F([]string{"string", "string", "string"}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
