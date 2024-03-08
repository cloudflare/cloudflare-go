// File generated from our OpenAPI spec by Stainless.

package cloudflare_test

import (
	"context"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/cloudflare-go/internal/testutil"
	"github.com/cloudflare/cloudflare-go/option"
)

func TestUsage(t *testing.T) {
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
	zoneNewResponse, err := client.Zones.New(context.TODO(), cloudflare.ZoneNewParams{
		Account: cloudflare.F(cloudflare.ZoneNewParamsAccount{
			ID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		}),
		Name: cloudflare.F("example.com"),
		Type: cloudflare.F(cloudflare.ZoneNewParamsTypeFull),
	})
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", zoneNewResponse.ID)
}
