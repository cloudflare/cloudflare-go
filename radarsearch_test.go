// File generated from our OpenAPI spec by Stainless.

package cloudflare_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-sdk-go"
	"github.com/cloudflare/cloudflare-sdk-go/internal/testutil"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

func TestRadarSearchGlobalWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Search.Global(context.TODO(), cloudflare.RadarSearchGlobalParams{
		Query:         cloudflare.F("United"),
		Exclude:       cloudflare.F([]cloudflare.RadarSearchGlobalParamsExclude{cloudflare.RadarSearchGlobalParamsExcludeSpecialEvents, cloudflare.RadarSearchGlobalParamsExcludeNotebooks, cloudflare.RadarSearchGlobalParamsExcludeLocations}),
		Format:        cloudflare.F(cloudflare.RadarSearchGlobalParamsFormatJson),
		Include:       cloudflare.F([]cloudflare.RadarSearchGlobalParamsInclude{cloudflare.RadarSearchGlobalParamsIncludeSpecialEvents, cloudflare.RadarSearchGlobalParamsIncludeNotebooks, cloudflare.RadarSearchGlobalParamsIncludeLocations}),
		Limit:         cloudflare.F(int64(5)),
		LimitPerGroup: cloudflare.F(0.000000),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
