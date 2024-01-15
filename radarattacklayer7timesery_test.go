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

func TestRadarAttackLayer7TimeseryListWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Attacks.Layer7.Timeseries.List(context.TODO(), cloudflare.RadarAttackLayer7TimeseryListParams{
		AggInterval:   cloudflare.F(cloudflare.RadarAttackLayer7TimeseryListParamsAggInterval1h),
		ASN:           cloudflare.F([]string{"string", "string", "string"}),
		Attack:        cloudflare.F([]cloudflare.RadarAttackLayer7TimeseryListParamsAttack{cloudflare.RadarAttackLayer7TimeseryListParamsAttackDdos, cloudflare.RadarAttackLayer7TimeseryListParamsAttackWaf, cloudflare.RadarAttackLayer7TimeseryListParamsAttackBotManagement}),
		DateEnd:       cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:     cloudflare.F([]cloudflare.RadarAttackLayer7TimeseryListParamsDateRange{cloudflare.RadarAttackLayer7TimeseryListParamsDateRange1d, cloudflare.RadarAttackLayer7TimeseryListParamsDateRange2d, cloudflare.RadarAttackLayer7TimeseryListParamsDateRange7d}),
		DateStart:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:        cloudflare.F(cloudflare.RadarAttackLayer7TimeseryListParamsFormatJson),
		Location:      cloudflare.F([]string{"string", "string", "string"}),
		Name:          cloudflare.F([]string{"string", "string", "string"}),
		Normalization: cloudflare.F(cloudflare.RadarAttackLayer7TimeseryListParamsNormalizationMin0Max),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
