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

func TestRadarAttackLayer7SummaryManagedRuleListWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.Summaries.ManagedRules.List(context.TODO(), cloudflare.RadarAttackLayer7SummaryManagedRuleListParams{
		ASN:               cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:           cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:         cloudflare.F([]cloudflare.RadarAttackLayer7SummaryManagedRuleListParamsDateRange{cloudflare.RadarAttackLayer7SummaryManagedRuleListParamsDateRange1d, cloudflare.RadarAttackLayer7SummaryManagedRuleListParamsDateRange2d, cloudflare.RadarAttackLayer7SummaryManagedRuleListParamsDateRange7d}),
		DateStart:         cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:            cloudflare.F(cloudflare.RadarAttackLayer7SummaryManagedRuleListParamsFormatJson),
		HTTPMethod:        cloudflare.F([]cloudflare.RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethod{cloudflare.RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethodGet, cloudflare.RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethodPost, cloudflare.RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethodDelete}),
		HTTPVersion:       cloudflare.F([]cloudflare.RadarAttackLayer7SummaryManagedRuleListParamsHTTPVersion{cloudflare.RadarAttackLayer7SummaryManagedRuleListParamsHTTPVersionHttPv1, cloudflare.RadarAttackLayer7SummaryManagedRuleListParamsHTTPVersionHttPv2, cloudflare.RadarAttackLayer7SummaryManagedRuleListParamsHTTPVersionHttPv3}),
		IPVersion:         cloudflare.F([]cloudflare.RadarAttackLayer7SummaryManagedRuleListParamsIPVersion{cloudflare.RadarAttackLayer7SummaryManagedRuleListParamsIPVersionIPv4, cloudflare.RadarAttackLayer7SummaryManagedRuleListParamsIPVersionIPv6}),
		Location:          cloudflare.F([]string{"string", "string", "string"}),
		MitigationProduct: cloudflare.F([]cloudflare.RadarAttackLayer7SummaryManagedRuleListParamsMitigationProduct{cloudflare.RadarAttackLayer7SummaryManagedRuleListParamsMitigationProductDdos, cloudflare.RadarAttackLayer7SummaryManagedRuleListParamsMitigationProductWaf, cloudflare.RadarAttackLayer7SummaryManagedRuleListParamsMitigationProductBotManagement}),
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
