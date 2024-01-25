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

func TestZoneSettingList(t *testing.T) {
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
	_, err := client.Zones.Settings.List(context.TODO(), "023e105f4ecef8ad9ca31a8372d0c353")
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithAPIToken("my-cloudflare-api-token"),
		option.WithUserServiceKey("my-cloudflare-user-service-key"),
	)
	_, err := client.Zones.Settings.Edit(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.ZoneSettingEditParams{
			Items: cloudflare.F([]cloudflare.ZoneSettingEditParamsItem{cloudflare.ZoneSettingEditParamsItemsZonesAlwaysOnline(cloudflare.ZoneSettingEditParamsItemsZonesAlwaysOnline{
				ID:    cloudflare.F(cloudflare.ZoneSettingEditParamsItemsZonesAlwaysOnlineIDAlwaysOnline),
				Value: cloudflare.F(cloudflare.ZoneSettingEditParamsItemsZonesAlwaysOnlineValueOn),
			}), cloudflare.ZoneSettingEditParamsItemsZonesBrowserCacheTtl(cloudflare.ZoneSettingEditParamsItemsZonesBrowserCacheTtl{
				ID:    cloudflare.F(cloudflare.ZoneSettingEditParamsItemsZonesBrowserCacheTtlIDBrowserCacheTtl),
				Value: cloudflare.F(cloudflare.ZoneSettingEditParamsItemsZonesBrowserCacheTtlValue18000),
			}), cloudflare.ZoneSettingEditParamsItemsZonesIPGeolocation(cloudflare.ZoneSettingEditParamsItemsZonesIPGeolocation{
				ID:    cloudflare.F(cloudflare.ZoneSettingEditParamsItemsZonesIPGeolocationIDIPGeolocation),
				Value: cloudflare.F(cloudflare.ZoneSettingEditParamsItemsZonesIPGeolocationValueOff),
			})}),
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
