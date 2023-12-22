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

func TestRadarBgpLeakEventListWithOptionalParams(t *testing.T) {
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
	_, err := client.Radars.Bgps.Leaks.Events.List(context.TODO(), cloudflare.RadarBgpLeakEventListParams{
		DateEnd:     cloudflare.F(time.Now()),
		DateRange:   cloudflare.F(cloudflare.RadarBgpLeakEventListParamsDateRange1d),
		DateStart:   cloudflare.F(time.Now()),
		Format:      cloudflare.F(cloudflare.RadarBgpLeakEventListParamsFormatJson),
		InvolvedASN: cloudflare.F(int64(0)),
		LeakASN:     cloudflare.F(int64(0)),
		Page:        cloudflare.F(int64(0)),
		PerPage:     cloudflare.F(int64(0)),
		SortBy:      cloudflare.F(cloudflare.RadarBgpLeakEventListParamsSortByID),
		SortOrder:   cloudflare.F(cloudflare.RadarBgpLeakEventListParamsSortOrderAsc),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
