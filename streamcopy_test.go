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

func TestStreamCopyNewWithOptionalParams(t *testing.T) {
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
	)
	_, err := client.Stream.Copies.New(context.TODO(), cloudflare.StreamCopyNewParams{
		AccountID:      cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		URL:            cloudflare.F("https://example.com/myvideo.mp4"),
		AllowedOrigins: cloudflare.F([]string{"example.com"}),
		Creator:        cloudflare.F("creator-id_abcde12345"),
		Meta: cloudflare.F[any](map[string]interface{}{
			"name": "video12345.mp4",
		}),
		RequireSignedURLs:     cloudflare.F(true),
		ScheduledDeletion:     cloudflare.F(time.Now()),
		ThumbnailTimestampPct: cloudflare.F(0.529241),
		Watermark: cloudflare.F(cloudflare.StreamCopyNewParamsWatermark{
			Uid: cloudflare.F("ea95132c15732412d22c1476fa83f27a"),
		}),
		UploadCreator:  cloudflare.F("creator-id_abcde12345"),
		UploadMetadata: cloudflare.F("name aGVsbG8gd29ybGQ=, requiresignedurls, allowedorigins ZXhhbXBsZS5jb20sdGVzdC5jb20="),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
