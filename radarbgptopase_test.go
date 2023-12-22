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

func TestRadarBgpTopAseListWithOptionalParams(t *testing.T) {
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
	_, err := client.Radars.Bgps.Tops.Ases.List(context.TODO(), cloudflare.RadarBgpTopAseListParams{
		ASN:        cloudflare.F([]string{"15169", "15169", "15169"}),
		DateEnd:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:  cloudflare.F([]cloudflare.RadarBgpTopAseListParamsDateRange{cloudflare.RadarBgpTopAseListParamsDateRange1d, cloudflare.RadarBgpTopAseListParamsDateRange7d, cloudflare.RadarBgpTopAseListParamsDateRange14d}),
		DateStart:  cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:     cloudflare.F(cloudflare.RadarBgpTopAseListParamsFormatJson),
		Limit:      cloudflare.F(int64(5)),
		Name:       cloudflare.F([]string{"main_series", "main_series", "main_series"}),
		Prefix:     cloudflare.F([]string{"1.1.1.0/24", "1.1.1.0/24", "1.1.1.0/24"}),
		UpdateType: cloudflare.F([]cloudflare.RadarBgpTopAseListParamsUpdateType{cloudflare.RadarBgpTopAseListParamsUpdateTypeAnnouncement, cloudflare.RadarBgpTopAseListParamsUpdateTypeWithdrawal}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
