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

func TestRadarSearchGlobalListWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Searches.Globals.List(context.TODO(), cloudflare.RadarSearchGlobalListParams{
		Query:         cloudflare.F("United"),
		Exclude:       cloudflare.F([]cloudflare.RadarSearchGlobalListParamsExclude{cloudflare.RadarSearchGlobalListParamsExcludeSpecialEvents, cloudflare.RadarSearchGlobalListParamsExcludeNotebooks, cloudflare.RadarSearchGlobalListParamsExcludeLocations}),
		Format:        cloudflare.F(cloudflare.RadarSearchGlobalListParamsFormatJson),
		Include:       cloudflare.F([]cloudflare.RadarSearchGlobalListParamsInclude{cloudflare.RadarSearchGlobalListParamsIncludeSpecialEvents, cloudflare.RadarSearchGlobalListParamsIncludeNotebooks, cloudflare.RadarSearchGlobalListParamsIncludeLocations}),
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
