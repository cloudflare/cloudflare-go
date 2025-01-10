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

func TestStreamNewWithOptionalParams(t *testing.T) {
	t.Skip("TODO: investigate broken test")
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
	err := client.Stream.New(context.TODO(), stream.StreamNewParams{
		AccountID:      cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Body:           map[string]interface{}{},
		TusResumable:   cloudflare.F(stream.StreamNewParamsTusResumable1_0_0),
		UploadLength:   cloudflare.F(int64(0)),
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

func TestStreamListWithOptionalParams(t *testing.T) {
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
	_, err := client.Stream.List(context.TODO(), stream.StreamListParams{
		AccountID:     cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Asc:           cloudflare.F(true),
		Creator:       cloudflare.F("creator-id_abcde12345"),
		End:           cloudflare.F(time.Now()),
		IncludeCounts: cloudflare.F(true),
		Search:        cloudflare.F("puppy.mp4"),
		Start:         cloudflare.F(time.Now()),
		Status:        cloudflare.F(stream.StreamListParamsStatusPendingupload),
		Type:          cloudflare.F("live"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestStreamDelete(t *testing.T) {
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
	err := client.Stream.Delete(
		context.TODO(),
		"ea95132c15732412d22c1476fa83f27a",
		stream.StreamDeleteParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		},
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestStreamEditWithOptionalParams(t *testing.T) {
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
	_, err := client.Stream.Edit(
		context.TODO(),
		"ea95132c15732412d22c1476fa83f27a",
		stream.StreamEditParams{
			AccountID:          cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			AllowedOrigins:     cloudflare.F([]stream.AllowedOriginsParam{"example.com"}),
			Creator:            cloudflare.F("creator-id_abcde12345"),
			MaxDurationSeconds: cloudflare.F(int64(1)),
			Meta: cloudflare.F[any](map[string]interface{}{
				"name": "video12345.mp4",
			}),
			RequireSignedURLs:     cloudflare.F(true),
			ScheduledDeletion:     cloudflare.F(time.Now()),
			ThumbnailTimestampPct: cloudflare.F(0.529241),
			UploadExpiry:          cloudflare.F(time.Now()),
		},
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestStreamGet(t *testing.T) {
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
	_, err := client.Stream.Get(
		context.TODO(),
		"ea95132c15732412d22c1476fa83f27a",
		stream.StreamGetParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		},
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
