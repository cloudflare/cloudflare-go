// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/stream"
)

func TestDirectUploadNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Stream.DirectUpload.New(context.TODO(), stream.DirectUploadNewParams{
		AccountID:          cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		MaxDurationSeconds: cloudflare.F(int64(1)),
		AllowedOrigins:     cloudflare.F([]stream.AllowedOriginsParam{"example.com"}),
		Creator:            cloudflare.F("creator-id_abcde12345"),
		Expiry:             cloudflare.F(time.Now()),
		Meta: cloudflare.F[any](map[string]interface{}{
			"name": "video12345.mp4",
		}),
		RequireSignedURLs:     cloudflare.F(true),
		ScheduledDeletion:     cloudflare.F(time.Now()),
		ThumbnailTimestampPct: cloudflare.F(0.529241),
		Watermark: cloudflare.F(stream.DirectUploadNewParamsWatermark{
			UID: cloudflare.F("ea95132c15732412d22c1476fa83f27a"),
		}),
		UploadCreator: cloudflare.F("creator-id_abcde12345"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
