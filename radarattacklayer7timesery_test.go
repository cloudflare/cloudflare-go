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
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithEmail("dev@cloudflare.com"),
	)
	_, err := client.Radars.Attacks.Layer7s.Timeseries.List(context.TODO(), cloudflare.RadarAttackLayer7TimeseryListParams{
		AggInterval:   cloudflare.F(cloudflare.RadarAttackLayer7TimeseryListParamsAggInterval15m),
		ASN:           cloudflare.F([]string{"15169", "15169", "15169"}),
		Attack:        cloudflare.F([]cloudflare.RadarAttackLayer7TimeseryListParamsAttack{cloudflare.RadarAttackLayer7TimeseryListParamsAttackDdos, cloudflare.RadarAttackLayer7TimeseryListParamsAttackWaf, cloudflare.RadarAttackLayer7TimeseryListParamsAttackBotManagement}),
		DateEnd:       cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:     cloudflare.F([]cloudflare.RadarAttackLayer7TimeseryListParamsDateRange{cloudflare.RadarAttackLayer7TimeseryListParamsDateRange1d, cloudflare.RadarAttackLayer7TimeseryListParamsDateRange7d, cloudflare.RadarAttackLayer7TimeseryListParamsDateRange14d}),
		DateStart:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:        cloudflare.F(cloudflare.RadarAttackLayer7TimeseryListParamsFormatJson),
		Location:      cloudflare.F([]string{"US,CA", "US,CA", "US,CA"}),
		Name:          cloudflare.F([]string{"main_series", "main_series", "main_series"}),
		Normalization: cloudflare.F(cloudflare.RadarAttackLayer7TimeseryListParamsNormalizationPercentageChange),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
