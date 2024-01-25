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

func TestRadarAttackLayer7SummaryIPVersionListWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.Summaries.IPVersion.List(context.TODO(), cloudflare.RadarAttackLayer7SummaryIPVersionListParams{
		ASN:               cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:           cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:         cloudflare.F([]cloudflare.RadarAttackLayer7SummaryIPVersionListParamsDateRange{cloudflare.RadarAttackLayer7SummaryIPVersionListParamsDateRange1d, cloudflare.RadarAttackLayer7SummaryIPVersionListParamsDateRange2d, cloudflare.RadarAttackLayer7SummaryIPVersionListParamsDateRange7d}),
		DateStart:         cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:            cloudflare.F(cloudflare.RadarAttackLayer7SummaryIPVersionListParamsFormatJson),
		HTTPMethod:        cloudflare.F([]cloudflare.RadarAttackLayer7SummaryIPVersionListParamsHTTPMethod{cloudflare.RadarAttackLayer7SummaryIPVersionListParamsHTTPMethodGet, cloudflare.RadarAttackLayer7SummaryIPVersionListParamsHTTPMethodPost, cloudflare.RadarAttackLayer7SummaryIPVersionListParamsHTTPMethodDelete}),
		HTTPVersion:       cloudflare.F([]cloudflare.RadarAttackLayer7SummaryIPVersionListParamsHTTPVersion{cloudflare.RadarAttackLayer7SummaryIPVersionListParamsHTTPVersionHttPv1, cloudflare.RadarAttackLayer7SummaryIPVersionListParamsHTTPVersionHttPv2, cloudflare.RadarAttackLayer7SummaryIPVersionListParamsHTTPVersionHttPv3}),
		Location:          cloudflare.F([]string{"string", "string", "string"}),
		MitigationProduct: cloudflare.F([]cloudflare.RadarAttackLayer7SummaryIPVersionListParamsMitigationProduct{cloudflare.RadarAttackLayer7SummaryIPVersionListParamsMitigationProductDdos, cloudflare.RadarAttackLayer7SummaryIPVersionListParamsMitigationProductWaf, cloudflare.RadarAttackLayer7SummaryIPVersionListParamsMitigationProductBotManagement}),
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
