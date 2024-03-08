// File generated from our OpenAPI spec by Stainless.

package cloudflare_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/cloudflare-go/internal/testutil"
	"github.com/cloudflare/cloudflare-go/option"
)

func TestRadarBGPTopPrefixesWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.BGP.Top.Prefixes(context.TODO(), cloudflare.RadarBGPTopPrefixesParams{
		ASN:        cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:  cloudflare.F([]cloudflare.RadarBGPTopPrefixesParamsDateRange{cloudflare.RadarBGPTopPrefixesParamsDateRange1d, cloudflare.RadarBGPTopPrefixesParamsDateRange2d, cloudflare.RadarBGPTopPrefixesParamsDateRange7d}),
		DateStart:  cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:     cloudflare.F(cloudflare.RadarBGPTopPrefixesParamsFormatJson),
		Limit:      cloudflare.F(int64(5)),
		Name:       cloudflare.F([]string{"string", "string", "string"}),
		UpdateType: cloudflare.F([]cloudflare.RadarBGPTopPrefixesParamsUpdateType{cloudflare.RadarBGPTopPrefixesParamsUpdateTypeAnnouncement, cloudflare.RadarBGPTopPrefixesParamsUpdateTypeWithdrawal}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
