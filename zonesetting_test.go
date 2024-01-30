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
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithEmail("dev@cloudflare.com"),
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
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithEmail("dev@cloudflare.com"),
	)
	_, err := client.Zones.Settings.Edit(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.ZoneSettingEditParams{
			Items: cloudflare.F([]cloudflare.ZoneSettingEditParamsItem{cloudflare.ZoneSettingEditParamsItemsObject(cloudflare.ZoneSettingEditParamsItemsObject{
				ID:    cloudflare.F(cloudflare.ZoneSettingEditParamsItemsObjectIDAlwaysOnline),
				Value: cloudflare.F(cloudflare.ZoneSettingEditParamsItemsObjectValueOn),
			}), cloudflare.ZoneSettingEditParamsItemsObject(cloudflare.ZoneSettingEditParamsItemsObject{
				ID:    cloudflare.F(cloudflare.ZoneSettingEditParamsItemsObjectIDBrowserCacheTtl),
				Value: cloudflare.F(cloudflare.ZoneSettingEditParamsItemsObjectValue18000),
			}), cloudflare.ZoneSettingEditParamsItemsObject(cloudflare.ZoneSettingEditParamsItemsObject{
				ID:    cloudflare.F(cloudflare.ZoneSettingEditParamsItemsObjectIDIPGeolocation),
				Value: cloudflare.F(cloudflare.ZoneSettingEditParamsItemsObjectValueOff),
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
