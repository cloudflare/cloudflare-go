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

func TestCacheVariantList(t *testing.T) {
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
	_, err := client.Cache.Variants.List(context.TODO(), cloudflare.CacheVariantListParams{
		ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCacheVariantDelete(t *testing.T) {
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
	_, err := client.Cache.Variants.Delete(context.TODO(), cloudflare.CacheVariantDeleteParams{
		ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCacheVariantEditWithOptionalParams(t *testing.T) {
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
	_, err := client.Cache.Variants.Edit(context.TODO(), cloudflare.CacheVariantEditParams{
		ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Value: cloudflare.F(cloudflare.CacheVariantEditParamsValue{
			Avif: cloudflare.F([]interface{}{"image/webp", "image/jpeg"}),
			Bmp:  cloudflare.F([]interface{}{"image/webp", "image/jpeg"}),
			Gif:  cloudflare.F([]interface{}{"image/webp", "image/jpeg"}),
			Jp2:  cloudflare.F([]interface{}{"image/webp", "image/avif"}),
			Jpeg: cloudflare.F([]interface{}{"image/webp", "image/avif"}),
			Jpg:  cloudflare.F([]interface{}{"image/webp", "image/avif"}),
			Jpg2: cloudflare.F([]interface{}{"image/webp", "image/avif"}),
			Png:  cloudflare.F([]interface{}{"image/webp", "image/avif"}),
			Tif:  cloudflare.F([]interface{}{"image/webp", "image/avif"}),
			Tiff: cloudflare.F([]interface{}{"image/webp", "image/avif"}),
			Webp: cloudflare.F([]interface{}{"image/jpeg", "image/avif"}),
		}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
