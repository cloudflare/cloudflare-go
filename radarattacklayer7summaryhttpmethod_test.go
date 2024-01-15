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

func TestRadarAttackLayer7SummaryHTTPMethodListWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.Summaries.HTTPMethod.List(context.TODO(), cloudflare.RadarAttackLayer7SummaryHTTPMethodListParams{
		ASN:               cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:           cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:         cloudflare.F([]cloudflare.RadarAttackLayer7SummaryHTTPMethodListParamsDateRange{cloudflare.RadarAttackLayer7SummaryHTTPMethodListParamsDateRange1d, cloudflare.RadarAttackLayer7SummaryHTTPMethodListParamsDateRange2d, cloudflare.RadarAttackLayer7SummaryHTTPMethodListParamsDateRange7d}),
		DateStart:         cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:            cloudflare.F(cloudflare.RadarAttackLayer7SummaryHTTPMethodListParamsFormatJson),
		HTTPVersion:       cloudflare.F([]cloudflare.RadarAttackLayer7SummaryHTTPMethodListParamsHTTPVersion{cloudflare.RadarAttackLayer7SummaryHTTPMethodListParamsHTTPVersionHttPv1, cloudflare.RadarAttackLayer7SummaryHTTPMethodListParamsHTTPVersionHttPv2, cloudflare.RadarAttackLayer7SummaryHTTPMethodListParamsHTTPVersionHttPv3}),
		IPVersion:         cloudflare.F([]cloudflare.RadarAttackLayer7SummaryHTTPMethodListParamsIPVersion{cloudflare.RadarAttackLayer7SummaryHTTPMethodListParamsIPVersionIPv4, cloudflare.RadarAttackLayer7SummaryHTTPMethodListParamsIPVersionIPv6}),
		Location:          cloudflare.F([]string{"string", "string", "string"}),
		MitigationProduct: cloudflare.F([]cloudflare.RadarAttackLayer7SummaryHTTPMethodListParamsMitigationProduct{cloudflare.RadarAttackLayer7SummaryHTTPMethodListParamsMitigationProductDdos, cloudflare.RadarAttackLayer7SummaryHTTPMethodListParamsMitigationProductWaf, cloudflare.RadarAttackLayer7SummaryHTTPMethodListParamsMitigationProductBotManagement}),
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
