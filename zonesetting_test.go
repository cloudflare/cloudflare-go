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

func TestZoneSettingEdit(t *testing.T) {
	t.Skip("oneOf doesnt match")
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
	_, err := client.Zones.Settings.Edit(context.TODO(), cloudflare.ZoneSettingEditParams{
		ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Items: cloudflare.F([]cloudflare.ZoneSettingEditParamsItem{cloudflare.ZonesAlwaysOnlineParam(cloudflare.ZonesAlwaysOnlineParam{
			ID:    cloudflare.F(cloudflare.ZonesAlwaysOnlineIDAlwaysOnline),
			Value: cloudflare.F(cloudflare.ZonesAlwaysOnlineValueOn),
		}), cloudflare.ZonesBrowserCacheTTLParam(cloudflare.ZonesBrowserCacheTTLParam{
			ID:    cloudflare.F(cloudflare.ZonesBrowserCacheTTLIDBrowserCacheTTL),
			Value: cloudflare.F(cloudflare.ZonesBrowserCacheTTLValue18000),
		}), cloudflare.ZonesIPGeolocationParam(cloudflare.ZonesIPGeolocationParam{
			ID:    cloudflare.F(cloudflare.ZonesIPGeolocationIDIPGeolocation),
			Value: cloudflare.F(cloudflare.ZonesIPGeolocationValueOff),
		})}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestZoneSettingGet(t *testing.T) {
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
	_, err := client.Zones.Settings.Get(context.TODO(), cloudflare.ZoneSettingGetParams{
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
