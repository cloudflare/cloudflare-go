// File generated from our OpenAPI spec by Stainless.

package cloudflare_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/cloudflare-go/internal/shared"
	"github.com/cloudflare/cloudflare-go/internal/testutil"
	"github.com/cloudflare/cloudflare-go/option"
)

func TestCachePurgeWithOptionalParams(t *testing.T) {
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
	_, err := client.Cache.Purge(context.TODO(), cloudflare.CachePurgeParams{
		ZoneID:          cloudflare.F("string"),
		Files:           cloudflare.F([]cloudflare.CachePurgeParamsFile{shared.UnionString("http://www.example.com/css/styles.css"), shared.UnionString("http://www.example.com/css/styles.css"), shared.UnionString("http://www.example.com/css/styles.css")}),
		Hosts:           cloudflare.F([]string{"www.example.com", "images.example.com"}),
		Prefixes:        cloudflare.F([]string{"www.example.com/foo", "images.example.com/bar/baz"}),
		PurgeEverything: cloudflare.F(true),
		Tags:            cloudflare.F([]string{"some-tag", "another-tag"}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
