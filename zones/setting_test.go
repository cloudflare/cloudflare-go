// File generated from our OpenAPI spec by Stainless.

package zones_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/cloudflare-go/internal/testutil"
	"github.com/cloudflare/cloudflare-go/option"
	"github.com/cloudflare/cloudflare-go/zones"
)

func TestSettingEdit(t *testing.T) {
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
	_, err := client.Zones.Settings.Edit(context.TODO(), zones.SettingEditParams{
		ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Items: cloudflare.F([]zones.SettingEditParamsItem{zones.SettingEditParamsItemsZonesAlwaysOnline(zones.SettingEditParamsItemsZonesAlwaysOnline{
			ID:    cloudflare.F(zones.SettingEditParamsItemsZonesAlwaysOnlineIDAlwaysOnline),
			Value: cloudflare.F(zones.SettingEditParamsItemsZonesAlwaysOnlineValueOn),
		}), zones.SettingEditParamsItemsZonesBrowserCacheTTL(zones.SettingEditParamsItemsZonesBrowserCacheTTL{
			ID:    cloudflare.F(zones.SettingEditParamsItemsZonesBrowserCacheTTLIDBrowserCacheTTL),
			Value: cloudflare.F(zones.SettingEditParamsItemsZonesBrowserCacheTTLValue18000),
		}), zones.SettingEditParamsItemsZonesIPGeolocation(zones.SettingEditParamsItemsZonesIPGeolocation{
			ID:    cloudflare.F(zones.SettingEditParamsItemsZonesIPGeolocationIDIPGeolocation),
			Value: cloudflare.F(zones.SettingEditParamsItemsZonesIPGeolocationValueOff),
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

func TestSettingGet(t *testing.T) {
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
	_, err := client.Zones.Settings.Get(context.TODO(), zones.SettingGetParams{
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
