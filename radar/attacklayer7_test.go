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

func TestAttackLayer7TimeseriesWithOptionalParams(t *testing.T) {
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
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.Radar.Attacks.Layer7.Timeseries(context.TODO(), radar.AttackLayer7TimeseriesParams{
		AggInterval:   cloudflare.F(radar.AttackLayer7TimeseriesParamsAggInterval1h),
		ASN:           cloudflare.F([]string{"string", "string", "string"}),
		Attack:        cloudflare.F([]radar.AttackLayer7TimeseriesParamsAttack{radar.AttackLayer7TimeseriesParamsAttackDDoS, radar.AttackLayer7TimeseriesParamsAttackWAF, radar.AttackLayer7TimeseriesParamsAttackBotManagement}),
		Continent:     cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:       cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:     cloudflare.F([]radar.AttackLayer7TimeseriesParamsDateRange{radar.AttackLayer7TimeseriesParamsDateRange1d, radar.AttackLayer7TimeseriesParamsDateRange2d, radar.AttackLayer7TimeseriesParamsDateRange7d}),
		DateStart:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:        cloudflare.F(radar.AttackLayer7TimeseriesParamsFormatJson),
		Location:      cloudflare.F([]string{"string", "string", "string"}),
		Name:          cloudflare.F([]string{"string", "string", "string"}),
		Normalization: cloudflare.F(radar.AttackLayer7TimeseriesParamsNormalizationMin0Max),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
